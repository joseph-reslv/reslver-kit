package kit

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var GeneratorFilename = "./reslvergraph"

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
				log.Printf("error os.WriteFile error: %v", err)
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

func runGraphGenerator(inputPath, outputPath, yamlPath, sourceCodePath string) (string, error) {
	err := copyFilesFromEmbedFS(".", GraphGeneratorSource, sourceCodePath, &GraphGeneratorFileSystem)
	if err != nil {
		return "", err
	}
	if yamlPath, err = moveYamlConfig(yamlPath, filepath.Dir(inputPath) + "/"); err != nil {
		return "", err
	}
	fmt.Println("RUN: graph generator...")
	cmd := exec.Command(sourceCodePath + GeneratorFilename, "--yaml-config", yamlPath, "--output", outputPath)
	if err = cmd.Run(); err != nil {
		return "", err
	}
	fmt.Println("DONE: graph generated")
	return "", nil
}