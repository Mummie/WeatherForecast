FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/weatherapp

FROM scratch

COPY --from=builder /go/bin/weatherapp /weatherapp

EXPOSE 8080

ENTRYPOINT ["/weatherapp"]
