tinygo build -o fyx.wasm -target wasi ./main.go
gzip -f -k fyx.wasm