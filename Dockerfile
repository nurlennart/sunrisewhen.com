FROM python:latest

RUN pip install --upgrade pip
COPY requirements.txt .
RUN pip install -r requirements.txt

COPY /app /

CMD [ "gunicorn", "-b", "0.0.0.0:8000", "app:app"]
