# Mockster - Suite for mocking time series and byte range responses

Mockster is a suite of golang packages and accompanying runnable applications that facilitate unit testing of components by providing mock data on-the-fly.

Mockster aims to support all of the output formats that are accelerated by [Trickster](https://github.com/Comcast/trickster).

Any project is welcome to import the Mockster packages for testing these supported formats.

Mockster is the merger of two previously separate applications (PromSim and RangeSim) that were formerly part of the Trickster core project codebase.

## Running as a Standalone App

The project provides a standalone server implementation of Mockster. This is useful for backing full simulation environments or running a local background app for querying during development of a data consumer app. You can find it at `github.com/trickstercache/mockster/cmd/mockster`, and, from that working directory, simply run `go run main.go \[PORT\]`. If a port number is not provided, it defaults to 8482.

## Running from Docker

We offer Mockster as an image on Docker Hub at `trickstercache/mockster`:

```bash
$ docker run --rm -p 8482:8482 trickstercache/mockster:latest &
  Starting up mockster 1.1.1 on port 8482


$ curl -v http://127.0.0.1:8482/prometheus/api/v1/query?query=up
> GET /prometheus/api/v1/query?query=up HTTP/1.1
> Host: 127.0.0.1:8482
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Tue, 05 May 2020 14:17:10 GMT
< Content-Length: 117
<
{"status":"success","data":{"resultType":"vector","result":[{"metric":{"series_id":"0"},"value":[1588688230,"76"]}]}}


$ curl -v -H 'Range: bytes=0-10,12-25' http://127.0.0.1:8482/byterange/test1
> GET /byterange/test1 HTTP/1.1
> Host: 127.0.0.1:8482
> User-Agent: curl/7.64.1
> Accept: */*
> Range: bytes=0-10,12-25
>
< HTTP/1.1 206 Partial Content
< Cache-Control: max-age=60
< Content-Type: multipart/byteranges; boundary=TestRangeServerBoundary
< Last-Modified: Wed, 01 Jan 2020 00:00:00 UTC
< Date: Tue, 05 May 2020 14:17:15 GMT
< Content-Length: 263
<
--TestRangeServerBoundary
Content-Range: bytes 0-10/1224
Content-Type: text/plain; charset=utf-8

Lorem ipsum
--TestRangeServerBoundary
Content-Range: bytes 12-25/1224
Content-Type: text/plain; charset=utf-8

dolor sit amet
--TestRangeServerBoundary--


$ fg
$ <ctrl + c>

```

## Output Formats

Mockster currently supports mocking the following outputs:

- HTTP Byte Ranges
- Prometheus

### mockster/byterange

mockster/byterange simply prints out the `Lorem ipsum ...` sample text, pared down to the requested range or multipart ranges, with a few bells and whistles that allow you to customize its response for unit testing purposes. We make extensive use of mockster/byterange in unit testing to verify the integrity of Trickster's output after performing operations like merging disparate range parts, extracting ranges from other ranges, or from a full body, compressing adjacent ranges into a single range in the cache, etc.

Read more about using mockster/byterange for your needs [here](./byterange.md).

### mockster/prometheus

mockster/prometheus can output data in the Prometheus format, which consists of values that are repeatably generatable for the provided query and timerange inputs. The data output by mockster/prometheus does not represent reality in any way, and is only useful for unit testing and integration testing, by providing a synthesized Prometheus environment that outputs meaningless data. None of mockster/prometheus's result sets are stored on or retrieved from disk, and are calculated just-in-time on every request, using simple mathematical computations. In Trickster, we use mockster/prometheus to conduct end-to-end testing of our DeltaProxyCache during unit testing, without requiring a real Prometheus server.

Read more about using and customizing mockster/prometheus for your needs [here](./prometheus.md).
