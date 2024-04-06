package loop

import (
	"bufio"
	"fmt"
	cmdGT "git.jereileu.ch/gotables/client/gt-cli/gt/cmd"
	"git.jereileu.ch/gotables/client/gt-cli/shared"
	cmdSQL "git.jereileu.ch/gotables/client/gt-cli/sql/cmd"
	"os"
	"strings"
)

func GT(https bool, ip string, port string, token string) {
	var host string
	if https {
		host = "https://" + ip + ":" + port
	} else {
		host = "http://" + ip + ":" + port
	}
	shared.Location = "/"
	for {
		fmt.Print(start(shared.Location))
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			shared.HandleRequest("", err)
		}
		if strings.HasPrefix(scanner.Text(), "exit") {
			os.Exit(0)
		}
		request(host, scanner.Text(), false, token)
	}
}

func SQL() {}

func start(location string) string {
	out := "\n[ "
	out += location
	out += " ] --> "
	return out
}

func request(host string, input string, sql bool, token string) {
	if sql {
		cmdSQL.SetArgs(append(strings.Split(input, " "), "--host", host, "--token", token))
		cmdSQL.Execute()
	} else {
		cmdGT.SetArgs(append(strings.Split(input, " "), "--host", host, "--token", token))
		cmdGT.Execute()
	}
}
