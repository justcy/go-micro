package main

import (
	"embed"
	"go-micro.kanter.cn/v1/cmd"

	_ "go-micro.kanter.cn/v1/cmd/micro/cli"
	_ "go-micro.kanter.cn/v1/cmd/micro/run"
	"go-micro.kanter.cn/v1/cmd/micro/server"
)

//go:embed web/styles.css web/main.js web/templates/*
var webFS embed.FS

var version = "5.0.0-dev"

func init() {
	server.HTML = webFS
}

func main() {
	cmd.Init(
		cmd.Name("micro"),
		cmd.Version(version),
	)
}
