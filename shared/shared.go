package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	Location string
)

type Body struct {
	Query     string `json:"query"`
	SessionId string `json:"session_id"`
}

type Column struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default any    `json:"default"`
}

type Table struct {
	Columns []Column         `json:"columns"`
	Rows    []map[string]any `json:"rows"`
}

func (t Table) Ascii() string {
	out := ""
	x := len(t.Columns) + 1
	y := len(t.Rows) + 1
	maxLengths := make([]int, x)
	maxLengths[0] = digits(y - 1)
	for i := 0; i < len(t.Columns); i++ {
		maxLengths[i+1] = len(t.Columns[i].Name)
		for j := 0; j < len(t.Rows); j++ {
			length := len(t.Rows[j][t.Columns[i].Name].(string))
			if length > maxLengths[i+1] {
				maxLengths[i+1] = length
			}
		}
	}
	out += drawSeparator(0, maxLengths)
	content := make([]string, x)
	for i := 0; i < x; i++ {
		if i == 0 {
			content[0] = ""
		} else {
			content[i] = t.Columns[i-1].Name
		}
	}
	out += drawContent(content, maxLengths)
	out += drawSeparator(1, maxLengths)
	for i := 1; i < y; i++ {
		content[0] = strconv.Itoa(i)
		for j := 1; j < x; j++ {
			content[j] = t.Rows[i-1][t.Columns[j-1].Name].(string)
		}
		out += drawContent(content, maxLengths)
		out += drawSeparator(0, maxLengths)
	}
	return out
}

func ValidUrl(path string) bool {
	parts := strings.Split(path, "/")[1:]
	if len(parts) == 1 || len(parts) == 2 {
		return true
	}
	return false
}

func TableSet(path string) (string, bool) {
	if len(strings.Split(path, "/")[1:]) == 2 {
		return path, true
	}
	return "", false
}

func DbSet(path string) (string, bool) {
	parts := strings.Split(path, "/")[1:]
	if len(parts) > 0 && parts[0] != "" {
		return "/" + parts[0], true
	}
	return "", false
}

func Request(host, path, query, sessionId string) (string, error) {
	path = strings.TrimSuffix(path, "/")
	url := host + path
	body := Body{
		Query:     query,
		SessionId: sessionId,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return "", errors.New("Failed to create request: " + err.Error())
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return "", errors.New("Failed to create request: " + err.Error())
	}

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	// Read the response body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Failed to read response body: " + err.Error())
	}
	var tbl Table
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return "", errors.New("Failed to read response body: " + err.Error())
	}
	return tbl.Ascii(), nil
}

func HandleRequest(data string, err error) {
	if err != nil {
		fmt.Println("Error: Failed to send http request: " + err.Error())
	} else {
		fmt.Println(data)
	}
}

func Login(u, p string) (string, error) {
	return "login not implemented yet on the server side", nil
}

func MakeRequest(query, host, token string) {
	if ValidUrl(Location) {
		ret, err := Request(host, Location, query, token)
		HandleRequest(ret, err)
	} else {
		fmt.Println("Path is invalid. Use 'set db', 'set table' and 'unset' to modify the path")
	}
}

func ConnectArgs(args []string) string {
	out := ""
	for i := 0; i < len(args); i++ {
		out += args[i]
		out += " "
	}
	out = strings.TrimSpace(out)
	return out
}

func digits(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func drawSeparator(style int, lengths []int) string {
	out := ""
	if style == 0 {
		out = "+"
		for i := 0; i < len(lengths); i++ {
			for j := 0; j < lengths[i]+4; j++ {
				out += "-"
			}
			out += "+"
		}
	} else if style == 1 {
		out = "+"
		for i := 0; i < len(lengths); i++ {
			for j := 0; j < lengths[i]+4; j++ {
				out += "="
			}
			out += "+"
		}
	}
	return out + "\n"
}

func drawContent(content []string, lengths []int) string {
	out := "|"
	for i := 0; i < len(lengths); i++ {
		length := lengths[i] - len(content[i]) + 4
		for j := 0; j < length/2; j++ {
			out += " "
		}
		out += content[i]
		if length%2 != 0 {
			length += 1
		}
		for j := 0; j < length/2; j++ {
			out += " "
		}
		out += "|"
	}
	return out + "\n"
}
