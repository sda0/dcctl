package cmd

import (
	"fmt"

	"../docker"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down services",
	Long: `Down one, few or all services listed in docker compose files. For example:

	dcctl down		Stops all launched services
	dcctl down <service>	Stop <service>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = []string{"all"}
		}
		composeFiles, servicesList := docker.GetComposeFilesAndServicesByArg(args)

		command := "docker-compose " + composeFiles + " down " + servicesList
		println(command)

		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
