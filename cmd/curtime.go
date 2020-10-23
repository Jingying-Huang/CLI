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
	"time"

	"github.com/spf13/cobra"
)

type color string

const (
	colorBlack  color = "\u001b[30m"
	colorRed          = "\u001b[31m"
	colorGreen        = "\u001b[32m"
	colorYellow       = "\u001b[33m"
	colorBlue         = "\u001b[34m"
	colorReset        = "\u001b[0m"
)

// CurtimeCmd represents the curtime command
var curtimeCmd = &cobra.Command{
	Use: "curtime",
	RunE: func(cmd *cobra.Command, args []string) error {
		now := time.Now()
		prettyTime := now.Format(time.RubyDate)
		cmd.Println("Hey Gophers! The current time is", prettyTime)
		return nil
	},
}

// func colorize(color color, str string) {
// 	fmt.Println(string(color), str, string(colorReset))
// }

// func init() {
// 	cobra.OnInitialize(initConfig)
// 	curtimeCmd.PersistentFlags().Bool("color", false, "display colorized output")

// }

// func initConfig() {
// 	clr, _ := curtimeCmd.Flags().GetString("curtime")
// 	// if clr != "" {
// 	// 	colorize(colorGreen, "hello world")
// 	// 	return
// 	// }
// }
