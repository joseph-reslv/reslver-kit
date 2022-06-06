package main

import (
	"embed"
	"os"

	"git.k8s.app/resolve/reslver-kit/cmd"
	"git.k8s.app/resolve/reslver-kit/kit"
	log "git.k8s.app/resolve/reslver-kit/logger"
	"git.k8s.app/resolve/reslver-kit/types"
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

//go:embed sources/reslver-configs
var reslverConfigsFileSystem embed.FS

//go:embed templates
var templatesFileSystem embed.FS

func RunInit(flags *types.CommandFlag, root string) (error) {
	log.SetLogger("Reslver Kit", flags.Debug)
	return kit.Init(flags, root)
}

func RunApply(flags *types.CommandFlag, root string) (error) {
	log.SetLogger("Reslver Kit", flags.Debug)
	err := kit.Build(flags, root)
	if err != nil {
		log.Logger.Printf("ERROR: %s", err.Error())
		if !flags.Debug {
			kit.CleanUpSysFiles()
		}
		return err
	}
	return nil
}

func main() {
	root, _ := os.Getwd()
	root += "/"
	
	cmd.VERSION = VERSION
	cmd.DEFAULT_CONFIG_PATH = DEFAULT_CONFIG_PATH
	// map kit fs to embed fs
	kit.TfLoaderFileSystem = tfLoaderFileSystem
	kit.ReslverFileSystem = reslverFileSystem
	kit.GraphModuleFileSystem = graphModuleFileSystem
	kit.GraphGeneratorFileSystem = graphGeneratorFileSystem
	kit.ReslverConfigsFileSystem = reslverConfigsFileSystem
	kit.TemplatesFileSystem = templatesFileSystem
	kit.KIT_ROOT = KIT_ROOT
	kit.CONFIGS_REPO = CONFIGS_REPO
	kit.DEFAULT_CONFIG_PATH = DEFAULT_CONFIG_PATH

	cli := cmd.GetCmd(root, RunInit, RunApply)
	err := cli.Run(os.Args)
	if err != nil {
		os.Exit(2)
	}
}