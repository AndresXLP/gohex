package app

import (
	"log"
	"os/exec"
)

func (fl *Service) InitGoModule(module string) {
	launchGoModInitAndGoModTidy(module)
}

func launchGoModInitAndGoModTidy(module string) {
	_, err := exec.Command("go", "mod", "init", module).Output()
	if err != nil {
		log.Fatal("error mod ", err)
	}
	_, err = exec.Command("go", "mod", "tidy").Output()
	if err != nil {
		log.Fatal("error tidy ", err)
	}
}
