# build the Go APP
FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# copy the build assets to a minimal
# alpine image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main .
RUN chmod +x ./main
EXPOSE 8080
CMD ./main