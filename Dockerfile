FROM golang:1.16.3 AS builder
WORKDIR src 
COPY . .
RUN GOOS=darwin GOARCH=amd64 go build -o micro-service-handler
FROM scratch 
COPY --from=builder src/micro-service-handler /opt/micro-service-handler
CMD ["/opt/micro-service-handler"]
