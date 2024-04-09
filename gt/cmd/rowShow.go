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

	"github.com/spf13/cobra"
)

// rowShowCmd represents the rowShow command
var rowShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display rows",
	Long: `This command displays rows in a table
Requires db AND table to be set
Syntax: row show [index1:index2:index3 ... ]`,
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()
		token := cmd.Flag("token").Value.String()
		query := "row show " + shared.ConnectArgs(args)
		shared.MakeRequest(query, host, token)
	},
}

func init() {
	rowCmd.AddCommand(rowShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rowShowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rowShowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
