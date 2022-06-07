package kit

import (
	"embed"

	graphExporter "git.k8s.app/resolve/reslver-graph-exporter/core"
	logger "git.k8s.app/resolve/reslver-graph-exporter/logger"
	graphExporterUtils "git.k8s.app/resolve/reslver-graph-exporter/utils"
)

var GraphModuleFileSystem embed.FS
var GraphModuleSource = "sources/reslver-graph-exporter/"

func runGraphExporter(profile string, inputPath string, outputPath string, configPath string, root string, debug bool) (string, error) {
	logger.SetLogger("Graph Exporter", debug)
	profilePath, err := graphExporterUtils.GetProfilePath(profile, configPath, root)
	if err != nil {
		return "", err
	}
	err = graphExporter.Build(profilePath, inputPath, outputPath, debug, GraphModuleSource, configPath, &GraphModuleFileSystem)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}