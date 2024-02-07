FROM golang:1.22-alpine3.18 as build
RUN apk add build-base libpcap-dev
ADD . /src
WORKDIR /src
RUN go mod tidy
RUN go test ./...
RUN go vet ./...
RUN CGO_ENABLED=1 go build -o /app main.go

FROM alpine:3.18
COPY --from=build /app /app
ENTRYPOINT ["/app"]
