About
-------
[False Sharing](https://en.wikipedia.org/wiki/False_sharing)

Result
-------
```
≈60000000000ns 44053681537 ops # golang no paddiing
≈60000000000ns 48919048088 ops # golang padding
≈60000000000ns 25754666867 ops # c++ no padding
≈60000000000ns 49408074006 ops # c++ padding
```