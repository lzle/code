# gRPC

### Prerequisites 

1、Go, any one of the two latest major releases of Go.

For installation instructions, see Go’s [Getting Started](https://go.dev/doc/install) guide.


2、Protocol buffer compiler, protoc, version 3.

For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).


3、Go plugins for the protocol compiler:

* Install the protocol compiler plugins for Go using the following commands:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

* Update your PATH so that the protoc compiler can find the plugins:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

检查版本：

```bash
$ protoc --version
libprotoc 3.20.1   # 3.x 版本支持 syntax = "proto3"

$ protoc-gen-go --version
protoc-gen-go v1.28.0

$ protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.2.0
```

### Defining the service

`gRPC` 可以定义四种服务类型：

* 普通 rpc，客户端向服务器发送一个请求，然后得到一个响应，就像普通的函数调用；

* 服务器流式 rpc，客户端向服务器发送一个请求，然后得到一个流式的响应；

* 客户端流式 rpc，客户端向服务器发送一个流式的请求，然后得到一个响应；

* 双向流式 rpc，客户端向服务器发送一个流式的请求，然后得到一个流式的响应。

```proto
// Interface exported by the server.
service RouteGuide {
  // A simple RPC.
  //
  // Obtains the feature at a given position.
  //
  // A feature with an empty name is returned if there's no feature at the given
  // position.
  rpc GetFeature(Point) returns (Feature) {}

  // A server-to-client streaming RPC.
  //
  // Obtains the Features available within the given Rectangle.  Results are
  // streamed rather than returned at once (e.g. in a response message with a
  // repeated field), as the rectangle may cover a large area and contain a
  // huge number of features.
  rpc ListFeatures(Rectangle) returns (stream Feature) {}

  // A client-to-server streaming RPC.
  //
  // Accepts a stream of Points on a route being traversed, returning a
  // RouteSummary when traversal is completed.
  rpc RecordRoute(stream Point) returns (RouteSummary) {}

  // A Bidirectional streaming RPC.
  //
  // Accepts a stream of RouteNotes sent while a route is being traversed,
  // while receiving other RouteNotes (e.g. from other users).
  rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
}
```

### Generating client and server code 

执行以下命令在 `pb` 目录下生成 `route_guide.pb.go` 和 `route_guide_grpc.pb.go` 文件：

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
```

### Run the sample code
To compile and run the server, assuming you are in the root of the `route_guide`
folder, i.e., `.../examples/route_guide/`, simply:

```sh
$ go run server/server.go
```

Likewise, to run the client:

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```

and

```sh
$ go run client/client.go -tls=true
```

### References

[Quick start](https://grpc.io/docs/languages/go/quickstart/)

[Basics tutorial](https://grpc.io/docs/languages/go/basics/)