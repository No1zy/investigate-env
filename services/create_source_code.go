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

func (p *property) ReadTemplateDir(src string, langs []string) ([][]string, error) {

	var targets []string
	if len(langs) < 1 {
		for _, srv := range p.services {
			targets = append(targets, srv.Name)
		}
	} else {
		targets = langs
	}

	dirs, err := ioutil.ReadDir(src)

	if err != nil {
		fmt.Println("read error: " + src)
		log.Fatal(err)
	}

	var path [][]string

	for _, d := range dirs {
		for _, target := range targets {
			if d.Name() == target {
				path = append(path, recursiveMkDir(src, d.Name(), p.GetService(target))...)
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

func recursiveMkDir(src, pathParts string, service *service) [][]string {
	fmt.Printf("pathParts: %v\n", pathParts)
	target := filepath.Join(src, pathParts)
	files, err := ioutil.ReadDir(target)
	
	if err != nil {
		log.Fatal(err)
	}

	path := [][]string{}
	
	for _, f := range files {
		if f.IsDir() {
			if err := os.MkdirAll(filepath.Join(service.Dist, pathParts, f.Name()), 0755); err != nil {
			        fmt.Println(err)
			    }
			path = append(path, recursiveMkDir(src, filepath.Join(pathParts, f.Name()), service)...)
			continue
		}
		fmt.Printf("path: %v\n", filepath.Join(service.Dist, pathParts, f.Name()))
		tmp := []string{
			filepath.Join(src, pathParts, f.Name()),
			filepath.Join(service.Dist, pathParts, f.Name()),
		}
		path = append(path, tmp)
	}
	return path
}
