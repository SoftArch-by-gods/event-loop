package commands

import (
	"fmt"
	"strings"
)

func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)
	switch len(parts) {
	case 0:
		return printCommand("SYNTAX ERROR: empty string")
	case 1:
		return printCommand(fmt.Sprintf("SYNTAX ERROR: no arg or undefined command: %s", parts[0]))
	case 2:
		if parts[0] == "print" {
			return printCommand(parts[1])
		}
		if parts[0] == "palindrom" {
			return palindromCommand(parts[1])
		}
		return printCommand(fmt.Sprintf("SYNTAX ERROR: undefined command: %s", parts[0]))
	default:
		return printCommand(fmt.Sprintf("SYNTAX ERROR: too many arguments"))
	}
}
