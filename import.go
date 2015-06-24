package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

func Import(name string) (*Code, error) {
	if name == "sys" {
		return NewSys(), nil
	}

	cwd, err := ioutil.ReadDir(".")

	if err != nil {
		return nil, err
	}

	for _, file := range cwd {
		if strings.HasPrefix(file.Name(), name) {
			if strings.HasSuffix(file.Name(), ".pyc") {
				// TODO(flowlo): Compile if necessary
				return nil, errors.New("Not implemented.")
			} else if strings.HasSuffix(file.Name(), ".py") {
				// TODO(flowlo): Check pycache
				return nil, errors.New("Not implemented.")
			}
		}
	}

	// search cwd
	// search $PYTHONPATH
	// search lel

	return nil, errors.New("Module importing is not implemented.")
}
