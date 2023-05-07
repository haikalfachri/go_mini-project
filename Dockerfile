FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o entry

EXPOSE 8000

ENTRYPOINT [ "./entry" ]