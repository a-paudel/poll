FROM python:3.10-alpine
RUN pip install pdm
WORKDIR /app
COPY pyproject.toml .
RUN pdm install
COPY . .