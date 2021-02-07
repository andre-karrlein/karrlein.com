build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm app/app.go
	go build -o karrlein

run: build
	./karrlein