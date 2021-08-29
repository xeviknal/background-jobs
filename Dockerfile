# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

LABEL maintainer=xeviknal@gmail.com

WORKDIR /app

copy go.mod ./
copy go.sum ./

RUN go mod download

copy . ./

RUN go build -o background-jobs
RUN chmod +x background-jobs

EXPOSE 3306

CMD [ "./background-jobs" ]