/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/lucaskatayama/app-ideas/Projects/1-Beginner/bin2dec"
	"github.com/spf13/cobra"
)

// bin2decCmd represents the bin2dec command
var bin2decCmd = &cobra.Command{
	Use:   "bin2dec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			res, err := bin2dec.Bin2Dec(arg)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				fmt.Println(res)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(bin2decCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bin2decCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bin2decCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
