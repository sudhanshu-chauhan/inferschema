FROM golang:1.17

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . /app/
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install supervisor -y
RUN go mod download && go mod verify
RUN go test inferschema/app
RUN go build -v -o /app/inferschema

CMD ["supervisord", "-c", "/app/supervisord.conf"]

