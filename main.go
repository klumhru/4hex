package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/klumhru/4hex/viz"
)

func main() {
	opts := &viz.Options{}
	flags.Parse(opts)
	viz.Run(opts)
}
