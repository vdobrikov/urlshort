# urlshort
URL shortener implemented with Go lang

## How to Use
Start the server:
```
$ cd $GOPATH/src
$ git clone https://github.com/vdobrikov/urlshort.git
$ cd urlshort
$ go run cmd/app/main.go
Starting the server on :8080
```
Now we can check HTTP redirect:
```
$ curl -i http://localhost:8080/urlshort

HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: https://github.com/gophercises/urlshort
Date: Sat, 23 Mar 2019 07:45:48 GMT
Content-Length: 74

<a href="https://github.com/gophercises/urlshort">Moved Permanently</a>.
```
URLs are stored in ``resources/urls.yaml`` file.