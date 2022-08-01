import json
from flask import Flask, render_template, request, redirect, url_for, flash, jsonify
from app.models import Question, db

app = Flask(__name__)


@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "POST":
        data = request.get_json()
        if data:
            question = Question.from_dict(data)
            return jsonify(question.to_dict())
        return jsonify({"error": "Invalid data"}), 400

    return render_template("index.html")


# vote page gui
@app.route("/<code>", methods=["GET", "POST"])
def vote(code):
    question = Question.from_code(code)
    if request.method == "POST":
        data = request.get_json()
        if data:
            if question:
                index = data["index"]
                with db.atomic():
                    votes = json.loads(question.votes)
                    votes[index] += 1
                    question.votes = json.dumps(votes)
                    question.save()
                print(question)
                return jsonify(question.to_dict())
            return jsonify({"error": "No data"}), 404
        return jsonify({"error": "Invalid data"}), 400

    # get request
    if question:
        return render_template("vote.html")
    return "Not found", 404


# vote page api
@app.route("/api/<code>", methods=["GET"])
def vote_api(code):
    question = Question.from_code(code)
    # get request
    if question:
        return jsonify(question.to_dict())
    return jsonify({"error": "No data"}), 404
