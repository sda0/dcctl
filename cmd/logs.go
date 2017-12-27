package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.corp.pushwoosh.com/Backend/pwctl/powodock"
	"log"
	"os/exec"
	"strings"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "View output from containers.",
	Long:  `View output from containers.

Usage: logs [options] [SERVICE...]

Options:
    --no-color          Produce monochrome output.
    -f, --follow        Follow log output.
    -t, --timestamps    Show timestamps.
    --tail="all"        Number of lines to show from the end of the logs
                        for each container.
`,
	Run: func(cmd *cobra.Command, args []string) {
		composeFiles, _ := powodock.GetComposeFilesAndServicesByArg([]string{"all"})
		command := "docker-compose " + composeFiles + " logs " + strings.Join(args, " ")
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).CombinedOutput()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
