// Copyright © 2018 Shajir Haider
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		// c := exec.Command("/home/shajir/go/src/ayvu-cli/hello.py")
		c := exec.Command("python3.6", "pythonfile.py")
		output, err := c.Output()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
