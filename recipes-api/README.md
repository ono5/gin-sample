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
