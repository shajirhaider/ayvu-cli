// Copyright © 2018 Shajir Haider
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/viper"
	"fmt"

	"github.com/spf13/cobra"
)

// sayhelloCmd represents the sayhello command
var sayhelloCmd = &cobra.Command{
	Use:   "hello",
	Short: " then type -n and name or anything ",
	Long: `name will print if use with -n`,
	Run: func(cmd *cobra.Command, args []string) {
		greeting := "Hello"
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		if viper.GetString("name")!=""{
			name = viper.GetString("name")
		}
		if viper.GetString("greeting")!=""{
			greeting = viper.GetString("greeting")
		}
		fmt.Println(greeting + " " + name)
	},
}

func init() {
	sayCmd.AddCommand(sayhelloCmd)
	sayhelloCmd.Flags().StringP("name", "n", viper.GetString("NAME") , "Set your name")
}
