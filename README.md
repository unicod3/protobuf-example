# INSTALL

## Download and Install the compiler https://github.com/protocolbuffers/protobuf/releases/

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go

$ export PATH=$PATH:$GOPATH/bin
```


# USAGE
```bash
$ protoc --go_out=. *.proto
```


```go
import proto "github.com/golang/protobuf/proto"
```

```bash
$ go run .
```
