/*
Copyright © 2024 Jeroen Leuenberger <jereileu@proton.me>

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
package cmd

import (
	"git.jereileu.ch/gotables/client/gt-cli/shared"
	"strings"

	"github.com/spf13/cobra"
)

// unsetTableCmd represents the unsetTable command
var unsetTableCmd = &cobra.Command{
	Use:   "table",
	Short: "Unset table",
	Long: `Unset the name of the table to interact with
Requires NOTHING to be set OR
Requires db to be set OR
Requires db AND table to be set
Syntax: unset table`,
	Run: func(cmd *cobra.Command, args []string) {
		shared.Location = "/" + strings.Split(shared.Location, "/")[1]
	},
}

func init() {
	unsetCmd.AddCommand(unsetTableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unsetTableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unsetTableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
