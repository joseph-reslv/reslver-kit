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

func checkIsExist(configsPath string) (bool) {
	if _, err := os.Stat(configsPath); os.IsNotExist(err) {
		return false	
	}
	return true
}

func downloadConfigsFromGit(configsPath string, force bool) (error) {
	log.Logger.Println("Initializing reslver configurations")
	if checkIsExist(configsPath) {
		if !force {
			log.Logger.Println("Configurations folder is found [", configsPath,"], please remove it manually or run with `--force` to force clear the folder")
			return nil
		} else {
			log.Logger.Println("Clearing old reslver configurations")
			if err := cleanUp(configsPath); err != nil {
				return err
			}
		}
	}
	log.Logger.Println("Downloading reslver configurations")
	_, err := git.PlainClone(configsPath, false, &git.CloneOptions{
		URL: CONFIGS_REPO,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Logger.Println("Unabled to clone reslver configuration, please try clone the configuration manually:", CONFIGS_REPO)
		log.Logger.Println("Copying default configurations instead")
		if err := copyFilesFromEmbedFS(".", ReslverConfigsSource, configsPath, &ReslverConfigsFileSystem); err != nil {
			log.Logger.Println("Unabled to copy default configurations")
			os.Exit(0)
			return err
		}
	}
	log.Logger.Println("Initialized reslver configurations")
	return nil
}