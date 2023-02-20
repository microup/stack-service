package stack

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"stack-service/internal/utils"
)

const dataFile = "stack.json"

type Stack struct {
	data []string
}

func New() *Stack {
	return &Stack{
		data: make([]string, 0),
	}
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
	s.Save()
}

func (s *Stack) Pop() (string, bool) {
	if len(s.data) == 0 {
		return "", false
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.Save()
	return value, true
}

func (s *Stack) Top() (string, bool) {
	if len(s.data) == 0 {
		return "", false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack) Save() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(s.data)
}

func (s *Stack) LoadStack() error {
	if !utils.FileExists(dataFile) {
		return nil
	}

	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &s.data); err != nil {
		return err
	}

	return nil
}
