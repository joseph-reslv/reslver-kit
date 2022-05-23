package kit

import (
	"embed"
	"os"
	"path/filepath"

	graphExporter "git.k8s.app/joseph/reslver-graph-exporter/core"
	graphExporterUtils "git.k8s.app/joseph/reslver-graph-exporter/utils"
	"git.k8s.app/joseph/reslver-kit/types"
	tfLoader "git.k8s.app/joseph/reslver-tf-loader/core"
	reslver "git.k8s.app/joseph/reslver/core"
)

var KIT_ROOT string

var TfLoaderFileSystem embed.FS
var GraphModuleFileSystem embed.FS
var ReslverFileSystem embed.FS
var GraphGeneratorFileSystem embed.FS

var TFLoaderSource = "sources/reslver-tf-loader/"
var GraphModuleSource = "sources/reslver-graph-exporter/"
var	ReslverSource = "sources/reslver/"
var	GraphGeneratorSource = "sources/reslver-static-graph-exporter/"

func getAbsPath(path, root string) (string) {
	result, _ := filepath.Abs(root + path)
	return result
}

func runReslver(inputPath string, outputPath string, configPath string, debug bool) (string, error) {
	err := reslver.Build(inputPath, outputPath, debug, ReslverSource, configPath, &ReslverFileSystem)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}

func runTFLoader(inputPath string, outputPath string, configPath string, debug bool) (string, error) {
	err := tfLoader.Build(inputPath, outputPath, debug, TFLoaderSource, configPath, &TfLoaderFileSystem)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}

func runGraphExporter(profile string, inputPath string, outputPath string, configPath string, root string, debug bool) (string, error) {
	profilePath, err := graphExporterUtils.GetProfilePath(profile, configPath, root)
	if err != nil {
		return "", err
	}
	err = graphExporter.Build(profilePath, inputPath, outputPath, debug, GraphModuleSource, configPath, &GraphModuleFileSystem)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}

func Build(flags *types.CommandFlag, root string) (error) {
	input := getAbsPath(flags.InputPath, root)
	outputTemp := root + KIT_ROOT
	output := getAbsPath(flags.OutputPath, root)
	config := flags.ConfigsPath + "/"
	debug := flags.Debug
	yaml, err := getYamlConfig(flags.ConfigYAML)
	if err != nil {
		return err
	}
	profile, err := getYamlDiagram(yaml)
	if err != nil {
		return err
	}
	downloadConfigsFromGit(config)
	
	// run tf loader
	input, err = runTFLoader(input, outputTemp + "_tfloader/", config, debug)
	if err != nil {
		return err
	}

	// run reslver core
	input, err = runReslver(input, outputTemp + "_reslver/", config, debug)
	if err != nil {
		return err
	}

	// run graph exporter
	input, err = runGraphExporter(profile, input + "output.json", outputTemp + "_graphexporter/", config, root, debug)
	if err != nil {
		return err
	}

	// run graph generator
	_ggenInput := input
	_ggenOutput := output
	_ggenYaml := yaml
	_ggenSource := outputTemp + "_graph_generator/"
	_, err = runGraphGenerator(_ggenInput, _ggenOutput, _ggenYaml, _ggenSource)
	if err != nil {
		return err
	}

	// remove generated files
	if err := os.RemoveAll(KIT_ROOT); err != nil {
		return err
	}
	
	return nil
}