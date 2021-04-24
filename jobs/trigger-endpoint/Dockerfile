FROM python:3.9.4-slim-buster

ADD requirements.txt .
ADD trigger_endpoint.py .
RUN python3 -m pip install -r requirements.txt


CMD ["python3", "trigger_endpoint.py"]