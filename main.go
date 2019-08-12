package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"text/template"
)

type templateArgs struct {
	URL string
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("引数にはURLを指定してください")
	}

	templateArgs := &templateArgs{args[0]}

	createSourceCode(templateArgs)
}

// Create source code from template file
func createSourceCode(args *templateArgs) {
	FilePath := [][]string{
		{"template/template.go", "go/main.go"},
		{"template/template.java", "java/src/main/Main.java"},
		{"template/template.php", "php/main.php"},
		{"template/template.py", "python/main.py"},
		{"template/template.rb", "ruby/main.rb"},
		{"template/template.pl", "perl/main.pl"},
	}

	var wg sync.WaitGroup

	for _, path := range FilePath {
		wg.Add(1)

		go func(src string, dist string) {
			defer wg.Done()

			tpl, err := template.ParseFiles(src)

			if err != nil {
				log.Fatal(err)
			}

			buf := new(bytes.Buffer)

			if err := tpl.Execute(buf, args); err != nil {
				log.Fatal(err)
			}

			ioutil.WriteFile(dist, buf.Bytes(), 0655)

		}(path[0], path[1])

		wg.Wait()
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
