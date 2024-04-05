FROM python:3.12.2-slim-bookworm

WORKDIR /app

COPY . .

RUN pip install -r requirnments.txt

EXPOSE 9500

CMD ["python3", "main.py"]
