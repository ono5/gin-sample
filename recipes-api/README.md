# recipes-api

## cURL

```
curl -s -X GET 'http://localhost:8080/recipes' | jq length
492
```

## Open API

- [goswagger](https://goswagger.io/install.html)
  - [binary](https://github.com/go-swagger/go-swagger/releases)

```
brew tap go-swagger/go-swagger
brew install go-swagger

or
docker pull quay.io/goswagger/swagger
alias swagger="docker run --rm -it --user $(id -u):$(id -g) -e XDG_CACHE_HOME=/tmp/.cache -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
```

```
swagger version
version: v0.27.0
```

```
swagger generate spec -o ./swagger.json
```

## Redis

```
docker ps
0c20a2e39d6e   redis:latest          "docker-entrypoint.s…"   15 minutes ago   Up 2 minutes   0.0.0.0:6379->6379/tcp, :::6379->6379/tcp       recipes-api_redis_1

docker exec -it 0c20a2e39d6e bash

redis-cli

```

## Performance benchmark

- [Apache Benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html)

```
# キャッシュなし
ab -n 2000 -c 100 -g without-cache.data http://localhost:8080/recipes

Time taken for tests:   148.191 seconds
Time per request:       7409.556 [ms] (mean)
Time per request:       74.096 [ms] (mean, across all concurrent requests)

# キャッシュあり
ab -n 2000 -c 100 -g with-cache.data http://localhost:8080/recipes

Time taken for tests:   134.045 seconds
Time per request:       6702.274 [ms] (mean)
Time per request:       67.023 [ms] (mean, across all concurrent requests)
```

- Time taken for tests: This means the total time to complete the 2,000 rrequests
- Time per request: This means how many milliseconds it takes to complete one request.

グラフ化

```
brew install gnuplot

gnuplot apache-benchmark.p
```

## JWT Decode

[jwt](https://jwt.io/)
