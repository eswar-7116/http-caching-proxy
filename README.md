# HTTP Caching Proxy
HTTP Caching Proxy is a simple, minimal caching proxy server implemented in Go. It serves as a basic solution for caching HTTP responses, using an in-memory map to store the cache entries.

## Features
- ### In-Memory Caching:
  The server uses an in-memory map combined with an LRU list to store cached responses. This approach allows for fast lookups and reduces the time needed to retrieve cached data.

- ### Proxy Functionality:
  The server forwards client requests to the target server and caches the responses. If the same request is made again, the server returns the cached response, saving the time and resources of making a new request to the target server.

- ### Accepts only HTTP/HTTPS URLs
  The proxy accepts only HTTP or HTTPS URLs for upstream requests.

- ### Time-based and LRU Eviction Policy:
  Cache entries use lazy TTL expiration combined with an LRU eviction policy. Expired and least-recently-used entries are removed when capacity exceeds and refetched from upstream for freshness.

- ### Simple Design:
  This is a minimalistic implementation aimed at demonstrating the core concepts of a caching proxy server. It is not suitable for production use without further enhancements.

## Usage

### Start the server:

```bash
go run ./cmd/proxy
```

### Invoke the server
```bash
curl 'http://localhost:8000?url=<URL>'
```

## Limitations

- ### Memory Usage:
  Since the cache is stored in memory, it is limited by the available RAM. For large-scale use cases, consider implementing a more sophisticated caching mechanism that can handle larger datasets or persist the cache to disk.

- ### No Size-based Eviction Policy:

  This implementation does not include a size-based cache eviction policy. Eviction is based on entry count, not response size. Large responses may consume disproportionate memory.

- ### Not a real forward-proxy

  This is an application-level fetch proxy and cannot be configured as a browser/system forward proxy.

## Conclusion
This project is a minimal, educational HTTP fetch-based caching proxy designed to demonstrate core caching concepts such as TTL handling, LRU eviction, header handling, and request forwarding. It is not a full forward proxy and is intended as a foundation for learning and experimentation rather than production deployment.

<p hidden>Built as a solution to https://roadmap.sh/projects/caching-server</p>

---
#### <div align="center">If you like this project, please give this repo a star 🌟</div>
---
