# gin-sample

## Utility Command

### Dependency module output into file

```
go mod graph | sed -Ee 's/@[^[:blank:]]+//g' | sort | uniq >unver.txt
cat unver.txt | awk '{print "\""$1"\" -> \""$2"\""};' >>graph.dot
echo "}" >>graph.dot
sed -i '' 's+\("github.com/[^/]*/\)\([^"]*"\)+\1\\n\2+g' graph.dot

brew install graphviz
open graph.svg
```

## MongoDB

- [mongo-drivert](https://pkg.go.dev/go.mongodb.org/mongo-driver)

```
go get go.mongodb.org/mongo-driver/mongo
```
