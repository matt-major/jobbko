FROM alpine:3.13.6

RUN mkdir /app

ADD ./jobbko /app

WORKDIR /app

CMD ["./jobbko"]
