.PHONY: sw sw-serve sw-serve2 go

sw:
	alias swagger="docker run --rm -it -p 9999:9999 --user $(id -u):$(id -g) -e XDG_CACHE_HOME=/tmp/.cache -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
	swagger generate spec -o ./swagger.json

sw-serve:
	swagger serve ./swagger.json --no-open -p 9999

sw-serve2:
	swagger serve -F swagger ./swagger.json --no-open -p 9999

go:
	MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE="demo" go run main.go
