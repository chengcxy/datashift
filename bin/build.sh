cd ../cmd/
echo $PWD
echo "go build -o datashift main.go"
go build -o datashift main.go
echo "$GOPATH/bin"
mv datashift $GOPATH/bin/
