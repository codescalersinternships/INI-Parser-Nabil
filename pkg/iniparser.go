package iniparser

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type section map[string]string

type IniParser struct {
	data map[string]section
}

func NewIniParser() *IniParser {
	return &IniParser{
		data: make(map[string]section),
	}
}

func (iniparser *IniParser) createIni(scanner *bufio.Scanner) error {
	var currSection string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}
		if line[0] == '[' {
			possibleCurSec := line[1 : len(line)-1]
			possibleCurSec = strings.TrimSpace(possibleCurSec)
			if len(possibleCurSec) == 0 {
				return fmt.Errorf("sections name can't be empty")
			}
			currSection = possibleCurSec
			iniparser.data[currSection] = make(map[string]string)
			continue
		}
		keyVal := strings.SplitN(line, "=", 2)
		key, val := keyVal[0], keyVal[1]
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if len(key) == 0 || len(val) == 0 {
			return fmt.Errorf("key val aren't valid")
		}
		curval, isKeyFound := iniparser.data[currSection][key]
		if isKeyFound {
			return fmt.Errorf("key is dublicated already has value = %s", curval)
		}
		iniparser.data[currSection][key] = val
	}
	return nil
}

func (iniparser *IniParser) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("incorrect file path or file does not exist %e", err)
	}
	defer file.Close()
	if path.Ext(filePath) != ".ini" {
		return fmt.Errorf("unsupported file format ")
	}
	in := bufio.NewScanner(file)
	return iniparser.createIni(in)
}

func (iniparser *IniParser) LoadFromString(fileString string) error {
	in := bufio.NewScanner(strings.NewReader(fileString))
	return iniparser.createIni(in)
}
func (iniparser *IniParser) GetSectionNames() []string {
	var sections []string
	for section := range iniparser.data {
		sections = append(sections, section)
	}
	return sections
}
func (iniparser *IniParser) GetSections() map[string]section {
	return iniparser.data
}

func (iniparser *IniParser) Get(sectionName string, key string) (string, error) {
	if len(sectionName) == 0 {
		return "", fmt.Errorf("section name is invalid")
	}
	val, found := iniparser.data[sectionName][key]
	if !found {
		return "", fmt.Errorf("section name and key aren't found")
	}
	return val, nil
}
func (iniparser *IniParser) Set(sectionName string, key string, val string) error {
	if len(sectionName) == 0 {
		return fmt.Errorf("section name is invalid")
	}
	curval, found := iniparser.data[sectionName][key]
	if !found {
		return fmt.Errorf("section name and key aren't found %s", curval)
	}
	iniparser.data[sectionName][key] = val
	return nil
}

func (iniparser *IniParser) String() string {
	var str string
	for sectionName, section := range iniparser.data {
		str += fmt.Sprintf("[%v]\n", sectionName)
		for k, v := range section {
			str += fmt.Sprintf("%v=%v\n", k, v)
		}
	}

	return str
}

func (iniparser *IniParser) SaveToFile(filePath string) error {
	if fileExt := filepath.Ext(filePath); fileExt != ".ini" {
		return fmt.Errorf("unsupported file format: %s", fileExt)
	}

	return os.WriteFile(filePath, []byte(iniparser.String()), 0644)
}
