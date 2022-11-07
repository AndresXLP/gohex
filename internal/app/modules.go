package app

import (
	"fmt"
	"os/exec"
)

func (s *Service) InitGoModule() error {
	res, err := launchGoModInitAndGoModTidy()
	if err != nil {
		return fmt.Errorf("error in executing: %s", string(res))
	}
	return nil
}

func launchGoModInitAndGoModTidy() ([]byte, error) {
	res, err := exec.Command("go", "mod", "init").CombinedOutput()
	if err != nil {
		return res, err
	}
	res, err = exec.Command("go", "mod", "tidy").CombinedOutput()
	if err != nil {
		return res, err
	}

	return res, nil
}
