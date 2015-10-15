package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/jberkhahn/v3_beta/commands"
)

type V3Plugin struct {
}

func main() {
	plugin.Start(new(V3Plugin))
}

func (v3plugin *V3Plugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "v3-push" {
		commands.Push(cliConnection, args)
	} else if args[0] == "v3-apps" {
		if len(args) == 1 {
			commands.Apps(cliConnection, args)
		} else {
			//print apps help
		}
	} else if args[0] == "v3-processes" {
		if len(args) == 1 {
			commands.Processes(cliConnection, args)
		} else {
			//print processes help
		}
	} else if args[0] == "v3-delete" {
		if len(args) == 2 {
			commands.Delete(cliConnection, args)
		} else {
			//print processes help
		}
	}
}

func (v3plugin *V3Plugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "v3_beta",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 4,
			Build: 4,
		},
		Commands: []plugin.Command{
			{
				Name:     "v3-push",
				Alias:    "",
				HelpText: "pushes current dir as a v3 process",
				UsageDetails: plugin.Usage{
					Usage: "v3-push APPNAME",
					Options: map[string]string{
						"p": "path to dir or zip to push",
					},
				},
			},
			{
				Name:     "v3-apps",
				Alias:    "v3-a",
				HelpText: "displays all v3 apps",
				UsageDetails: plugin.Usage{
					Usage:   "v3-apps",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-processes",
				Alias:    "v3-p",
				HelpText: "displays all v3 processes",
				UsageDetails: plugin.Usage{
					Usage:   "v3-processes",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-delete",
				Alias:    "v3-d",
				HelpText: "delete a v3 app",
				UsageDetails: plugin.Usage{
					Usage:   "v3-delete APPNAME",
					Options: map[string]string{},
				},
			},
		},
	}
}
