FROM golang:1.16-alpine

ENV PORT=8000

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod download

RUN go build -o ./out/docker-test-app

EXPOSE ${PORT}

CMD [ "./out/docker-test-app" ]