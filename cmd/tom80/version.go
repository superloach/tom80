package main

import (
	rt_debug "runtime/debug"
	"strings"
)

func version_info() {
	info, ok := rt_debug.ReadBuildInfo()
	if ok {
		println("tom80 version: " + info.Main.Version)
		for _, dep := range info.Deps {
			path := strings.Split(dep.Path, "/")
			auth := path[len(path)-2]
			name := path[len(path)-1]
			if auth == "hajimehoshi" && name == "ebiten" {
				println("ebiten version: " + dep.Version)
			}
		}
	}
}
