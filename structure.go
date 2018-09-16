package main

import (
	"fmt"
	"errors"
	"io/ioutil"
	"encoding/json"
)

type Structure struct {
	Columns []Column `json:"columns"`
}

func (s *Structure) Load(structureFile string) error {
    b, err := ioutil.ReadFile(structureFile)
    if err != nil {
        return errors.New("Cannot read file!")
    }
	err = json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("Cannot json unmarshal! %s", err)
	}
	return nil	
}

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Objects []Object `json:"objects"`
}

type Object struct {
	Name string `json:"name"`
	Patterns [][]string `json:"patterns"`
	Sub []Object `json:"sub"`
}
