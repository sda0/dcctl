package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.corp.pushwoosh.com/Backend/pwctl/powodock"
	"log"
	"path/filepath"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "pwctl",
	Short: "Pushwoosh control utility",
	Long:  `Manage your pushwoosh environment like a pro. It is just wrapper around docker-compose. Better place executable file in your PW dock directory and link it by /usr/bin/pwctl.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		pwDockDir := viper.GetString("pw_dock")
		if pwDockDir != "" {
			if err := os.Chdir(pwDockDir); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	var pwDockDir string
	var pwHomeDir string
	var ymlPattern string
	var defaultComposeFile string

	exe, _ := os.Executable()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")
	rootCmd.PersistentFlags().StringVar(&pwDockDir, "pw_dock", filepath.Dir(exe), "powodock directory path")
	rootCmd.PersistentFlags().StringVar(&pwHomeDir, "pw_home", "", "pw home directory path")
	rootCmd.PersistentFlags().StringVar(&ymlPattern, "composer_pattern", powodock.YmlPattern, "pw home directory path")
	rootCmd.PersistentFlags().StringVar(&defaultComposeFile, "composer_default", powodock.DefaultComposeFile, "pw home directory path")

	viper.BindPFlag("pw_dock", rootCmd.PersistentFlags().Lookup("pw_dock"))
	viper.BindPFlag("pw_home", rootCmd.PersistentFlags().Lookup("pw_home"))
	viper.BindPFlag("composer_pattern", rootCmd.PersistentFlags().Lookup("composer_pattern"))
	viper.BindPFlag("composer_default", rootCmd.PersistentFlags().Lookup("composer_default"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pwctl" (without extension).
		viper.AddConfigPath(home)

		//Find current executable directory
		exe, err := os.Executable()
		if err != nil {
			panic(err)
		}

		viper.AddConfigPath(filepath.Dir(exe))
		viper.SetConfigName(".pwctl")
		viper.SupportedExts = append(viper.SupportedExts, "")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
