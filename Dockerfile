FROM golang:1.25 as build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base

COPY --from=build /app/main .

COPY --from=build /app/static ./static

EXPOSE 8081

CMD ["/main"]
