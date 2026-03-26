package internal

import "strings"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
	ParamType   string
}

var Registry map[string]CliCommand

func CleanInput(text string) []string {
	var returnSlice []string
	returnSlice = strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return returnSlice
}
