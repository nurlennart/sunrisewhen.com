FROM python:latest

RUN pip install --upgrade pip
COPY requirements.txt .
RUN pip install -r requirements.txt

ADD /app/src/app.py /
ADD /app/src/config.py /
ADD /app/src/routes.py /
ADD /app/src/localizer.py /
ADD /app/templates /templates
ADD /app/static /static
ADD GeoLite2-City.mmdb /

CMD [ "gunicorn", "-b", "0.0.0.0:8000", "app:app"]
