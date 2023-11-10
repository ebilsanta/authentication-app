FROM python:3.11

WORKDIR /resource

COPY . .
RUN pip install --no-cache-dir -r requirements.txt

WORKDIR /resource/src

CMD ["uvicorn", "main:app"] 