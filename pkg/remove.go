package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func Remove(packages []string) (err error) {
	if len(packages) == 0 {
		return fmt.Errorf("Error: no targets specified\n")
	}
	for _, packageName := range packages {
		var (
			pdkg     Pdkg
			pdkgJson []byte
		)
		if pdkgJson, err = ioutil.ReadFile(AppData + "/" + packageName + "/" + PdkgJson); err != nil {
			return err
		}
		if err = json.Unmarshal(pdkgJson, &pdkg); err != nil {
			return err
		}
		fmt.Println(Indent1 + "Removing " + pdkg.Name + " v" + pdkg.Version)
		//CALLBACK_SCRIPT
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command(AppData + "/" + packageName + "/uninstall")
			out, _ := cmd.CombinedOutput()
			if len(out) != 0 {
				fmt.Println(string(out))
			}
		}
		for _, file := range pdkg.Files {
			if err := os.Remove(AppRoot + "/" + file); err != nil {
				return err
			}
			fmt.Println(Indent2 + "Deleted " + file)
		}
	}
	return nil
}
