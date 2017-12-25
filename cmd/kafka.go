package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Manage raker-kafka daemon inside cli",
	Long: `Check status of raker-kafka and start/stop commands

	pwctl kafka 			Show raker-kafka:* status
	pwctl kafka start		Start raker-kafka:*
	pwctl kafka stop		Stop raker-kafka:*`,

	Run: func(cmd *cobra.Command, args []string) {

		command := "docker exec -t \"dock_cli_1\" supervisorctl status | grep raker-kafka"
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
