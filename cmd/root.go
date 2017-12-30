package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/kardianos/osext"
	"./docker"
	"log"
	"path/filepath"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dcctl",
	Short: "Docker-compose control utility",
	Long:  `Manage your docker compose.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		dockDir := viper.GetString("project_dock")
		if dockDir != "" {
			if err := os.Chdir(dockDir); err != nil {
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

	var dockDir string
	var srcDir string
	var ymlPattern string
	var defaultComposeFile string

	path := getRealPath()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dcctl.yaml)")
	rootCmd.PersistentFlags().StringVar(&dockDir, "project_dock", path, "docker-compose directory path (default is current dir)")
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

		// Search config in home directory with name ".pwctl" (without extension).
		viper.AddConfigPath(home)

		//Find current executable directory
		path := getRealPath()

		viper.AddConfigPath(path)
		viper.SetConfigName(".dcctl")
		viper.SupportedExts = append(viper.SupportedExts, "")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getRealPath() (string)  {
	exe, err := osext.Executable()
	if err != nil {
		panic(err)
	}

	ln, err := filepath.EvalSymlinks(exe)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ln)
}