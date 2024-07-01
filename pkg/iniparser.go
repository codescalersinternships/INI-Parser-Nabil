// packafe iniparser  implements a basic configuration language
// which provides a structure similar to what’s found in Microsoft Windows INI files.
// You can use this to write go programs which can be manpulate ini files by end users easily.
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

// IniParser implementing parser that loads and manipulates ini files as requested..
type IniParser struct {
	data map[string]section
}

// NewIniParser implementing new parser.
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

// LoadFromFile Read and parse a filename
// Files that cannot be opened return error, On success returns nil
// It try to parse my file
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

// LoadFromString Read and parse a string
// On success returns nil
// It try to parse my string
func (iniparser *IniParser) LoadFromString(fileString string) error {
	in := bufio.NewScanner(strings.NewReader(fileString))
	return iniparser.createIni(in)
}

// GetSectionNames Return a list of section names
func (iniparser *IniParser) GetSectionNames() []string {
	var sections []string
	for section := range iniparser.data {
		sections = append(sections, section)
	}
	return sections
}

// GetSections returns a map of sections and their key-value pairs
func (iniparser *IniParser) GetSections() map[string]section {
	return iniparser.data
}

// Get returns the value of the key in the sectionName, if the section or key is not found it returns an error
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

// Set set the value of the key in the sectionName, if the section or key is not found it returns an error
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

// String converts the parsed INI into a string
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

// SaveToFile saves the parsed INI to a file at the given path.
// On success returns nil.
func (iniparser *IniParser) SaveToFile(filePath string) error {
	if fileExt := filepath.Ext(filePath); fileExt != ".ini" {
		return fmt.Errorf("unsupported file format: %s", fileExt)
	}

	return os.WriteFile(filePath, []byte(iniparser.String()), 0644)
}
