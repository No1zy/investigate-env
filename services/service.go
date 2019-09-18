package services

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os/exec"
	"regexp"
)

type property struct {
	services []service
}

type service struct {
	Name      string
	Extension string
	Dist      string
}

type Daemon interface {
	Start([]string)
	Remove()
	Log([]string)
	IsValidName(string) bool
	ReadTemplateDir(string) ([][]string, error)
}

func New(opts ...Option) (Daemon, error) {
	prop := &property{}

	for _, opt := range append(defaultOptions, opts...) {
		if err := opt(prop); err != nil {
			return nil, errors.Wrap(err, "error create service")
		}
	}

	return prop, nil
}

func (p *property) IsValidName(name string) bool {
	for _, srv := range p.services {
		if name == srv.Name {
			return true
		}
	}
	return false
}

func (p *property) Start(args []string) {
	cmd := exec.Command("docker-compose", args...)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

// remove containers
func (p *property) Remove() {
	cmd := exec.Command("docker-compose", "rm", "-fv")

	cmd.Start()
	fmt.Println("Removing container...")
	cmd.Wait()

	removeDockerImages()
}

// remove docker images
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

// print logs on each container
func (p *property) Log(names []string) {
	if len(names) < 1 {
		for _, srv := range p.services {
			printLog(srv.Name)
		}
	} else {
		for _, name := range names {
			printLog(name)
		}
	}
}

func printLog(name string) {
	out, err := exec.Command("docker-compose", "logs", name).Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
