FROM golang as builder

WORKDIR /go/build/
COPY main.go .
COPY app .
COPY go.mod .
COPY go.sum .
COPY Makefile .

RUN go mod download
RUN make build

FROM alpine:3.6
RUN apk --no-cache add ca-certificates

WORKDIR /app/
RUN mkdir web

COPY web/css web/.
COPY --from=builder /go/build/karrlein .
COPY --from=builder /go/build/web/app.wasm web/.

EXPOSE 8080

CMD ["./karrlein"]
