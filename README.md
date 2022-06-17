# "Word of Wisdom” TCP Client/Server 

Test task for Server Engineer - Design and implement “Word of Wisdom” tcp server.  

Simple "Word of Wisdom” TCP-server implementation.

For each incoming request, generated and send to client a token with given accuracy (1/difficulty). 
The client needs to find **nonce** such that the final hash will be with a given accuracy.

``` 
hash(token || nonce) < 1/difficulty
```

SHA-256 is used as the PoW hash-function.

The current difficulty is determined based on the average number of incoming requests.

```
difficulty ~ averageCountOfRequests
```

The difficulty increases proportionally as the number of requests increases.
So difficulty limits the number of incoming requests.

## Docker container 
### Build docker
```
make
```

### Run server
```
docker run -p 8080:8080 wow-server
``` 

### Run client
```
docker run -it wow-client
```
