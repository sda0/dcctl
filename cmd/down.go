package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.corp.pushwoosh.com/Backend/pwctl/powodock"
	"log"
	"os/exec"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down services",
	Long: `Down one, few or all services listed in powodock compose files. For example:

	pwctl down		Stops all launched services
	pwctl down <service>	Stop <service>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = []string{"all"}
		}
		composeFiles, servicesList := powodock.GetComposeFilesAndServicesByArg(args)

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
