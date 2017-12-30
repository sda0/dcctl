package cmd

import (
	"fmt"

	"../docker"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop services",
	Long: `Stop one, few or all services listed in docker compose files. For example:

	dcctl stop		Stops all launched services
	dcctl stop <service>	Stop <service>`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			args = []string{"all"}
		}
		composeFiles, servicesList := docker.GetComposeFilesAndServicesByArg(args)
		command := "docker-compose " + composeFiles + " stop " + servicesList
		println(command)

		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
