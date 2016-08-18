# Concurrence Request Go #

- Use to test the time os request for websites.
- Concurrence Request 
- Version: 1.0.0


## Getting Started

The minimum requirement of GO is 1.5. ([Install Go](https://golang.org/dl/))

To install **crgo**
```
$ go get github.com/rlopes-fixeads/crgo
```
    
To compile **crgo**
```
$ go build -x crgo main.go
```
 
## How to use

How to use the **crgo**
```
$ crgo [url] [url] [url] ...
```
    
Example to use with urls.    
```
$ ./crgo http://www.google.com/ http://www.facebook.com/ http://www.terra.com.br/

[0.13s] to request with [11517] for [http://www.google.com/]
[0.95s] to request with [88022] for [http://www.facebook.com/]
[0.92s] to request with [261321] for [http://www.terra.com.br/]
[2.19s] Request time.
```

## License

This project is under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for the full license text.