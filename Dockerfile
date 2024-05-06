FROM golang:1.22.2 as BUILD
WORKDIR /src
COPY . .
RUN env GOOS=linux GOARCH=amd64 go build -o golang-app ./cmd/main.go
EXPOSE 3000
ENTRYPOINT ["./golang-app"]