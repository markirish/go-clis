/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/markirish/go-clis/clis/cobra/cmd"
	_ "github.com/markirish/go-clis/clis/cobra/cmd/list"
	_ "github.com/markirish/go-clis/clis/cobra/cmd/list/pods"
)

func main() {
	cmd.Execute()
}
