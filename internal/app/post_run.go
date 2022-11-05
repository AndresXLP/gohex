package app

import (
	"fmt"
	"log"
	"os/exec"
)

func (fl *Service) InitGoModule(module string) {
	launchGoModInitAndGoModTidy(module)
}

func launchGoModInitAndGoModTidy(module string) {
	if err := exec.Command("go", "mod", "init", module).Run(); err != nil {
		log.Fatal("Error init", err)
	}

	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		log.Fatal("Error tidy ", err)
	}
	fmt.Println("Go mod init and Go mod tidy executed successfully")
}
