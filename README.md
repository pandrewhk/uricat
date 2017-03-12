# uricat
streaming parallel downloader (multi-file, not multi-part)

Takes JSON stream as input and produces a JSON stream with results. Requests are processed in parallel and output as soon as complete.

```json
% echo '{ "Url": "http://httpbin.org/get?1" }'|go run uricat.go|jq .
{
  "Url": "http://httpbin.org/get?1",
  "Status": "200 OK",
  "StatusCode": 200,
  "Proto": "HTTP/1.0",
  "Header": {
    "Access-Control-Allow-Credentials": [
      "true"
    ],
    "Access-Control-Allow-Origin": [
      "*"
    ],
    "Connection": [
      "close"
    ],
    "Content-Length": [
      "256"
    ],
    "Content-Type": [
      "application/json"
    ],
    "Date": [
      "Sun, 12 Mar 2017 09:51:09 GMT"
    ],
    "Server": [
      "nginx"
    ]
  },
  "Body": {
    "args": {
      "1": ""
    },
    "headers": {
      "Accept-Encoding": "gzip",
      "Cache-Control": "max-age=0",
      "Host": "httpbin.org",
      "User-Agent": "Go-http-client/1.1"
    },
    "origin": "212.46.218.245",
    "url": "http://httpbin.org/get?1"
  }
}
```
