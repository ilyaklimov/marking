package main

import (
	"fmt"
	"log"
	"flag"
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
	"encoding/csv"
	"os"
)

var err error

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()

	var args []string = flag.Args()

	if len(args) < 3 {
		log.Fatalf("There are not enough arguments!")
	}

	keywords, err := loadKeywords(args[0])
	if err != nil {
		log.Fatalf("Cannot load keywords! %s", err)
	}

	var structure Structure
	if err = structure.Load(args[1]); err != nil {
		log.Fatalf("Cannot load structure! %s", err)
	}

	marks, err := mark(keywords, structure)
	if err != nil {
		log.Fatalf("Cannot mark keywords! %s", err)
	}

	err = save(args[2], marks)
	if err != nil {
		log.Fatalf("Cannot save CSV file")
	}

	fmt.Println("OK! BayBay...")

}


func loadKeywords(keywordsFile string) ([]string, error) {
	var keywords []string
    b, err := ioutil.ReadFile(keywordsFile)
    if err != nil {
        return keywords, errors.New("Cannot read file!")
    }
    keywords = strings.Split(string(b), "\n")
    return keywords, nil	
}


func mark(keywords []string, structure Structure) ([][]string, error) {
	var marks [][]string
	var mark []string
	var name string
	var header []string = getHeader(structure.Columns)
	marks = append(marks, header)
	for _, k := range keywords {
		mark = append(mark, k)
		for _, column := range structure.Columns {
			switch column.Type {
			case "tree":
				name = findTreeObjects(k, column.Objects)
			case "tags":
				name = findTagsObjects(k, column.Objects)
			default: 
				return marks, errors.New("Type column is empty!")
			}
			mark = append(mark, name)
		}
		marks = append(marks, mark)
		mark = mark[len(mark):]
	}
	return marks, nil
}


func getHeader(columns []Column) []string {
	var header []string
	header = append(header, "Phrase")
	for _, column := range columns {
		header = append(header, column.Name)
	}
	return header	
}


func findTreeObjects(keywords string, objects []Object) string {
	var name string = ""
	for _, object := range objects {
		if findPatterns(keywords, object.Patterns) {
			if len(object.Sub) != 0 {
				name = findTreeObjects(keywords, object.Sub)
				if name != "" {
					return name
				}
			}
			return object.Name
		}
	}
	return name
}



func findTagsObjects(keywords string, objects []Object) string {
	tags := make(map[string][]string)
	for _, object := range objects {
		for _, subobject := range object.Sub {
			if findPatterns(keywords, subobject.Patterns) {
				tags[object.Name] = append(tags[object.Name], subobject.Name)
			}			
		}
	}
	b, err := json.Marshal(tags)
	if err != nil {
		log.Fatalf("Cannot json marshal! %s", err)
	}	
	return string(b)
}


func findPatterns(keywords string, patterns [][]string) bool {
	var matched int
	for _, pattern := range patterns {
		matched = 0
		for _, p := range pattern {
			if strings.Contains(keywords, p) {
				matched++
			}
		}
		if len(pattern) == matched {
			return true
		}
	}
	return false
}


func save(outfile string, marks [][]string) error {
	csvfile, err := os.Create(outfile)
	if err != nil {
		log.Fatalf("Cannot create %s file! %s", outfile, err)
	}
	defer csvfile.Close()
	w := csv.NewWriter(csvfile)
	w.Comma = ';'
	w.WriteAll(marks)
	if err := w.Error(); err != nil {
		log.Fatalf("Cannot write in %s file! %s", outfile, err)
	}	
	return nil
}


