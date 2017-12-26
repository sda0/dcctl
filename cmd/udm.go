package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var udmCmd = &cobra.Command{
	Use:   "udm",
	Short: "Manage UDM daemon inside cli container",
	Long: `Check status of unregister-devices-manager in cli container and start/stop commands

	dcctl udm 		Status of UDM
	dcctl udm start		Start UDM
	dcctl udm stop		Stop UDM`,

	Run: func(cmd *cobra.Command, args []string) {

		command := "container=$(docker ps -f name=_cli_1 -q | head -n1); docker exec -t $container supervisorctl status | grep unregister-devices-manager"
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(udmCmd)
}
