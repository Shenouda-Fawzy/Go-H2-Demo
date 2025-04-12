**MIRRO REPO**  
Source: https://gitlab.com/Shenouda-Fawzy/ht2-demo

# Go-H2-Demo
This repository contain a quick demo for spinning up an HTTP2/TLS server

It contains three endpoints  
- `https://localhost:1234/` the main page
- `https://localhost:1234/syahello` I think you know what it does
- `https://localhost:1234/hooks` will print the request content which is expected to be encoded as `application/x-www-form-urlencoded`


## Running it
```bash
$ go run main.go
```

