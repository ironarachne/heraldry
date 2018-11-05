# build stage
FROM golang:1.11 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/armsapi

# final stage
FROM scratch
COPY --from=build-env /go/bin/armsapi /go/bin/armsapi
EXPOSE 7476
CMD ["/go/bin/armsapi"]