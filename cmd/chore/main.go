package chore

import (
	"fmt"
	"os"

	"github.com/rskvp/necode/pkg/fs"
	"github.com/rskvp/necode/services/chore/version"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/server"
)

// main entry point for the web server
func main() {
	app := buildCLI()
	err := app.Run(os.Args)
	if err != nil {
		// An unhandled error was returned, wrap it and run it through the default exit code handler. Any errors
		// that make it here should be caught further up the call stack and wrapped with cli.Exit and the proper exit code.
		cli.HandleExitCoder(cli.Exit(fmt.Sprintf("Unexpected error encountered: %v.", err), 9))
	}
}

// buildCLI is the main entry point for the web server
func buildCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "Temporal UI"
	app.Usage = "https://github.com/temporalio/ui"
	app.Version = version.UIVersion
	app.ArgsUsage = " "
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "root",
			Aliases: []string{"r"},
			Value:   ".",
			Usage:   "root directory of execution environment",
			EnvVars: []string{fs.EnvKeyRoot},
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "config",
			Usage:   "config dir path relative to root",
			EnvVars: []string{fs.EnvKeyConfigDir},
		},
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Value:   "development",
			Usage:   "runtime environment",
			EnvVars: []string{fs.EnvKeyEnvironment},
		}}

	app.Commands = []*cli.Command{
		{
			Name:      "start",
			Usage:     "Start Web UI server",
			ArgsUsage: " ",
			Flags:     []cli.Flag{},
			Action: func(c *cli.Context) error {
				//env := c.String("env")
				//configDir := path.Join(c.String("root"), c.String("config"))
				//cfgProvider := fs_config_provider.NewFSConfigProvider(configDir, env)

				//cfg, err := cfgProvider.GetConfig()
				//if err != nil {
				//	return cli.Exit(err, 1)
				//}

				//opts := []server_options.ServerOption{
				//	server_options.WithConfigProvider(cfgProvider),
				//	server_options.WithAPIMiddleware([]api.Middleware{
				//			headers.WithForwardHeaders(cfg.ForwardHeaders),
				//		}),
				//}

				s := server.NewServer()
				defer s.Stop()
				//err = s.Start()

				//if err != nil {
				return cli.Exit(fmt.Sprintf("Unable to start server"), 1)
				//)
				//cli.Exit("All services are stopped.", 0)
			},
		},
	}
	return app
}
