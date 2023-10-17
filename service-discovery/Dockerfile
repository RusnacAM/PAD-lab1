FROM python:3.11-slim-bookworm

WORKDIR /service-discovery

COPY ./requirements.txt /service-discovery

RUN pip3 install -r requirements.txt

COPY . .

EXPOSE 5051

ENV FLASK_APP=main.py

CMD ["python3", "-m", "flask", "run", "--host=0.0.0.0", "--port=5051"]