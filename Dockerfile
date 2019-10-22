FROM python:3.7.5-alpine

COPY delete-junk.py /
COPY requirements.txt /tmp
RUN pip3 install -r /tmp/requirements.txt

WORKDIR /
CMD ["python3", "delete-junk.py"]
