# Simple Poll App

Simple poll app built using flask for the back-end and alpinejs to increase interactivity on the front-end.

One of the very first apps I made to learn python.

Now dockerized and running [here](https://poll.paudel.me).

## How it works

The app has a page to create polls, and a page to vote on polls.

### Create Page `/`

The create page uses two methods, GET and POST. The GET method is used to render the form. The POST method takes json of the format

```json
{
    "question":"some question",
    "answers":["answer1", "answer2"]
}
```



 and returns 

```json
{
    "code":"someCode"
}
```

.



### Vote Page `/<code>`

The vote page again has a GET and POST methods. The GET method is used to render the vote page. The POST method again takes a json of format

```json
{
    "index":3 //can be any number
}
```



.

### API Route `/api/<code>`

A GET request to this route returns json of the format

```json
{
    "question":"some question",
    "answers":["answer1", "answer2"],
    "votes":[0, 1],
    "code":"some code"
}
```



AlpineJs is used to hydrate the page with this data, switch to result view and finally use browser localstorage to limit one user to one vote.
