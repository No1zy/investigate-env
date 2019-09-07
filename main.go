package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sync"
	"text/template"
)

type templateArgs struct {
	URL string
}

const DIST_PREFIX = "dist"

func main() {

	distPath := map[string]string{
		"go":   DIST_PREFIX + "/go/main.go",
		"java": DIST_PREFIX + "/java/src/main/Main.java",
		"php":  DIST_PREFIX + "/php/main.php",
		"py":   DIST_PREFIX + "/python/main.py",
		"rb":   DIST_PREFIX + "/ruby/main.rb",
		"pl":   DIST_PREFIX + "/perl/main.pl",
		"js":   DIST_PREFIX + "/node/main.js",
	}

	var lang = flag.String("lang", "",
		"Specify the language you want to test")
	var templateDir = flag.String("template", "template",
		"Specify the template directory to use")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("引数にはURLを指定してください")
		return
	}

	if templateDir == nil {
		fmt.Println("テンプレートファイルが含まれるディレクトリを指定してください")
		return
	}

	files, err := ioutil.ReadDir(*templateDir)

	if err != nil {
		log.Fatal(err)
	}

	langs := []string{
		"go",
		"java",
		"php",
		"python",
		"ruby",
		"perl",
		"node",
	}

	var exts []string

	commandArgs := []string{"up"}

	if len(*lang) > 1 {
		for _, l := range langs {
			if *lang == l {
				exts = append(exts, getExtFromLang(l))
				commandArgs = append(commandArgs, "--build")
				commandArgs = append(commandArgs, l)
			}
		}
	} else {
		exts = []string{
			"go",
			"java",
			"php",
			"py",
			"rb",
			"pl",
			"js",
		}
	}

	if len(exts) < 1 {
		fmt.Println("--lang 引数には対応している言語を指定してください: go, java, php, python, ruby, perl, javascript")
		return
	}

	path := [][]string{}

	for _, f := range files {
		for _, e := range exts {
			if filepath.Ext(f.Name())[1:] == e {
				tmp := []string{
					filepath.Join(*templateDir, f.Name()),
					distPath[e],
				}
				path = append(path, tmp)
			}
		}
	}

	if len(path) < 1 {
		fmt.Println("指定された言語のテンプレートを取得できませんでした。")
		return
	}

	templateArgs := &templateArgs{args[0]}

	createSourceCode(templateArgs, path)

	_, err = exec.Command("docker-compose", commandArgs...).Output()
	if err != nil {
		log.Fatal("Failed build :" + err.Error())
	}

	printLogDockerCompose(exts)
}

// Create source code from template file
func createSourceCode(args *templateArgs, filePath [][]string) {

	var wg sync.WaitGroup

	for _, path := range filePath {
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

func getExtFromLang(lang string) string {
	exts := map[string]string{
		"go":     "go",
		"java":   "java",
		"php":    "php",
		"python": "py",
		"ruby":   "rb",
		"perl":   "pl",
		"node":   "js",
	}

	return exts[lang]
}

func getLangFromExt(ext string) string {
	exts := map[string]string{
		"go":   "go",
		"java": "java",
		"php":  "php",
		"py":   "python",
		"rb":   "ruby",
		"pl":   "perl",
		"js":   "node",
	}

	return exts[ext]

}

func printLogDockerCompose(exts []string) {

	for _, e := range exts {
		out, err := exec.Command("docker-compose", "logs", getLangFromExt(e)).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", out)
	}

	removeDockerCompose()
	removeDockerImages()

}

func removeDockerCompose() {
	cmd := exec.Command("docker-compose", "rm", "-fv")

	cmd.Start()
	fmt.Println("Removing container...")
	cmd.Wait()
}

func removeDockerImages() {

	images, err := exec.Command("sudo", "docker", "images").Output()

	regNone := regexp.MustCompile(`<none>.*([0-9abcdef]{12})`)
	regId := regexp.MustCompile(`([0-9abcdef]{12})`)
	res := regNone.FindAll(images, -1)

	var ids [][]byte

	for _, b := range res {
		findId := regId.Find(b)
		ids = append(ids, findId)
	}

	if len(ids) < 1 {
		return
	}

	var out []byte
	for _, id := range ids {
		out, err = exec.Command("sudo", "docker", "rmi", string(id)).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", out)
	}

}
