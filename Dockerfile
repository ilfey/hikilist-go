FROM ubuntu:latest
LABEL authors="ilfey"

ENTRYPOINT ["top", "-b"]