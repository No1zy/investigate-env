package main

import (
	"flag"
	"fmt"
	"github.com/No1zy/investigate-env/config"
	"github.com/No1zy/investigate-env/services"
	"log"
	"os"
	"sync"
	"strings"
)

func main() {

	var lang = flag.String("service", "",
		"Specify the service you want to test")
	var templateDir = flag.String("template", "template",
		"Specify the template directory to use")
	var createSourceCodeFlag = flag.Bool("only-create-source", false, "Only create source code, not build docker compose.")

	
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		args = []string{""}
	}

	if templateDir == nil {
		fmt.Println("テンプレートファイルが含まれるディレクトリを指定してください")
		return
	}

	loadConfig()

	srvs, err := services.New()

	commandArgs := []string{"up", "--build"}

	var langs []string
	if len(*lang) > 1 {
		langs = strings.Split(*lang, ",")
		for _, l := range langs {
			if srvs.IsValidName(l) {
				commandArgs = append(commandArgs, l)
			} else {
				fmt.Println("Invalid service")
				return
			}
		}
	}

	filePath, err := srvs.ReadTemplateDir(*templateDir, langs)

	if err != nil {
		fmt.Println("read template error")
		log.Fatal(err)
	}

	templateArgs := &config.Variable{args[0]}

	var wg sync.WaitGroup

	for _, path := range filePath {
		wg.Add(1)

		go func(src, dist string) {
			defer wg.Done()

			services.CreateSourceCode(src, dist, templateArgs)

		}(path[0], path[1])

		wg.Wait()
	}

	if !*createSourceCodeFlag {
		srvs.Start(commandArgs)

		if len(commandArgs) > 2 {
			srvs.Log(commandArgs[2:])
		} else {
			srvs.Log(nil)
		}
		srvs.Remove()
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

// load config
func loadConfig() {
	err := config.Load("docker-compose.yml")
	if err != nil {
		log.Fatal(err)
	}
}
