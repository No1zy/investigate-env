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

func createSourceCode(args *templateArgs) {
	FilePath := [][]string{
		{"template/template.go", "go/main.go"},
		{"template/template.java", "java/src/main/java/Main.java"},
		{"template/template.php", "php/main.php"},
	}

	var (
		templateText string
		wg           sync.WaitGroup
	)

	for _, path := range FilePath {
		wg.Add(1)

		go func(src string, dist string) {
			defer wg.Done()

			if fileExists(src) {
				f, err := os.Open(src)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				buf, err := ioutil.ReadAll(f)
				if err != nil {
					log.Fatal(err)
				}
				templateText = string(buf)
			}

			tpl := template.Must(template.New("template").Parse(templateText))

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
