package cmd

import (
	"playbook/cmd/check"
	"playbook/pkg/output"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/component-base/logs"
	"os"
	"runtime"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "playbook",
	Short: "playbook protection levels",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		printInfo()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		output.PrintCliError(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.playbook.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(check.NewCmdCheck())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			output.PrintCliError(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".playbook" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".playbook")
	}

	viper.Set("Verbose", false)
	//viper.Set("LogFile", LogFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		output.PrintCliInfo(fmt.Sprintf("Using config file:", viper.ConfigFileUsed()))
	}
}

func printInfo() {
	output.PrintCliInfo(fmt.Sprintf("%s\n%s\n%s\n%s",
		fmt.Sprintf("Operating System: %s\nArchitecture: %s", runtime.GOOS, runtime.GOARCH),
		fmt.Sprint("Check our Sources at https://github.com/Continuous-X/continuousSec"),
		fmt.Sprintf("Get in contact via github issue....."),
		fmt.Sprintf("Author: %s", viper.GetString("author")),
	))

}
