package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var kafkaStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start raker-kafka ",
	Long: `Start raker-kafka inside cli container

This is a wrapper for docker exec -t "dock_cli_1" supervisorctl start raker-kafka:*`,

	Run: func(cmd *cobra.Command, args []string) {

		command := "container=$(docker ps -f name=_cli_1 -q | head -n1); docker exec -t $container supervisorctl start raker-kafka:*"
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	kafkaCmd.AddCommand(kafkaStartCmd)
}
