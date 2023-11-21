# photogo
A photo album in Go.


### Mod

```bash
go mod init github.com/prantoran/photogo
go mod tidy
```

### Live preloading

#### Install
```bash
go install github.com/cosmtrek/air@latest
go get github.com/gorilla/mux
go get github.com/gorilla/schema
```
#### Run directly
```
air --build.cmd "go build ." --build.bin "./photogo"
```
#### Run using config
```
air init
air
```