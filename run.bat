@echo off
rem pushd %~dp0
rem pushd ..
echo -- build started
go build -o "bin\debug\rest_api_server.exe" -i -v "cmd\server\main.go"
echo -- build completed
bin\debug\rest_api_server.exe %*
del bin\debug\rest_api_server.exe
echo -- binary removed
rem popd
rem popd