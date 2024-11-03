FROM golang:latest

WORKDIR /app

COPY Makefile ./
COPY go.mod go.sum ./
RUN make tidy

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping main.go

EXPOSE 8080

CMD [ "/docker-gs-ping" ]