CURRENT=$(cd $(dirname $0);pwd)
echo $CURRENT
cd ../

# win用にbuild
GOOS=windows GOARCH=amd64 go build -o dist/agse2e.exe ./main.go
# mac用にbuild
go build -o dist/agse2e ./main.go