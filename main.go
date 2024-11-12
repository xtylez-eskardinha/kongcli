/*
Copyright © 2024 Miguel Angel Sanchez <EMAIL ADDRESS>
*/

package main

import (
	"kongcli/cmd"
	_ "kongcli/cmd/consumer"
	_ "kongcli/cmd/route"
)

func main() {
	cmd.Execute()
}
