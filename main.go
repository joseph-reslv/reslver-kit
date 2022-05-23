package main

import (
	"embed"
	"os"

	"git.k8s.app/joseph/reslver-kit/cmd"
	"git.k8s.app/joseph/reslver-kit/kit"
	log "git.k8s.app/joseph/reslver-kit/logger"
)

var VERSION = "1.0.0"
var DEFAULT_CONFIG_PATH = ".reslver/configs/"
var KIT_ROOT = ".reslver_sys/"
var CONFIGS_REPO = "https://git.k8s.app/joseph/reslver-configs"

//go:embed sources/reslver-tf-loader
var tfLoaderFileSystem embed.FS

//go:embed sources/reslver-graph-exporter
var graphModuleFileSystem embed.FS

//go:embed sources/reslver
var reslverFileSystem embed.FS

//go:embed sources/reslver-static-graph-exporter
var graphGeneratorFileSystem embed.FS

func main() {
	root, _ := os.Getwd()
	root += "/"
	
	cmd.VERSION = VERSION
	cmd.DEFAULT_CONFIG_PATH = DEFAULT_CONFIG_PATH
	cli, flags := cmd.GetCmd()
	err := cli.Run(os.Args)
	if err != nil {
		return
	}
	log.SetLogger("Reslver Kit", flags.Debug)

	// map kit fs to embed fs
	kit.TfLoaderFileSystem = tfLoaderFileSystem
	kit.ReslverFileSystem = reslverFileSystem
	kit.GraphModuleFileSystem = graphModuleFileSystem
	kit.GraphGeneratorFileSystem = graphGeneratorFileSystem
	kit.KIT_ROOT = KIT_ROOT
	kit.CONFIGS_REPO = CONFIGS_REPO

	err = kit.Build(flags, root)
	if err != nil {
		log.DebugLogger.Println(err)
		return
	}
}