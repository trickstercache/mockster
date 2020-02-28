# Mockster - Suite for mocking time series and byte range responses

Mockster is a suite of golang packages and accompanying runnable applications that facilitate unit testing of components by providing mock data on-the-fly.

Mockster aims to support all of the output formats that are accelerated by [Trickster](https://github.com/Comcast/trickster).

Any project is welcome to import the Mockster packages for testing these supported formats.

Mockster is the merger of two previously separate applications (PromSim and RangeSim) that were formerly part of the Trickster core project codebase.

## Running as a Standalone App

The project provides a standalone server implementation of Mockster. This is useful for backing full simulation environments or running a local background app for querying during development of a data consumer app. You can find it at `github.com/tricksterproxy/mockster/cmd/mockster`, and, from that working directory, simply run `go run main.go \[PORT\]`. If a port number is not provided, it defaults to 8482.

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
