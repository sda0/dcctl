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

	dcctl kafka 			Show raker-kafka:* status
	dcctl kafka start		Start raker-kafka:*
	dcctl kafka stop		Stop raker-kafka:*`,

	Run: func(cmd *cobra.Command, args []string) {

		command := "container=$(docker ps -f name=_cli_1 -q | head -n1); docker exec -t $container supervisorctl status | grep raker-kafka"
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
