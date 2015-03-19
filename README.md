## tlsget

A simple curl-like diagnostics utility, useful for fetching SSL websites from a host where DNS does not resolve to the hostname you are interested in.


### Usage

Grab [precompiled binary](https://github.com/porjo/tlsget/releases/), or compile using Go

```
./tlsget -h mywebsite.com https://clusterhost25.com
```

Optionally, specify `-i` flag to see response headers.

