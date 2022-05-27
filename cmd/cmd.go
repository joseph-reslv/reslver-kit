package cmd

import (
	"io"
	"os"

	"git.k8s.app/joseph/reslver-kit/types"
	"github.com/urfave/cli/v2"
)

var VERSION string
var DEFAULT_CONFIG_PATH string

func GetCmd(root string, init types.InitFunc, apply types.ApplyFunc) (*cli.App) {
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
		Name: "reslver-kit",
		Usage: "Reslver Toolkit",
		Version: VERSION,
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "init",
				Usage: "initialize reslver toolkit",
				Action: func(ctx *cli.Context) error {
					return init(commands, root)
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
						Name: "config",
						Aliases: []string{"c"},
						Usage: "Load configuration from `DIR`",
						Destination: &commands.ConfigsPath,
						Value: root + DEFAULT_CONFIG_PATH,
						EnvVars: []string{"RESLVER_PATH"},
					},
					&cli.BoolFlag{
						Name: "force",
						Aliases: []string{"f"},
						Usage: "Force initialize all reslver configurations",
						Destination: &commands.Force,
						Value: false,
					},
					&cli.StringFlag{
						Name: "template",
						Aliases: []string{"t"},
						Usage: "Indicate which default YAML configuration template should be generated [ overall | level2 ]",
						Destination: &commands.Template,
						Value: "",
					},
				},
			},
			{
				Name: "apply",
				Usage: "generate diagrams from terraform states",
				Action: func(ctx *cli.Context) error {
					return apply(commands, root)
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
						Value: root + "graph.yaml",
					},
					&cli.StringFlag{
						Name: "config",
						Aliases: []string{"c"},
						Usage: "Load configurations from `DIR`",
						Destination: &commands.ConfigsPath,
						Value: root + DEFAULT_CONFIG_PATH,
						EnvVars: []string{"RESLVER_PATH"},
					},
					&cli.StringFlag{
						Name: "input",
						Aliases: []string{"i"},
						Usage: "Load terraform states from `DIR`",
						Destination: &commands.InputPath,
						Value: root,
					},
					&cli.StringFlag{
						Name: "output",
						Aliases: []string{"o"},
						Usage: "Output results to `DIR`",
						Destination: &commands.OutputPath,
						Value: root,
					},
				},
			},
		},
	}
	return app
}