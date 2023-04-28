```sh
$ GOOS=wasip1 GOARCH=wasm gotip build -o main.wasm .

$ wazero run -mount / -env-inherit main.wasm
.gitignore
README.md
.git
main.wasm
main.go
go.mod

$ wasmtime run --dir / --env PWD=$PWD -- main.wasm
panic: open /home/jabaile/work/gowasi: Is a directory

goroutine 1 [running]:
main.main()
        /home/jabaile/work/gowasi/main.go:16 +0x23

$ wasmer run --dir / --env PWD=$PWD -- main.wasm     
panic: stat .: Bad file number

goroutine 1 [running]:
main.main()
        /home/jabaile/work/gowasi/main.go:11 +0x26
```
