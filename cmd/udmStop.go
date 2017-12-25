package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var udmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop unregister device manager",
	Long: `Stop unregister device manager inside cli container

This is a wrapper for docker exec -t "dock_cli_1" supervisorctl stop unregister-devices-manager:*`,


	Run: func(cmd *cobra.Command, args []string) {

		command := "container=$(docker ps -f name=_cli_1 -q | head -n1); docker exec -t $container supervisorctl stop unregister-devices-manager:*"
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	udmCmd.AddCommand(udmStopCmd)
}
