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
[0.14s] elapsed time for request [http://www.google.com/] with [11546] 
[0.86s] elapsed time for request [http://www.facebook.com/] with [88051] 
[0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] 
[1.77s] elapsed time.
```

## License

This project is under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for the full license text.