package main

import (
	"flag"

	"github.com/fujimaru-lab/fuditor/pkg/fuditor"
)

func main() {
	fuditor.RunFuditor(flag.Args())
}
