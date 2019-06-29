echo 'delete old file'
rm -rf dist
echo 'delete success ~~'

echo 'build...'
go build -o dist/iris-admin-api main.go
echo 'build success ~~'