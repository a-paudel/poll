import json
from secrets import token_urlsafe
import peewee as p
from playhouse.db_url import connect
import os

# set db_name based on env
# db_url = "dev.db" if os.getenv("ENV", "PROD") == "DEV" else "data/db.sqlite3"
db_url = os.getenv("DATABASE_URL")

# create db connection
# db = p.SqliteDatabase(db_url)
print("=================================")
print(db_url)
print("=================================")
db = connect(db_url)

# base model
class BaseModel(p.Model):
    class Meta:
        database = db


# Question model
# string question, answers, votes, code
class Question(BaseModel):
    question = p.TextField()
    answers = p.TextField()
    votes = p.TextField()
    code = p.TextField()

    def __str__(self):
        return self.question

    # from dict
    @classmethod
    def from_dict(cls, data):
        def _generate_code():
            random_code = token_urlsafe()[:5]
            all_codes = [q.code for q in cls.select()]
            while random_code in all_codes:
                return _generate_code()
            return random_code

        instance = cls()
        instance.question = data["question"]
        instance.answers = json.dumps(data["answers"])
        instance.votes = json.dumps([0 for _ in data["answers"]])
        instance.code = _generate_code()
        instance.save()
        return instance

    # to dict
    def to_dict(self):
        return {
            "question": self.question,
            "answers": json.loads(self.answers),
            "votes": json.loads(self.votes),
            "code": self.code,
        }

    # from code
    @classmethod
    def from_code(cls, code):
        instance = cls.get_or_none(cls.code == code)
        if instance:
            return instance
        return None


# create tables
db.create_tables([Question])
