package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Remove(packages []string) (err error) {
	if len(packages) == 0 {
		return fmt.Errorf("Error: no targets specified\n")
	}
	for _, packageName := range packages {
		var pdkg Pdkg
		var pdkgJson []byte
		if pdkgJson, err = ioutil.ReadFile(AppData + "/" + packageName + "/" + PdkgJson); err != nil {
			return err
		}
		if err = json.Unmarshal(pdkgJson, &pdkg); err != nil {
			return err
		}
		for _, file := range pdkg.Files {
			if PATH, err := filepath.Abs(file); err != nil {
				return err
			} else if err := os.Remove(PATH); err != nil {
				return err
			}
		}
	}
	return nil
}
