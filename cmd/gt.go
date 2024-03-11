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
	"fmt"
	"git.jereileu.ch/gotables/client/gt-cli/loop"
	"git.jereileu.ch/gotables/client/gt-cli/shared"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"log"
	"os"
)

var (
	https    bool
	ip       string
	port     string
	username string
	password string
)

// gtCmd represents the gt command
var gtCmd = &cobra.Command{
	Use:   "gt",
	Short: "connect to gt-server",
	Long:  `Connect to a GoTables server using the GT query language`,
	Run: func(cmd *cobra.Command, args []string) {
		if password == "" {
			fmt.Print("Password: ")
			pw, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatalln("Failed to receive password from command line")
			}
			password = string(pw)
		}
		token, err := shared.Login(username, password)
		if err != nil {
			log.Fatalln("Wrong username or password")
		}
		fmt.Println("")
		fmt.Println("Authentication successful")
		loop.GT(https, ip, port, token)
	},
}

func init() {
	connectCmd.AddCommand(gtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	gtCmd.Flags().BoolVar(&https, "https", false, "")
	gtCmd.Flags().StringVar(&ip, "ip", "127.0.0.1", "ip of the GoTables server")
	gtCmd.Flags().StringVar(&port, "port", "5678", "port of the GoTables server")
	gtCmd.Flags().StringVarP(&username, "username", "u", "", "username to be used to connect to GoTables")
	gtCmd.Flags().StringVarP(&password, "password", "p", "", "password to be used to connect to GoTables")

	err := gtCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
}
