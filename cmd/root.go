/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package cmd

import (
	"os"
	"path/filepath"
	"runtime"

	"a13s.io/pleiades/pkg/configuration"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pleiades",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	config *viper.Viper
)

func init() {
	config = configuration.Get()

	viper.SetConfigName("pleiades") // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name

	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "darwin" {
		dir, _ := homedir.Dir()
		viper.AddConfigPath(filepath.Join(dir, "Library", "pleiades"))
	} else {
		viper.AddConfigPath(configuration.DefaultBaseConfigPath) // path to look for the config file in
	}

	viper.AddConfigPath("$HOME/.pleiades") // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().Bool("trace", false, "enable trace logging")
	config.BindPFlag("trace", rootCmd.PersistentFlags().Lookup("trace"))

	rootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")
	config.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.MarkFlagsMutuallyExclusive("debug", "trace")
}

func setupLogger(cmd *cobra.Command, args []string) zerolog.Logger {
	var logger zerolog.Logger
	if config.GetBool("trace") {
		logger = configuration.NewRootLogger().Level(zerolog.TraceLevel)
	} else if config.GetBool("debug") {
		logger = configuration.NewRootLogger().Level(zerolog.DebugLevel)
	} else {
		logger = configuration.NewRootLogger().Level(zerolog.InfoLevel)
	}
	return logger
}
