FROM python:3.10-alpine
# RUN adduser app -D
# USER app
WORKDIR /app

RUN pip install pipenv
COPY Pipfile .
COPY Pipfile.lock .
RUN pipenv install --deploy --system
COPY . .
CMD [ "gunicorn", "app.app:app" ]