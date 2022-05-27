package kit

import (
	"embed"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	log "git.k8s.app/joseph/reslver-kit/logger"
)

var TemplatesFileSystem embed.FS
var TemplatesSource = "templates"
var DEFAULT_TEMPLATE = "graph.yaml"

func copyFileFromEmbedFS(fromPath, toPath string, f *embed.FS, replace bool) (error) {
	content, err := f.ReadFile(fromPath)
	if err != nil {
		return err
	}
	counter := 1
	ext := filepath.Ext(toPath)
	path := strings.ReplaceAll(toPath, ext, "")
	for checkIsExist(toPath) {
		toPath = path + " (" + strconv.FormatInt(int64(counter), 10) + ")" + ext
		counter += 1
	}
	if err := os.WriteFile(toPath, content, 0777); err != nil {
		log.DebugLogger.Printf("error os.WriteFile error: %v", err)
		return err
	}
	log.DebugLogger.Println("file created: ", toPath)
	return nil
}

func getTemplates(root string, template string, force bool) (error) {
	templatePath := root + DEFAULT_TEMPLATE
	templateSource := toFolder(TemplatesSource) + template + ".yaml"
	log.Logger.Println("Initializing yaml configuration template")
	if checkIsExist(template) {
		if !force {
			log.Logger.Println("Default configuration is found [", template,"], please remove it manually or run with `--force` to force remove the configuration")
			return nil
		} else {
			log.Logger.Println("Removing old default configuration")
			if err := cleanUp(template); err != nil {
				return err
			}
		}
	}
	log.Logger.Println("Creating default configuration")
	if err := copyFileFromEmbedFS(templateSource, templatePath, &TemplatesFileSystem, false); err != nil {
		log.Logger.Println("Unabled to create default configuration template")
		os.Exit(0)
		return err
	}
	log.Logger.Println("Initialized yaml configuration template")
	return nil
}