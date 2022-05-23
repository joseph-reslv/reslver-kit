package kit

import (
	excel "git.k8s.app/joseph/reslver-excel-exporter/core"
	logger "git.k8s.app/joseph/reslver-excel-exporter/logger"
)

func runExcel(inputPath string, outputPath string, debug bool) (string, error) {
	logger.SetLogger("Excel Exporter", debug)
	err := excel.Build(inputPath, outputPath)
	if err != nil {
		return "", err
	}
	return outputPath, nil
}