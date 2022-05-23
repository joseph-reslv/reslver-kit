package kit

import (
	"embed"

	tfLoader "git.k8s.app/joseph/reslver-tf-loader/core"
	logger "git.k8s.app/joseph/reslver-tf-loader/logger"
)

var TfLoaderFileSystem embed.FS
var TFLoaderSource = "sources/reslver-tf-loader/"

func runTFLoader(inputPath string, outputPath string, configPath string, debug bool) (string, error) {
	logger.SetLogger("Terraform Loader", debug)
	err := tfLoader.Build(inputPath, outputPath, debug, TFLoaderSource, configPath, &TfLoaderFileSystem)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}
