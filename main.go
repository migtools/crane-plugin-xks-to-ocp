package main

import (
	"github.com/konveyor/crane-lib/transform/cli"
	xkstoocp "github.com/migtools/crane-plugin-xks-to-ocp/xks-to-ocp"
)

func main() {
	cli.RunAndExit(&xkstoocp.Plugin{})
}
