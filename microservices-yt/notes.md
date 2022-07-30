# Notes

## Basics of Server
[Official Docs](https://pkg.go.dev/net/http)
- Servemux is an HTTP request multiplexer
  1. Incoming URL - matched against registered patterns
  2. Appropriate handler is called
- Can close a server gracefully using a channel and signals (`chan` and `os.Signal`)
- Also, server's `Shutdown(context)` method handles active requests, then shuts down

## JSON Encoding
[Official Docs](https://pkg.go.dev/encoding/json)
- Marshal : JSON --> string
- instead of calling `marshal()`, use `Encoder.encode()`
- difference
  - `json.marshal(v any)` - returns `[]byte, err`
  - Use an io.Writer to write
  - `Encoder.encode(v any)` - returns `err`
  - `Encoder` already takes a io.Writer, so this writes right away
  - `Encoder.encode` doesn't use memory, so it is preferable

## POST && PUT
- Implement `FromJSON` for `Product` using `Decoder.decode()`
- Implement `getNextID()` and `findProduct(id int)`
- Used the [regexp package in the standard library](https://pkg.go.dev/regexp) to parse URL params (for PUT)
  - Very tedious and potentially error prone
- Declare an Error, `ErrProductNotFound` to be used in the controller

