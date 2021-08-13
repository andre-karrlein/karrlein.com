build:
	@go build -o karrlein ./app

wasm:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm ./app

run: build wasm
	./karrlein