package stack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"  //nolint:staticcheck
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
	_ = s.Save()
}

func (s *Stack) Pop() (string, bool) {
	if len(s.data) == 0 {
		return "", false
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	_ = s.Save()

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
		return fmt.Errorf("%w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(s.data)

	if err != nil {
		return fmt.Errorf("%w", err)
	} 

	return nil
}

func (s *Stack) LoadStack() error {
	if !utils.FileExists(dataFile) {
		return nil
	}

	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(file, &s.data); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
