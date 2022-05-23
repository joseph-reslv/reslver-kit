package kit

import (
	"os"
	"path/filepath"

	"git.k8s.app/joseph/reslver-kit/types"
)

var KIT_ROOT string

func getAbsPath(path, root string) (string) {
	result, _ := filepath.Abs(root + path)
	return result
}

func Build(flags *types.CommandFlag, root string) (error) {
	input := getAbsPath(flags.InputPath, root)
	outputTemp := root + KIT_ROOT
	output := getAbsPath(flags.OutputPath, root) + "/"
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

	// run excel exporter
	_, err = runExcel(input + "output.json", output + "output.xlsx", debug)
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