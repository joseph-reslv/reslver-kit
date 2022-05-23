package cmd

import (
	"io"
	"os"

	"git.k8s.app/joseph/reslver-kit/types"
	"github.com/urfave/cli/v2"
)

var VERSION string
var DEFAULT_CONFIG_PATH string

func GetCmd() (*cli.App, *types.CommandFlag) {
	root, _ := os.Getwd()
	commands := &types.CommandFlag{}
	// workaround for exit program when help flag is ON
	oldHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
			oldHelpPrinter(w, templ, data)
			os.Exit(0)
	}
	oldVersionPrinter := cli.VersionPrinter
	cli.VersionPrinter = func(c *cli.Context) {
			oldVersionPrinter(c)
			os.Exit(0)
	}
	//
	
	app := &cli.App{
		HideHelpCommand: true,
		Name: "reslver-kit",
		Usage: "generate diagrams from terraform states",
		Version: VERSION,
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "debug",
				Aliases: []string{"d"},
				Usage: "Enable debug mode",
				Destination: &commands.Debug,
				Value: false,
			},
			&cli.StringFlag{
				Name: "yaml-config",
				Aliases: []string{"y"},
				Usage: "Load graph YAML configuration from `FILE`",
				Destination: &commands.ConfigYAML,
				Value: root + "/" + "graph.yaml",
			},
			&cli.StringFlag{
				Name: "config",
				Aliases: []string{"c"},
				Usage: "Load configuration from `DIR`",
				Destination: &commands.ConfigsPath,
				Value: root + "/" + DEFAULT_CONFIG_PATH,
				EnvVars: []string{"RESLVER_PATH"},
			},
			&cli.StringFlag{
				Name: "input",
				Aliases: []string{"i"},
				Usage: "Load components from `DIR`",
				Destination: &commands.InputPath,
				Value: root ,
			},
			&cli.StringFlag{
				Name: "output",
				Aliases: []string{"o"},
				Usage: "Output results to `DIR`",
				Destination: &commands.OutputPath,
				Value: root,
			},
		},
	}
	return app, commands
}