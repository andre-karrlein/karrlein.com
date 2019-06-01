FROM golang as builder

WORKDIR /go/build/
COPY karrlein.go .

RUN go get -u golang.org/x/crypto/...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o karrlein .

FROM alpine:3.6
RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY public/ public/
COPY template.html .
COPY resume.html .
COPY --from=builder /go/build/karrlein .

EXPOSE 8080

CMD ["./karrlein"]
