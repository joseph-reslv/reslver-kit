package kit

import (
	"embed"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "git.k8s.app/joseph/reslver-kit/logger"
	"gopkg.in/yaml.v2"
)

var GeneratorFilename = "./reslvergraph"
var GraphGeneratorFileSystem embed.FS
var	GraphGeneratorSource = "sources/reslver-static-graph-exporter/"

type YamlConfig struct {
	Diagram string `yaml:"diagram"`
}

func readYaml(yamlPath string) (YamlConfig, error) {
	var config YamlConfig
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func getYamlDiagram(yamlPath string) (string, error) {
	config, err := readYaml(yamlPath)
	if err != nil {
		return "", err
	}
	return config.Diagram, nil
}

func getYamlConfig(yamlPath string) (string, error) {
	fileInfo, err := os.Stat(yamlPath)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		if _, err := os.Stat(yamlPath); err != nil {
			return "", err
		}
		return yamlPath, nil
	}
	return yamlPath, nil
}

func copyFilesFromEmbedFS(rootPath, fromPath, toPath string, f *embed.FS) (error) {
	err := fs.WalkDir(f, rootPath, func(s string, d fs.DirEntry, e error) error {
		if e != nil { return e }
		if !d.IsDir() {
			content, err := f.ReadFile(s)
			if err != nil {
				return err
			}
			path := toPath + strings.ReplaceAll(s, fromPath, "")
			if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
        return err
    	}
			if err := os.WriteFile(path, content, 0777); err != nil {
				log.DebugLogger.Printf("error os.WriteFile error: %v", err)
				return err
			}
		}
		return nil
 	})
	if err != nil {
		return err
	}
	return nil
}

func moveYamlConfig(yamlPath, toPath string) (string, error) {
	content, err := os.ReadFile(yamlPath)
	if err != nil {
		return "", err
	}
	path := toPath + filepath.Base(yamlPath)
	if err := os.WriteFile(path, content, 0777); err != nil {
		return "", err
	}
	return path, nil
}

func unzipGraphGenerator(dirPath string, filePath string) (error) {
	filePath += ".zip"
	log.DebugLogger.Printf("Zips graph generator with [unzip, %s] in [%s]", filePath, dirPath)
	cmd := exec.Command("unzip", filePath)
	cmd.Dir = dirPath
	if err := cmd.Run(); err != nil {
		return errors.New("unable to unzip graph generator")
	}
	return nil
}

func runGraphGenerator(inputPath, outputPath, yamlPath, sourceCodePath string) (string, error) {
	err := copyFilesFromEmbedFS(".", GraphGeneratorSource, sourceCodePath, &GraphGeneratorFileSystem)
	if err != nil {
		return "", err
	}
	if yamlPath, err = moveYamlConfig(yamlPath, filepath.Dir(inputPath) + "/"); err != nil {
		return "", err
	}
	log.Logger.Println("Running graph generator")
	if err = unzipGraphGenerator(sourceCodePath, GeneratorFilename); err != nil {
		return "", err
	}
	log.DebugLogger.Printf("Graph generator runs with [%s, --yaml-config, %s, --output, %s]", sourceCodePath + GeneratorFilename, yamlPath, outputPath)
	cmd := exec.Command(sourceCodePath + GeneratorFilename, "--yaml-config", yamlPath, "--output", outputPath)
	if err = cmd.Run(); err != nil {
		return "", errors.New("unable to run graph generator")
	}
	log.Logger.Println("Graph generator run successfully")
	return "", nil
}