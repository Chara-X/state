# Concurrency, consistency, TTL state manager

```go
func ExampleClient() {
    var client = &Client{
        Address: "http://localhost:8080",
        Client:  &http.Client{},
    }
    var value []byte
    var ok bool
    var etag string
    value, ok, etag = client.Get("0")
    fmt.Println(string(value), ok, etag)
    client.Set("0", []byte("a"), 3*time.Second, "")
    value, ok, etag = client.Get("0")
    fmt.Println(string(value), ok, etag)
    client.Set("0", []byte("b"), 3*time.Second, etag)
    value, ok, etag = client.Get("0")
    fmt.Println(string(value), ok, etag)
    // Output:
    //  false
    // a true 3a3af51b-fbb5-49fa-983b-604b2251fac5
    // b true 554de156-d361-4f5f-b2b1-8a86d7817ba6
}
```

## Reference

[State management overview](https://docs.dapr.io/developing-applications/building-blocks/state-management/state-management-overview)
