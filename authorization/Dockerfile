FROM python:3.11

WORKDIR /authz

COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

WORKDIR /authz/src

CMD [ "uvicorn", "main:app"] 