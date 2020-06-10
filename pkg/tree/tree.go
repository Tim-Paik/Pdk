package tree

import (
	"fmt"
	"io/ioutil"
)

func PrintPath(path string, indent string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error: Read file path error\n", err)
		return
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
				PrintPath(path+"\\"+dirs[i], indent+"   ")
			}
		} else {
			fmt.Println(indent + "├── " + dirs[i])
			if i >= lenFile {
				PrintPath(path+"\\"+dirs[i], indent+"│  ")
			}
		}
	}
}
