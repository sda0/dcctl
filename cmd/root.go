package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/sda0/dcctl/docker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os/exec"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dcctl",
	Short: "Docker-compose control pedal",
	Long:  `Manage your docker compose.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		dockDir := viper.GetString("project_dock")
		if dockDir != "" {
			if err := os.Chdir(dockDir); err != nil {
				log.Fatal(err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		composeFiles, _ := docker.GetComposeFilesAndServicesByArg([]string{"all"})
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

	var dockDir string
	var srcDir string
	var ymlPattern string
	var defaultComposeFile string

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dcctl.yaml)")
	rootCmd.PersistentFlags().StringVar(&dockDir, "project_dock", "", "docker-compose directory path (default is current dir)")
	rootCmd.PersistentFlags().StringVar(&srcDir, "project", "", "project source files directory path")
	rootCmd.PersistentFlags().StringVar(&ymlPattern, "composer_pattern", docker.YmlPattern, "docker-compose files pattern")
	rootCmd.PersistentFlags().StringVar(&defaultComposeFile, "composer_default", docker.DefaultComposeFile, "default docker-compose file name")

	viper.BindPFlag("project_dock", rootCmd.PersistentFlags().Lookup("project_dock"))
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
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

		// Search config in home directory with name ".dcctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dcctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
