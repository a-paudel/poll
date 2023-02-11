package models

type Question struct {
	ID       int    `json:"id" db:"id"`
	Question string `json:"question" db:"question"`

	Answers []Answer `json:"answers" db:"-"`
}

func createQuestionTable() {
	DB.Exec(`
		create table if not exists questions (
			id integer not null primary key autoincrement,
			question text not null
		);
	`)
}

func (q *Question) getAnswers() error {
	rows, err := DB.Query(`
		select * from answers where question_id = ?;
	`, q.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ans Answer
		if err := rows.Scan(&ans.ID, &ans.Answer, &ans.Votes, &ans.QuestionID); err != nil {
			return err
		}
		q.Answers = append(q.Answers, ans)
	}
	return nil
}

func (q *Question) Create(question string) error {
	var err = DB.QueryRow(`
		insert into questions (question) values (?) returning *;
	`, question).Scan(&q.ID, &q.Question)
	if err != nil {
		return err
	}
	err = q.getAnswers()
	return err
}

func (q *Question) Get(id int) error {
	var err = DB.QueryRow(`
		select * from questions where id = ?;
	`, id).Scan(&q.ID, &q.Question)
	if err != nil {
		return err
	}
	err = q.getAnswers()
	return err
}

func (q *Question) Delete() error {
	_, err := DB.Exec(`
		delete from questions where id = ?;
	`, q.ID)
	if err != nil {
		return err
	}
	// delete all answers for this question
	_, err = DB.Exec(`
		delete from answers where question_id = ?;
	`, q.ID)
	return err
}
