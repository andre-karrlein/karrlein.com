FROM golang:alpine as builder

WORKDIR /go/build/
RUN mkdir app
COPY main.go .
COPY app/app.go app/.
COPY go.mod .
COPY go.sum .
COPY Makefile .

RUN go mod download
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm app/app.go
RUN go build -o karrlein

FROM alpine:3.6
RUN apk --no-cache add ca-certificates

WORKDIR /app/
RUN mkdir -p web/css
RUN mkdir -p web/images

COPY web/css/main.css web/css/.
COPY web/images/ak_logo_*.png web/images/
COPY --from=builder /go/build/karrlein .
COPY --from=builder /go/build/web/app.wasm web/.

EXPOSE 8080

CMD ["./karrlein"]
