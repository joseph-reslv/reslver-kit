package kit

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

var CONFIGS_REPO string

func checkConfigsFolderIsExist(configsPath string) (bool) {
	if _, err := os.Stat(configsPath); os.IsNotExist(err) {
		return false	
	}
	return true
}

func downloadConfigsFromGit(configsPath string) (error) {
	if !checkConfigsFolderIsExist(configsPath) {
		log.Println("download reslver configurations")
		_, err := git.PlainClone(configsPath, false, &git.CloneOptions{
			URL: CONFIGS_REPO,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println("Unabled to clone reslver configuration, please try clone the configuration manually:", CONFIGS_REPO)
			os.Exit(0)
			return err
		}
	}
	return nil
}