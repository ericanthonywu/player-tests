FROM golang:latest

RUN apt-get update && apt-get install -y git

RUN go env -w GOPRIVATE=gitlab.com/socialbread/mobile/jwt-middleware
RUN go env -w GO111MODULE=on

RUN echo "machine gitlab.com\nlogin ibnufajar\npassword glpat-Ai81zrQDx_wwMwyrmegg" > ~/.netrc

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

RUN touch .env

EXPOSE 6722

CMD ["sh", "-c", "/app/main"]
