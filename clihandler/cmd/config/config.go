package config

import (
	"github.com/spf13/cobra"
)

var (
	configExample = `
  # View cached configurations 
  kubescape config view

  # Delete cached configurations
  kubescape config delete

  # Set cached configurations
  kubescape config set --help
`
	setConfigExample = `
  # Set account id
  kubescape config set accountID <account id>

  # Set client id
  kubescape config set clientID <client id> 

  # Set access key
  kubescape config set secretKey <access key>
`
)

func GetConfigCmd() *cobra.Command {

	// configCmd represents the config command
	configCmd := &cobra.Command{
		Use:     "config",
		Short:   "handle cached configurations",
		Example: configExample,
	}

	configCmd.AddCommand(getDeleteCmd())
	configCmd.AddCommand(getSetCmd())
	configCmd.AddCommand(getViewCmd())

	return configCmd
}
