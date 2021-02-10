echo "-- build started"
go build -o "bin/debug/rest_api_server.elf" -i -v "cmd/server/main.go"
echo "-- build completed"
exec ./bin/debug/rest_api_server.elf %*
rm bin/debug/rest_api_server.elf
echo "-- binary removed"