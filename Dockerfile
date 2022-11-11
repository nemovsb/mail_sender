FROM alpine

RUN apk update  && apk add --no-cache ca-certificates

CMD ["/bin/sh", "-c", "./mail_sender"]

EXPOSE 8081

WORKDIR /mail_sender

COPY . /mail_sender/

RUN chmod +x mail_sender