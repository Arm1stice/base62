base62
====== 

This package allows you to convert between uint64 numbers and base62 encoded strings. It was designed with URL shorteners in mind, which is why it only encodes uint64 numbers (I created this package to be used in [my own personal URL shortener](https://wcal.xyz)).

[![GoDoc](https://godoc.org/github.com/wcalandro/base62?status.svg)](http://godoc.org/github.com/wcalandro/base62)  

## It has two functions

```go 
// Takes a uint64 typed number and returns a base62 encoded string
func ToB62(uint64) string {}

// Takes a base62 encoded string and returns a uint64 typed number and an optional error if there is an invalid character in the string
func FromB62(string) (uint64, error) {}
```

## Using the package
You can import the package using the following command

```bash
go get github.com/wcalandro/base62
```

