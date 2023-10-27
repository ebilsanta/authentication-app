FROM python:3.11

WORKDIR /authz

COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD [ "uvicorn", "main:app"] 