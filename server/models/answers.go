package models

type Answer struct {
	ID         int    `json:"id" db:"id"`
	Answer     string `json:"answer" db:"answer"`
	Votes      int    `json:"votes" db:"votes"`
	QuestionID int    `json:"question_id" db:"question_id"`
}

func createAnswerTable() {
	DB.Exec(`
		create table if not exists answers (
			id integer not null primary key autoincrement,
			answer text not null,
			votes integer not null default 0,
			question_id integer not null
		)
	`)
}

func (a *Answer) Create(answer string, questionID int) error {
	var err = DB.QueryRow(`
		insert into answers (answer, question_id) values (?, ?) returning *;
	`, answer, questionID).Scan(&a.ID, &a.Answer, &a.Votes, &a.QuestionID)
	return err
}

func (a *Answer) Delete() error {
	var err = DB.QueryRow(`
		delete from answers where id = ? returning *;
	`, a.ID).Scan(&a.ID, &a.Answer, &a.Votes, &a.QuestionID)
	return err
}

func (a *Answer) Vote(id int) error {
	var err = DB.QueryRow(`
		update answers set votes = votes + 1 where id = ? returning *;
	`, id).Scan(&a.ID, &a.Answer, &a.Votes, &a.QuestionID)
	return err
}
