---
layout: default
---

# Plugins

Plugins are scoped under each interface directory within this repository. To use a plugin, import it directly from the corresponding interface subpackage and pass it to your service via options.

Common interfaces and locations:
- Registry: `go-micro.kanter.cn/v1/registry/*` (e.g. `consul`, `etcd`, `nats`, `mdns`)
- Broker: `go-micro.kanter.cn/v1/broker/*` (e.g. `nats`, `rabbitmq`, `http`, `memory`)
- Transport: `go-micro.kanter.cn/v1/transport/*` (e.g. `nats`, default `http`)
- Server: `go-micro.kanter.cn/v1/server/*` (e.g. `grpc` for native gRPC compatibility)
- Client: `go-micro.kanter.cn/v1/client/*` (e.g. `grpc` for native gRPC compatibility)
- Store: `go-micro.kanter.cn/v1/store/*` (e.g. `postgres`, `mysql`, `nats-js-kv`, `memory`)
- Auth, Cache, etc. follow the same pattern under their respective directories.

## Registry Examples

Consul:
```go
import (
    "go-micro.kanter.cn/v1"
    "go-micro.kanter.cn/v1/registry/consul"
)

func main() {
    reg := consul.NewConsulRegistry()
    svc := micro.NewService(
        micro.Registry(reg),
    )
    svc.Init()
    svc.Run()
}
```

Etcd:
```go
import (
    "go-micro.kanter.cn/v1"
    "go-micro.kanter.cn/v1/registry/etcd"
)

func main() {
    reg := etcd.NewRegistry()
    svc := micro.NewService(micro.Registry(reg))
    svc.Init()
    svc.Run()
}
```

## Broker Examples

NATS:
```go
import (
    "go-micro.kanter.cn/v1"
    bnats "go-micro.kanter.cn/v1/broker/nats"
)

func main() {
    b := bnats.NewNatsBroker()
    svc := micro.NewService(micro.Broker(b))
    svc.Init()
    svc.Run()
}
```

RabbitMQ:
```go
import (
    "go-micro.kanter.cn/v1"
    "go-micro.kanter.cn/v1/broker/rabbitmq"
)

func main() {
    b := rabbitmq.NewBroker()
    svc := micro.NewService(micro.Broker(b))
    svc.Init()
    svc.Run()
}
```

## Transport Example (NATS)
```go
import (
    "go-micro.kanter.cn/v1"
    tnats "go-micro.kanter.cn/v1/transport/nats"
)

func main() {
    t := tnats.NewTransport()
    svc := micro.NewService(micro.Transport(t))
    svc.Init()
    svc.Run()
}
```

## gRPC Server/Client (Native gRPC Compatibility)

For native gRPC compatibility (required for `grpcurl`, polyglot gRPC clients, etc.), use the gRPC server and client plugins. Note: This is different from the gRPC transport.

```go
import (
    "go-micro.kanter.cn/v1"
    grpcServer "go-micro.kanter.cn/v1/server/grpc"
    grpcClient "go-micro.kanter.cn/v1/client/grpc"
)

func main() {
    svc := micro.NewService(
        micro.Server(grpcServer.NewServer()),
        micro.Client(grpcClient.NewClient()),
    )
    svc.Init()
    svc.Run()
}
```

See [Native gRPC Compatibility](guides/grpc-compatibility.md) for a complete guide.

## Store Examples

Postgres:
```go
import (
    "go-micro.kanter.cn/v1"
    postgres "go-micro.kanter.cn/v1/store/postgres"
)

func main() {
    st := postgres.NewStore()
    svc := micro.NewService(micro.Store(st))
    svc.Init()
    svc.Run()
}
```

NATS JetStream KV:
```go
import (
    "go-micro.kanter.cn/v1"
    natsjskv "go-micro.kanter.cn/v1/store/nats-js-kv"
)

func main() {
    st := natsjskv.NewStore()
    svc := micro.NewService(micro.Store(st))
    svc.Init()
    svc.Run()
}
```

## Notes
- Defaults: If you don’t set an implementation, Go Micro uses sensible in-memory or local defaults (e.g., mDNS for registry, HTTP transport, memory broker/store).
- Options: Each plugin exposes constructor options to configure addresses, credentials, TLS, etc.
- Imports: Only import the plugin you need; this keeps binaries small and dependencies explicit.
