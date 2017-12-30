package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.corp.pushwoosh.com/Backend/pwctl/powodock"
	"log"
	"os/exec"
	"fmt"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Long:  `List containers status`,
	Run: func(cmd *cobra.Command, args []string) {
		composeFiles, _ := powodock.GetComposeFilesAndServicesByArg([]string{"all"})
		command := "docker-compose " + composeFiles + " ps"
		println(command)
		stdoutStderr, err := exec.Command("sh", "-c", command).Output()
		fmt.Printf("%s", stdoutStderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
}
