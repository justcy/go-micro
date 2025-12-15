---
layout: default
---

# NATS Transport

Use NATS as the transport between services.

## In code

```go
package main

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

## Via environment

Run your service with env vars set:

```bash
MICRO_TRANSPORT=nats MICRO_TRANSPORT_ADDRESS=nats://127.0.0.1:4222 go run main.go
```
