package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.corp.pushwoosh.com/Backend/pwctl/powodock"
	"log"
	"os/exec"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "pwctl",
	Short: "Pushwoosh control pedal",
	Long:  `Manage your pushwoosh environment like a pro.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		pwDockDir := viper.GetString("pw_dock")
		if pwDockDir != "" {
			if err := os.Chdir(pwDockDir); err != nil {
				log.Fatal(err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		composeFiles, _ := powodock.GetComposeFilesAndServicesByArg([]string{"all"})
		command := "docker-compose " + composeFiles + " ps"
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwctl.yaml)")
	rootCmd.PersistentFlags().StringVar(&pwDockDir, "pw_dock", "", "powodock directory path (default is current dir)")
	rootCmd.PersistentFlags().StringVar(&pwHomeDir, "pw_home", "", "pw home directory path (default is current dir)")
	rootCmd.PersistentFlags().StringVar(&ymlPattern, "composer_pattern", powodock.YmlPattern, "pw home directory path (default is current dir)")
	rootCmd.PersistentFlags().StringVar(&defaultComposeFile, "composer_default", powodock.DefaultComposeFile, "pw home directory path (default is current dir)")

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
		viper.SetConfigName(".pwctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
