package util

import (
	"bufio"
	"io"
)

// State used in reading the front matter.
const (
	stateStart         = 0
	stateInFrontMatter = 1
	stateInBody        = 2
)

// Front matter reader.
type FrontMatter struct {
	tag string
}

func NewFrontMatter(tag string) *FrontMatter {
	return &FrontMatter{tag}
}

func (fm *FrontMatter) Parse(input io.Reader) (front, body string, err error) {
	s := bufio.NewScanner(input)
	front, body = "", ""
	state := stateStart
	lines := 0
	for s.Scan() {
		t := s.Text()
		if lines == 0 && t != fm.tag {
			state = stateInBody
		}
		if state == stateInFrontMatter && t != fm.tag {
			front += t + "\n"
		} else if state == stateInBody {
			body += t + "\n"
		}
		if t == fm.tag {
			if state < stateInBody {
				state++
			}
		}
		lines++
	}

	err = s.Err()
	if err != nil {
		return "", "", err
	}

	return front, body, nil
}
