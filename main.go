/*
Copyright © 2021 yuansudong

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"log"

	"github.com/spf13/cobra"
)

// GetClientCommand
func GetClientCommnd() *cobra.Command {

	mRootCommand := cobra.Command{
		Use:   "client",
		Long:  "client-long",
		Short: "client-short",
	}
	mUpdateGD := &cobra.Command{
		Use:   "update_gd",
		Long:  "update_gd_long",
		Short: "update_gd_short",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(cmd.PersistentFlags().GetString("listen_addr"))
			log.Println("call client update_gd", args, cmd.Example)
		},
	}
	mUpdateGD.Flags().String("listen_addr", "127.0.0.1", "--listen_addr=:8080")
	mRootCommand.AddCommand(mUpdateGD)

	return &mRootCommand
}

func main() {
	mRootCmd := cobra.Command{Use: "docker", Long: "虚拟化引擎-long", Short: "虚拟化引擎-short"}
	mRootCmd.AddCommand(GetClientCommnd())
	if err := mRootCmd.Execute(); err != nil {
		log.Println(err.Error())
	}
}
