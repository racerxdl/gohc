Golang Health Check
===================

Simple Health Check Library to make Docker / Kubernetes / Rancher / Swarm be able to detect if your small program is up.

Simple to use:

```go
(...)
hc := gohc.MakeHealthCheck(func() bool {
    // Test here for anything that should be up
    return true
})

hc.Listen(":8000")
(...)
```

If the argument function returns true, gohc will return `StatusCode 200` and `OK`. If not it will return `StatusCode 500` and `unhealthy` on body.

You can also use it in your own http server instance:

```go
(...)
hc := gohc.MakeHealthCheck(func() bool {
    // Test here for anything that should be up
    return true
})


m := mux.Router()
m.HandleFunc("/health-check", hc.ServeHTTP)
(...)
```