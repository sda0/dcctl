package cmd

import (
	"fmt"
	"github.com/sda0/dcctl/docker"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start services",
	Long: `Start services general services

	dcctl start		Start general service
	dcctl start <service>	Start <service>
	dcctl start all		Start all services`,

	Run: func(cmd *cobra.Command, args []string) {

		composeFiles, servicesList := docker.GetComposeFilesAndServicesByArg(args)

		if len(servicesList) > 0 && !strings.Contains(servicesList, "nginx") {
			servicesList += " nginx"
		}

		command := "docker-compose " + composeFiles + " start " + servicesList
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
