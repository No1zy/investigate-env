package services

import (
	"fmt"
	"bytes"
	"github.com/No1zy/investigate-env/config"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
	"os"
)

func (p *property) ReadTemplateDir(src string) ([][]string, error) {


	dirs, err := ioutil.ReadDir(src)

	if err != nil {
		fmt.Println("read error: " + src)
		log.Fatal(err)
	}

	var path [][]string

	for _, d := range dirs {
		for _, srv := range p.services {
			path = append(path, recursiveMkDir(filepath.Join(src, d.Name()), &srv)...)
		}
	}

	if len(path) < 1 {
		log.Fatal("テンプレートを取得できませんでした。")
	}

	return path, nil
}

// Create source code from template file
func CreateSourceCode(src, dist string, args *config.Variable) {
	tpl, err := template.ParseFiles(src)

	if err != nil {
		fmt.Println("parse error")
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	if err := tpl.Execute(buf, args); err != nil {
		fmt.Println("exec error")
		log.Fatal(err)
	}

	ioutil.WriteFile(dist, buf.Bytes(), 0655)


}

func recursiveMkDir(src string, service *service) [][]string {
	files, err := ioutil.ReadDir(src)
	
	if err != nil {
		log.Fatal(err)
	}

	path := [][]string{}
	
	for _, f := range files {
		if f.IsDir() {
			if _, err := os.Stat(f.Name()); err != nil {
				os.Mkdir(filepath.Join(service.Dist, src, f.Name()), 0755)
			}
			path = append(path, recursiveMkDir(filepath.Join(src, f.Name()), service)...)
			continue
		}
		if f.Name() == service.Name {
			tmp := []string{
				filepath.Join(src, service.Name),
				service.Dist,
			}
			path = append(path, tmp)
		}
	}
	return path
}
