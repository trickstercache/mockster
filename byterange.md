# mockster/byterange

mockster/byterange simply prints out the `Lorem ipsum ...` sample text, pared down to the requested byte range or multipart ranges, with a few bells and whistles that allow you to customize its response for unit testing purposes. We make extensive use of mockster/byterange in unit testing to verify the integrity of Trickster's output after performing operations like merging disparate range parts, extracting ranges from other ranges, or from a full body, compressing adjacent ranges into a single range in the cache, etc.

Provide a properly-formatted `Range` Request header, with 1 or more byte ranges.

The default output body is 1224 bytes. However, you can provide a `size=32768` parameter, where the value is any value integer representing the desired byte size of the output payload. Mockster will efficiently write to the wire the `Lorem ipsum` text until the byte size is met.

You can use the size parameter by itself to simply write out a very large payload for testing purposes.

Or you can combine a Range request with a size parameter, and the correct header and output bytes for the provided size and ranges will be efficiently written to the wire. If you provide a `Range` request without a size parameter, the content length will be set to highest-index byte in the range.

## Supported Simulation Endpoints

- `/byterange/*`

Any path under `/byterange/` can be requested. This way, if you are testing byte ranges with a caching layer, each test's URL can vary to avoid cache collisions. No matter the path provided, the response will always be the same, as described above.

Example paths:
    - `/byterange/test/1`
    - `/byterange/`
    - `/byterange/tests/testCacheMissWithRange`

For examples of using mockster/byterange for Unit Testing, see the [Trickster Unit Tests](https://github.com/tricksterproxy/trickster/blob/master/internal/proxy/engines/objectproxycache_test.go)

## Example Usage

```bash
$ go run cmd/mockster/main.go &
Starting up Mockster on port 8482

# make a multipart range request for a 2448-byte object
$ curl -v -H 'Range: bytes=1220-1260,1265-1270' -H 'Connection: close'  'http://0.0.0.0:8482/byterange/?size=2448'
*   Trying 0.0.0.0...
* TCP_NODELAY set
* Connected to 0.0.0.0 (127.0.0.1) port 8482 (#0)
> GET /byterange/?size=2448 HTTP/1.1
> Host: 0.0.0.0:8482
> User-Agent: curl/7.64.1
> Accept: */*
> Range: bytes=1220-1260,1265-1270
> Connection: close
>
< HTTP/1.1 206 Partial Content
< Cache-Control: max-age=60
< Content-Type: multipart/byteranges; boundary=TestRangeServerBoundary
< Last-Modified: Wed, 01 Jan 2020 00:00:00 UTC
< Date: Wed, 01 Jan 2020 00:00:00 UTC
< Content-Length: 292
< Connection: close
<
--TestRangeServerBoundary
Content-Range: bytes 1220-1260/2448
Content-Type: text/plain; charset=utf-8

m.

Lorem ipsum dolor sit amet, mel alia
--TestRangeServerBoundary
Content-Range: bytes 1265-1270/2448
Content-Type: text/plain; charset=utf-8

nieba
--TestRangeServerBoundary--
* Closing connection 0

# request a 500MB response of lorem ipsum text
$ curl -v -H 'Connection: close' -o /dev/null  'http://0.0.0.0:8482/byterange/?size=524288000'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying 0.0.0.0...
* TCP_NODELAY set
* Connected to 0.0.0.0 (127.0.0.1) port 8482 (#0)
> GET /byterange/?size=524288000 HTTP/1.1
> Host: 0.0.0.0:8482
> User-Agent: curl/7.64.1
> Accept: */*
> Connection: close
>
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: max-age=60
< Content-Length: 524288000
< Content-Type: text/plain; charset=utf-8
< Last-Modified: Wed, 01 Jan 2020 00:00:00 UTC
< Date: Thu, 26 Mar 2020 03:33:40 GMT
< Connection: close
<
{ [12050 bytes data]
100  500M  100  500M    0     0   710M      0 --:--:-- --:--:-- --:--:--  710M
* Closing connection 0

```
