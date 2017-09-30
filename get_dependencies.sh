#!/usr/bin/env bash
go get github.com/labstack/echo &
go get github.com/labstack/echo/middleware
wait
go run spa_server.go