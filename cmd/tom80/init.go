//+build !wasm

package main

import (
	"flag"
	"os"

	"github.com/superloach/tom80"
)

func init() {
	cons = tom80.MkTom80()

	game := flag.String("game", "", "`name` of game to load")
	version := flag.Bool("version", false, "print version info")
	ipf := flag.Int("ipf", tom80.DefaultIPF, "`number` of instructions per frame")
	flag.Parse()

	if *game == "" {
		version_info()
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		version_info()
		os.Exit(0)
	}

	cons.IPF = *ipf

	info, err := cons.MEM.LoadROMFile(*game)
	if err != nil {
		print("unable to load rom ")
		println(*game)
		print("error: ")
		println(err.Error())
		os.Exit(1)
	} else {
		print("loaded ")
		println(info.Name())
	}
}
