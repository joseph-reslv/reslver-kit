package kit

import (
	"os"
	"path/filepath"

	log "git.k8s.app/joseph/reslver-kit/logger"
	"git.k8s.app/joseph/reslver-kit/types"
)

var KIT_ROOT string

func getAbsPath(path, root string) (string) {
	if !filepath.IsAbs(path){
		path, _ = filepath.Abs(root + path)
	}
	return path
}

func createOutputFolder(folderPath string) (error) {
	log.DebugLogger.Println("create output folder: ", folderPath)
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err){
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func CleanUpSysFiles() (error) {
	// remove generated files
	log.DebugLogger.Println("remove reslver kit generated files")
	_, err := os.Stat(KIT_ROOT)
	if !os.IsNotExist(err){
		if err := os.RemoveAll(KIT_ROOT); err != nil {
			log.DebugLogger.Println(err)
			return err
		}
	}
	return nil
}

func Build(flags *types.CommandFlag, root string) (error) {
	// remove generated files
	CleanUpSysFiles()

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
	if err := createOutputFolder(output); err != nil {
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

	CleanUpSysFiles()
	return nil
}