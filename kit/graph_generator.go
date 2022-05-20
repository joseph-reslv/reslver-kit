package kit

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var DefaultYamlFilename = "graph.yaml"

func getYamlConfig(yamlPath string) (string, error) {
	fileInfo, err := os.Stat(yamlPath)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		if _, err := os.Stat(yamlPath + "/" + DefaultYamlFilename); err != nil {
			return "", err
		}
		return yamlPath + DefaultYamlFilename, nil
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
	cmd := exec.Command("python3", sourceCodePath, "--yaml-config", yamlPath, "--output", outputPath)
	if err = cmd.Run(); err != nil {
		return "", err
	}
	return "", nil
}