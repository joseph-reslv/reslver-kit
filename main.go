package main

import (
	"embed"
	"fmt"
	"os"

	"git.k8s.app/joseph/reslver-kit/cmd"
	"git.k8s.app/joseph/reslver-kit/kit"
)

//go:embed sources/reslver-tf-loader
var tfLoaderFileSystem embed.FS

//go:embed sources/reslver-graph-exporter
var graphModuleFileSystem embed.FS

//go:embed sources/reslver
var reslverFileSystem embed.FS

func main() {
	root, _ := os.Getwd()
	root += "/"
	cli, flags := cmd.GetCmd()
	err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	// map kit fs to embed fs
	kit.TfLoaderFileSystem = tfLoaderFileSystem
	kit.ReslverFileSystem = reslverFileSystem
	kit.GraphModuleFileSystem = graphModuleFileSystem
	err = kit.Build(flags, root)
	if err != nil {
		fmt.Println(err)
		return
	}
}