package kit

import (
	"embed"

	reslver "git.k8s.app/resolve/reslver/core"
	logger "git.k8s.app/resolve/reslver/logger"
)

var ReslverFileSystem embed.FS
var	ReslverSource = "sources/reslver/"

func runReslver(inputPath string, outputPath string, configPath string, debug bool) (string, error) {
	logger.SetLogger("Reslver", debug)
	err := reslver.Build(inputPath, outputPath, debug, ReslverSource, configPath, &ReslverFileSystem)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}