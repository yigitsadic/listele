FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY . .

RUN go build -o app

FROM alpine:3.13

WORKDIR /src

ARG PORT="3035"
ENV PORT=$PORT
EXPOSE $PORT

COPY --from=compiler /src/app/db /src/db
COPY --from=compiler /src/app/app /src/app
CMD ["/src/app"]
