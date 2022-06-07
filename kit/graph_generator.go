package kit

import (
	"archive/tar"
	"compress/gzip"
	"embed"
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "git.k8s.app/resolve/reslver-kit/logger"
	"gopkg.in/yaml.v2"
)

var ZipFilename = "reslvergraph"
var ExecFilename = "reslvergraph.py"
var GraphGeneratorFileSystem embed.FS
var	GraphGeneratorSource = "sources/reslver-static-graph-generator/"

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

/* utils */
func ExtractTarGz(tarReader *tar.Reader, path string) (error) {
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
				break
		}
		if header.Name == "./" || header.Name == "../" {
			continue
		}

		if err != nil {
				log.DebugLogger.Printf("ExtractTarGz: Next() failed: %s", err.Error())
				return err
		}

		switch header.Typeflag {
			case tar.TypeDir:
				if err := os.Mkdir(path + header.Name, 0755); err != nil {
						log.DebugLogger.Printf("ExtractTarGz: Mkdir() failed: %s", err.Error())
						return err
				}
			case tar.TypeReg:
				outFile, err := os.Create(path + header.Name)
				if err != nil {
						log.DebugLogger.Printf("ExtractTarGz: Create() failed: %s", err.Error())
						return err
				}
				if _, err := io.Copy(outFile, tarReader); err != nil {
						log.DebugLogger.Printf("ExtractTarGz: Copy() failed: %s", err.Error())
						return err
				}
				outFile.Close()

			default:
				log.DebugLogger.Printf(
						"ExtractTarGz: uknown type: %s in %s",
						string(header.Typeflag),
						string(path + header.Name))
						return err
		}
	}
	return nil
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
	filePath = dirPath + filePath + ".tar.gz"
	log.DebugLogger.Printf("extracting gzip tar graph generator")
	file, err := os.Open(filePath)
	if err != nil {
		log.DebugLogger.Println(err)
		return errors.New("unable to unzip graph generator")
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		log.DebugLogger.Println(err)
		return errors.New("unable to unzip graph generator")
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	ExtractTarGz(tr, dirPath)
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
	if err = unzipGraphGenerator(sourceCodePath, ZipFilename); err != nil {
		return "", err
	}
	log.DebugLogger.Printf("Graph generator runs with [python3, %s, --yaml-config, %s, --output, %s]", sourceCodePath + ExecFilename, yamlPath, outputPath)
	cmd := exec.Command("python3", sourceCodePath + ExecFilename, "--yaml-config", yamlPath, "--output", outputPath)
	out, err := cmd.Output()
	if err != nil {
		return "", errors.New("unable to run graph generator")
	}
	msg := string(out)
	if strings.Contains(msg, "error") {
		log.DebugLogger.Println(msg)
		return "", errors.New("graph generator got error, please check YAML configuration whether is correct")
	}
	log.Logger.Println("Graph generator run successfully")
	return "", nil
}