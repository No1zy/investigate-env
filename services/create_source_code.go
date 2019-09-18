package services

import (
	"github.com/No1zy/investigate-env/config"
	"path/filepath"
	"text/template"
	"io/ioutil"
	"log"
	"bytes"
	"sync"
)

func (p *property) ReadTemplateDir(src string) ([][]string, error) {
	dirs, err := ioutil.ReadDir(src)

	if err != nil {
		log.Fatal(err)
	}

	path := [][]string{}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		for _, srv := range p.services {
			if d.Name() == srv.Name {
				tmp := []string{
					filepath.Join(src, srv.Name),
					srv.Dist,
				}
				path = append(path, tmp)
			}
		}
	}

	if len(path) < 1 {
		log.Fatal("テンプレートを取得できませんでした。")
	}

	return path, nil
}

// Create source code from template file
func CreateSourceCode(src, dist string, args *config.Variable) {
	files, err := ioutil.ReadDir(src)
	
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, f := range files {
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

		}(filepath.Join(src, f.Name()), filepath.Join(dist, f.Name()))

		wg.Wait()
	}
}
