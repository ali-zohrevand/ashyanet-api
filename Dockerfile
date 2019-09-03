FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
#RUN apk add git

RUN go mod download
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' .
FROM scratch
COPY --from=builder /app/ashyanet-api /app/
EXPOSE 5000
ENTRYPOINT ["/app/ashyanet-api"]