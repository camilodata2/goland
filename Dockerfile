FROM ubuntu:latest
LABEL authors="camilo"

ENTRYPOINT ["top", "-b"]