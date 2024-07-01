# Unix-Based Commands Implemented in Go
This module provides the ConfigParser class which implements a basic configuration language which provides a structure similar to what’s found in Microsoft Windows INI files. You can use this to write go programs which can be manipulated by end users easily.

# Features

- LoadFromString - converts a given INI as string to a map stored in the parser
- LoadFromFile - converts a given INI file to a map stored in the parser
- GetSectionNames - list of all section names
- GetSections - serialize convert into a dictionary/map  { section_name: {key1: val1, key2, val2} ...}
- Get(section_name, key) - gets the value of key in section section_name
- Set(section_name, key, value)  - sets a key in section section_name to value value
- String - converts the currently stored map to string
- SaveToFile - converts the currently stored map to an INI file


# How to Use

1- import package

```golang
import github.com/codescalersinternships/INI-Parser-Nabil
```

2- create a new parser struct using NewParser()

```golang
parser := NewIniParser()
```

3- load from a file

```golang
_ = parser.LoadFromFile("./testdata/validini.ini")
```

4- load from a string

```golang
_ = parser.LoadFromString(validStringInput)
```

5- get a key value from a section

```golang
val, _ := parser.Get("Simple Values", "key")
fmt.Println(val)
// Output: value
```

6- set a value for a key in a section

```golang
_ = parser.Set("Simple Values", "key", "newvalue")
val, _ := parser.Get("Simple Values", "key")
fmt.Println(val)
// Output: newvalue
```

7- get section names

```golang
sectionsNames, _ := parser.GetSectionNames()
```

8- get parsed data

```golang
sections, _ := parser.GetSections()
```

9- convert data to string

```golang
str, _ := parser.String()
```

10- save data to file

```golang
_ = parser.SaveToFile("./testdata/output.ini")
```

## How to Test

- run go test ./... in root directory

```golang
go test ./...
```

- add the -v flag for more details about the specific tests that are running

```golang
go test -v ./...
```