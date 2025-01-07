package main

import (
	"embed"
	"json-server/cmd"
)

//go:embed embed/*
var templates embed.FS

func main() {
	cmd.Execute()
}
