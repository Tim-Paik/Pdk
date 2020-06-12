package tree

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PrintPath(path string, indent string) (err error) {
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(path); err != nil {
		return fmt.Errorf("Error: Read file path error\n")
	}

	for i := 0; i < len(files); i++ {
		if files[i].Name()[0] == '.' {
			files = append(files[:i], files[i+1:]...)
		}
	}
	dirs := make([]string, 0)

	lenFile := len(dirs)

	for _, fi := range files {
		if fi.IsDir() {
			dirs = append(dirs, fi.Name())
		}
	}

	for i := 0; i < len(dirs); i++ {
		if i == len(dirs)-1 {
			fmt.Println(indent + "└── " + dirs[i])
			if i >= lenFile {
				if err := PrintPath(path+"\\"+dirs[i], indent+"   "); err != nil {
					return err
				}
			}
		} else {
			fmt.Println(indent + "├── " + dirs[i])
			if i >= lenFile {
				if err := PrintPath(path+"\\"+dirs[i], indent+"│  "); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
