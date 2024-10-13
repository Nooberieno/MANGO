package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

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

func (e *ParseError) Error() string {
	errormsg := fmt.Sprintf("Error on line %d, %s:\n%s\n%s^",
		e.Line, e.Message, e.Context, strings.Repeat(" ", e.Column),
	)
	if e.Suggestion != "" {
		errormsg += fmt.Sprintf("\n%s", e.Suggestion)
	}
	return errormsg
}

func parse_file() error {
	file, err := os.Open(check_file())
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	in_target := false
	var current_target *Target

	for scanner.Scan() {
		i := 0
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		} else if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			varname := strings.TrimSpace(parts[0])
			varval := strings.TrimSpace(parts[1])
			variables[varname] = varval
		} else if strings.HasPrefix(line, "target") {
			targetname := strings.Split(line, " ")[1]
			if strings.HasPrefix(targetname, "{") {
				return &ParseError{
					Line:       i + 1,
					Column:     7,
					Message:    "target missing name",
					Context:    line,
					Suggestion: "Did you forget to name the target?",
				}
			}
			if !strings.Contains(line, "{") {
				return &ParseError{
					Line:       i + 1,
					Column:     7 + len(targetname),
					Message:    "target block not initialized, missing {",
					Context:    line,
					Suggestion: "Add '{' to open the target block.",
				}
			}
			in_target = true
			current_target = &Target{Name: strings.Trim(targetname, "{")}
			targets = append(targets, *current_target)
		} else if in_target && strings.HasPrefix(line, "-") {
			command := strings.TrimPrefix(line, "- ")
			true_command := variable_substitute(command)
			targets[len(targets)-1].Commands = append(targets[len(targets)-1].Commands, true_command)
		} else if strings.Contains(line, "}") && in_target {
			in_target = false
			current_target = nil
		} else if in_target && !scanner.Scan() && !strings.Contains(line, "}") {
			return &ParseError{
				Line:    i + 1,
				Column:  8 + len(current_target.Name),
				Message: "target block not ended, missing }",
				Context: "target " + current_target.Name + "{",
			}
		} else if in_target {
			return &ParseError{
				Line:       i + 1,
				Column:     0,
				Message:    "Unknown command inside of target block",
				Context:    "target " + current_target.Name,
				Suggestion: "Inside a target block use # for comments, - for command or use } to close the current target block",
			}
		} else if strings.TrimSpace(line) != "" {
			return &ParseError{
				Line:       i + 1,
				Column:     0,
				Message:    "Unknown command outside of target block",
				Context:    line,
				Suggestion: "Was this meant to be a comment?, if so use # at the beginning of the line",
			}
		}
		i++
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return nil
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
