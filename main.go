package main

import (
	"github.com/jfrog/jfrog-cli-core/plugins"
	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-cli-plugins/vault-frog/commands"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "vault-frog"
	app.Description = "Store Artifactory configuration in the Hashi vault and use them when running JFrog CLI commands."
	app.Version = "v1.0.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetStoreCommand(),
		commands.GetUseCommand(),
		commands.GetDeleteCommand(),
	}