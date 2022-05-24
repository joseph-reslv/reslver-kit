package kit

import (
	"embed"
	"os"

	log "git.k8s.app/joseph/reslver-kit/logger"
)

var TemplatesFileSystem embed.FS
var TemplatesSource = "templates"
var DEFAULT_TEMPLATE = "graph.yaml"

func copyFileFromEmbedFS(fromPath, toPath string, f *embed.FS) (error) {
	content, err := f.ReadFile(fromPath)
	if err != nil {
		return err
	}
	if err := os.WriteFile(toPath, content, 0777); err != nil {
		log.DebugLogger.Printf("error os.WriteFile error: %v", err)
		return err
	}
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
	if err := copyFileFromEmbedFS(templateSource, templatePath, &TemplatesFileSystem); err != nil {
		log.Logger.Println("Unabled to create default configuration template")
		os.Exit(0)
		return err
	}
	log.Logger.Println("Initialized yaml configuration template")
	return nil
}