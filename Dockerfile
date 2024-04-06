FROM golang:1.22-alpine as build-base

ARG DATABASE_URL

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY .env /app

COPY . .

RUN CGO_ENABLED=0 go test -tags=unit ./...

RUN go build -o ./out/go-app .

FROM alpine:3.16.2
COPY --from=build-base /app/out/go-app /app/go-app

ENV DATABASE_URL=postgres://root:password@localhost:5432/wallet?sslmode=disable

CMD ["/app/go-app"]

