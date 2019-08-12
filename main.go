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

const DIST_PREFIX = "dist"

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
		{"template/template.go", DIST_PREFIX + "/go/main.go"},
		{"template/template.java", DIST_PREFIX + "/java/src/main/Main.java"},
		{"template/template.php", DIST_PREFIX + "/php/main.php"},
		{"template/template.py", DIST_PREFIX + "/python/main.py"},
		{"template/template.rb", DIST_PREFIX + "/ruby/main.rb"},
		{"template/template.pl", DIST_PREFIX + "/perl/main.pl"},
		{"template/template.js", DIST_PREFIX + "/node/main.js"},
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
