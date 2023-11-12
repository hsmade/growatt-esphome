FROM golang:1.21 as build
ADD . /src
WORKDIR /src
RUN go mod tidy
RUN go mod vendor
RUN go test ./...
RUN go vet ./...
RUN go build -o /app main.go

FROM scratch
COPY --from=build /app /app
ENTRYPOINT ["/app"]
