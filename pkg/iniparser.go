package iniparser

import (
	"bufio"
	"fmt"
	"log"
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

func (iniparser *IniParser) createIni(scanner *bufio.Scanner) {
	var currSection string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}
		if line[0] == '[' {
			currSection = line[1 : len(line)-1]
			iniparser.data[currSection] = make(map[string]string)
			continue
		}
		temp := strings.SplitN(line, "=", 2)
		key, val := temp[0], temp[1]
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if len(key) == 0 || len(val) == 0 {
			log.Fatal("key val aren't valid")
		}
		curval, isKeyFound := iniparser.data[currSection][key]
		if isKeyFound {
			log.Fatal("key is dublicated already has value = ", curval)
		}
		iniparser.data[currSection][key] = val
	}
}

func (iniparser *IniParser) LoadFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Incorrect file path or file does not exist ", err)
	}
	defer file.Close()
	if path.Ext(filePath) != ".ini" {
		log.Fatal("unsupported file format ")
	}
	in := bufio.NewScanner(file)
	iniparser.createIni(in)
}

func (iniparser *IniParser) LoadFromString(fileString string) {
	in := bufio.NewScanner(strings.NewReader(fileString))
	iniparser.createIni(in)
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

func (iniparser *IniParser) Get(sectionName string, key string) string {
	if len(sectionName) == 0 {
		log.Fatal("section name is invalid")
	}
	val, found := iniparser.data[sectionName][key]
	if !found {
		log.Fatal("section name and key aren't found")
	}
	return val
}
func (iniparser *IniParser) Set(sectionName string, key string, val string) {
	if len(sectionName) == 0 {
		log.Fatal("section name is invalid")
	}
	curval, found := iniparser.data[sectionName][key]
	if !found {
		log.Fatal("section name and key aren't found", curval)
	}
	iniparser.data[sectionName][key] = val
}

func (iniparser *IniParser) ToString() string {
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

	return os.WriteFile(filePath, []byte(iniparser.ToString()), 0644)
}
