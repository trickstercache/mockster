# mockster/byterange

mockster/byterange simply prints out the `Lorem ipsum ...` sample text, pared down to the requested range or multipart ranges, with a few bells and whistles that allow you to customize its response for unit testing purposes. We make extensive use of mockster/byterange in unit testing to verify the integrity of Trickster's output after performing operations like merging disparate range parts, extracting ranges from other ranges, or from a full body, compressing adjacent ranges into a single range in the cache, etc.

## Supported Simulation Endpoints

- `/byterange/*` (Instantaneous)

Any path under `/byterange/` can be requested. This way, if you are testing byte ranges with a caching layer, each test's URL can very.

For examples of using mockster/byterange for Unit Testing, see the [Trickster Unit Tests](https://github.com/Comcast/trickster/blob/master/internal/proxy/engines/objectproxycache_test.go)
