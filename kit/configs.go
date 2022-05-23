package kit

import (
	"embed"
	"os"

	log "git.k8s.app/joseph/reslver-kit/logger"
	"github.com/go-git/go-git/v5"
)

var CONFIGS_REPO string
var ReslverConfigsFileSystem embed.FS
var ReslverConfigsSource = "sources/reslver-configs/"

func checkConfigsFolderIsExist(configsPath string) (bool) {
	if _, err := os.Stat(configsPath); os.IsNotExist(err) {
		return false	
	}
	return true
}

func downloadConfigsFromGit(configsPath string) (error) {
	if !checkConfigsFolderIsExist(configsPath) {
		log.Logger.Println("Copying reslver configurations")
		log.DebugLogger.Println("download reslver configurations")
		_, err := git.PlainClone(configsPath, false, &git.CloneOptions{
			URL: CONFIGS_REPO,
			Progress: os.Stdout,
		})
		if err != nil {
			log.DebugLogger.Println("Unabled to clone reslver configuration, please try clone the configuration manually:", CONFIGS_REPO)
			log.DebugLogger.Println("Copy default configurations instead")
			if err := copyFilesFromEmbedFS(".", ReslverConfigsSource, configsPath, &ReslverConfigsFileSystem); err != nil {
				log.DebugLogger.Println("Unabled to copy default configurations, kill process")
				os.Exit(0)
				return err
			}
		}
		log.Logger.Println("Copied reslver configurations")
	}
	return nil
}