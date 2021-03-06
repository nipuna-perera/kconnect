/*
Copyright 2020 The kconnect Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"flag"

	"github.com/fidelity/kconnect/internal/commands/use"
	"github.com/fidelity/kconnect/internal/commands/version"

	home "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	configFile  string
	logLevel    string
	logFormat   string
	interactive bool
)

// RootCmd creates the root kconnect command
func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "kconnect",
		Short: "The Kubernetes Connection Manager CLI",
		Run: func(c *cobra.Command, _ []string) {
			if err := c.Help(); err != nil {
				log.Debugf("ignoring cobra error %q", err.Error())
			}
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Configuration file (defaults to $HOME/.kconnect/config")
	//TODO: change the below back to INFO
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", log.DebugLevel.String(), "Log level for the CLI. Defaults to INFO")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "TEXT", "Format of the log output. Defaults to text.")
	rootCmd.PersistentFlags().BoolVar(&interactive, "non-interactive", false, "Run without interactive flag resolution. Defaults to false")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	rootCmd.AddCommand(use.Command())
	rootCmd.AddCommand(version.Command())

	cobra.OnInitialize(initConfig)

	// err := logging.Configure(logLevel, logFormat)
	// if err != nil {
	// 	return fmt.Errorf("configuring logging: %w", err)
	// }

	// if err := rootCmd.Execute(); err != nil {
	// 	return fmt.Errorf("executing root command: %w", err)
	// }

	return rootCmd
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := home.Dir()
		if err != nil {
			panic(err)
		}

		//TODO: construct path properyl
		viper.AddConfigPath(home)
		viper.SetConfigName(".kconnect")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
