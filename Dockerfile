FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY service service

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/server

FROM golang:1.23 AS runtime

WORKDIR /app

COPY --from=build /app/server server
COPY places.json .

EXPOSE 8080

CMD ["/app/server"]
