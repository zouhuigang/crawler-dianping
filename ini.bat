@echo off

if not exist config mkdir config
go run mainex/ini/ini-jianshu-5.go
go run mainex/ini/ini-jianshu-6.go
go run mainex/ini/ini-oschina-5.go

echo finished
