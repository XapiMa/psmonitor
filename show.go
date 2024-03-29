package pschecker

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Shower is struct for show command
type Shower struct {
	types      int
	outputPath string
}

// NewShower create new shower object
func NewShower(typesString string, outputPath string) (*Shower, error) {
	s := new(Shower)
	var err error
	s.types, err = parseTypes(typesString)
	if err != nil {
		return s, errors.Wrap(err, "cause in NewShower")
	}
	s.outputPath = outputPath
	return s, nil
}

type OpenFile struct {
	Path string `json:"path"`
	Fd   int    `json:"fd"`
}

func parseOpen(str string) (string, error) {
	file := OpenFile{}
	if err := json.Unmarshal([]byte(str), &file); err != nil {
		return "", err
	}
	return file.Path, nil
}

// Show shows prosesses information
func (shower *Shower) Show() error {

	targets, err := getProcessesInfo(shower.types)
	if err != nil {
		return errors.Wrap(err, "cause in Show")
	}
	if len(targets) != 0 {
		if err := clearFile(shower.outputPath); err != nil {
			return errors.Wrap(err, "cause in Show")
		}
	}
	for _, target := range targets {
		text := "- "
		if target.Exec != "" {
			if text != "- " {
				text += "  "
			}
			text += fmt.Sprintf("exec: %s\n", target.Exec)
		}
		if target.Cmd != "" {
			if text != "- " {
				text += "  "
			}
			text += fmt.Sprintf("cmd: %s\n", target.Cmd)
		}
		if len(target.Open) != 0 {
			if text != "- " {
				text += "  "
			}
			text += "open: \n"
			for _, file := range target.Open {
				path, err := parseOpen(file)
				if err != nil {
					return errors.Wrap(err, "cause in parseOpen")
				}
				if path != "" {
					text += fmt.Sprintf("    - %s\n", path)
				}
			}
		}
		if target.User != "" {
			if text != "- " {
				text += "  "
			}
			text += fmt.Sprintf("user: %s\n", target.User)
		}
		if target.Pid != 0 {
			if text != "- " {
				text += "  "
			}
			text += fmt.Sprintf("pid: %d\n", target.Pid)
		}
		if text == "- " {
			continue
		}
		if err := appendFile(shower.outputPath, text); err != nil {
			return errors.Wrap(err, "cause in Show:")
		}
	}

	return nil
}
