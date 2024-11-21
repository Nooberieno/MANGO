package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\$\{([a-zA-Z_][a-zA-Z0-9_]*)\}`)

type ParseError struct {
	Line       int
	Column     int
	Message    string
	Context    string
	Suggestion string
}

type Target struct {
	Name     string
	Commands []string
}

var variables = map[string]string{}
var targets = []Target{}

// If an error occurs while parsing, display it in a helpfull
func (e *ParseError) Error() string {
	errormsg := fmt.Sprintf("Error on line %d, %s:\n%s\n%s^",
		e.Line, e.Message, e.Context, strings.Repeat(" ", e.Column),
	)
	if e.Suggestion != "" {
		errormsg += fmt.Sprintf("\n%s", e.Suggestion)
	}
	return errormsg
}

// Parses the build file according to MANGO's parsing rules outlined in Rules.md
// A target must be initialized by the "target" keyword
// A target block starts with an "{" and ends with a "}"
// Commands can only be put inside a target block and be seperated by newlines
// A command must be preceded by a "-"
// Only line comments are allowed and are made by starting the line with a "#"
// Variables are initialzed as followed: <name> = <value>
// Variables are referenced as followed: ${variable_name}
func parse_file() error {
	file, err := os.Open(check_file())
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var current_target *Target
	for line_number := 1; scanner.Scan(); line_number++ {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 || trimmed[0] == '#' {
			continue
		}

		switch {
		case strings.Contains(trimmed, "="):
			parts := strings.SplitN(trimmed, "=", 2)
			variables[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		case strings.HasPrefix(trimmed, "target"):
			if err := handle_target(&current_target, trimmed, line_number); err != nil {
				return err
			}
		case current_target != nil && strings.HasPrefix(trimmed, "-"):
			command := strings.TrimPrefix(trimmed, "- ")
			true_command := variable_substitute(command)
			current_target.Commands = append(current_target.Commands, true_command)
		case strings.Contains(trimmed, "}") && current_target != nil:
			current_target = nil
		default:
			return &ParseError{
				Line:    line_number,
				Column:  0,
				Message: "Unknown command",
				Context: line,
			}
		}
	}

	return scanner.Err()
}

// Handles target information and errors, used by parse_file
func handle_target(current_target **Target, line string, line_number int) error {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return &ParseError{
			Line:       line_number,
			Column:     7,
			Message:    "target missing name",
			Context:    line,
			Suggestion: "Did you forget to name the target?",
		}
	}
	targetName := parts[1]
	if !strings.HasSuffix(targetName, "{") {
		return &ParseError{
			Line:       line_number,
			Column:     7 + len(targetName),
			Message:    "target block not initialized, missing {",
			Context:    line,
			Suggestion: "Add '{' to open the target block.",
		}
	}
	*current_target = &Target{Name: strings.TrimSuffix(targetName, "{")}
	targets = append(targets, **current_target)
	return nil
}

// Function to handle variable substitution, i.e. replacing a variable reference "${variable} with its value"
// If a variable reference "${variable_name}" is encounterd, check if it is defined
// If it is defined, return the variable
// If it isnt defined, return the reference as is
func variable_substitute(command string) string {
	return re.ReplaceAllStringFunc(command, func(varcall string) string {
		varname := varcall[2 : len(varcall)-1]
		if value, exists := variables[varname]; exists {
			return value
		}
		return varcall
	})
}
