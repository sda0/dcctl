package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"../docker"
	"log"
	"os/exec"
	"strings"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up services",
	Long: `Up services. For example:

	dcctl up		Up general services only (web, api, cli, nginx)
	dcctl up <service>	Up <service>
	dcctl up all		Up all services`,

	Run: func(cmd *cobra.Command, args []string) {

		composeFiles, servicesList := docker.GetComposeFilesAndServicesByArg(args)

		if len(servicesList) > 0 && !strings.Contains(servicesList, "nginx") {
			servicesList += " nginx"
		}

		command := "docker-compose " + composeFiles + " up -d " + servicesList
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
