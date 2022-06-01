package kit

import (
	"errors"
	"os"
	"path/filepath"

	log "git.k8s.app/resolve/reslver-kit/logger"
	"git.k8s.app/resolve/reslver-kit/types"
)

var KIT_ROOT string

func toFolder(path string) (string) {
	if path[len(path)-1:] != "/" {
		return path + "/"
	}
	return path
}

func checkConfigsFolder(configsPath string) (error) {
	if !checkIsExist(configsPath) {
		log.Logger.Println("Unabled to get reslver configurations, please try to run `reslver-kit init` to initialize reslver")
		return errors.New("invalid configuration folder")
	}
	return nil
}

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

func cleanUp(path string) (error) {
	_, err := os.Stat(path)
	if !os.IsNotExist(err){
		if err := os.RemoveAll(path); err != nil {
			log.DebugLogger.Println(err)
			return err
		}
	}
	return nil
}

func CleanUpSysFiles() (error) {
	return cleanUp(KIT_ROOT)
}

func Build(flags *types.CommandFlag, root string) (error) {
	// remove generated files
	CleanUpSysFiles()
	input := toFolder(getAbsPath(flags.InputPath, root))
	outputTemp := root + KIT_ROOT
	output := toFolder(getAbsPath(flags.OutputPath, root))
	config := toFolder(flags.ConfigsPath)
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
	if err := checkConfigsFolder(config); err != nil {
		return err
	}
	
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
	log.Logger.Println("All results generated in [", output, "]")

	if !flags.Debug {
		return CleanUpSysFiles()
	}
	return nil
}

func Init(flags *types.CommandFlag, root string) (error) {
	config := toFolder(flags.ConfigsPath)
	err := downloadConfigsFromGit(config, flags.Force)
	if err != nil {
		return err
	}
	if flags.Template != "" {
		if err := getTemplates(root, flags.Template, flags.Force); err != nil {
			return err
		}
	}
	return nil
}