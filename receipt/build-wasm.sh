tinygo build -o fyx.wasm -target wasi ./receipt.go
gzip -f -k fyx.wasm