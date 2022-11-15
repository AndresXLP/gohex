package app

import (
	"fmt"
	"os/exec"
)

func (s *Service) InitGoModule(module string) error {
	res, err := launchGoModInitAndGoModTidy(module)
	if err != nil {
		return fmt.Errorf("error in executing: %s", string(res))
	}
	return nil
}

func launchGoModInitAndGoModTidy(module string) ([]byte, error) {
	if module == "" {
		res, err := exec.Command("go", "mod", "init").CombinedOutput()
		if err != nil {
			return res, err
		}
	} else {
		res, err := exec.Command("go", "mod", "init", module).CombinedOutput()
		if err != nil {
			return res, err
		}
	}
	res, err := exec.Command("go", "mod", "tidy").CombinedOutput()
	if err != nil {
		return res, err
	}

	return res, nil
}
