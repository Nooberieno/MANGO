package main

import (
	"regexp"
	"strings"
)

type Target struct {
	Name     string
	Commands []string
}

var variables = map[string]string{}
var targets = []Target{}

func parse(contents string) {
	lines := strings.Split(contents, "\n")
	in_target := false
	var current_target *Target
	for _, line := range lines {
		line = strings.TrimSpace(line)
		in_command := false
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			varname := strings.TrimSpace(parts[0])
			varval := strings.TrimSpace(parts[1])
			variables[varname] = varval
		} else if strings.HasPrefix(line, "target") {
			targetname := strings.Split(line, " ")[1]
			if !strings.Contains(line, "{") {
				panic("Target block not missing { at initilization")
			}
			in_target = true
			current_target = &Target{Name: targetname}
			targets = append(targets, *current_target)
		} else if in_target && strings.HasPrefix(line, "-") {
			in_command = true
			command := strings.TrimPrefix(line, "- ")
			true_command := variable_substitute(command)
			targets[len(targets)-1].Commands = append(targets[len(targets)-1].Commands, true_command)
		} else if !in_command && in_target {
			strings.Contains(line, "}")
			in_target = false
			current_target = nil
		}
	}
}

func variable_substitute(command string) string {
	re := regexp.MustCompile(`\$\{([a-zA-Z_][a-zA-Z0-9_]*)\}`)
	return re.ReplaceAllStringFunc(command, func(varcall string) string {
		varname := varcall[2 : len(varcall)-1]
		if value, exists := variables[varname]; exists {
			return value
		}
		return varcall
	})
}
