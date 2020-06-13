package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Repositories struct {
	Name    string `json:"name"`
	Version int64  `json:"version"`
	Update  int64  `json:"update"`
	OS      string `json:"os"`
	Arch    string `json:"arch"`
	Base    string `json:"base"`
	URL     string `json:"URL"`
	Pkgs    []Pkg  `json:"pkgs"`
}

type Pkg struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Update      int64  `json:"update"`
	Md5         string `json:"md5"`
	URL         string `json:"URL"`
}

type Pdkg struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Update      int64    `json:"update"`
	Files       []string `json:"files"`
}

var (
	Root        = root()
	RepoRoot    = repoRoot()
	PackageRoot = packageRoot()
	AppRoot     = appRoot()
	//AppPath           = appPath()
	AppData           = appData()
	DefaultRepo       = "repo"
	PdkgJson          = "pdkg.json"
	Indent1           = "==> "
	Indent2           = "  -> "
	err         error = nil
)

func root() (root string) {
	if root, err = os.Executable(); err != nil {
		panic(err)
	}
	if root, err = filepath.EvalSymlinks(root); err != nil {
		panic(err)
	}
	root = filepath.Dir(root)
	return root
}

func repoRoot() (repoRoot string) {
	if err := os.MkdirAll(Root+"/repo", 0644); err != nil {
		panic(err)
	}
	return Root + "/repo"
}

func packageRoot() (repoRoot string) {
	if err := os.MkdirAll(Root+"/packages", 0644); err != nil {
		panic(err)
	}
	return Root + "/packages"
}

func appRoot() (repoRoot string) {
	switch runtime.GOOS {
	case "windows":
		if err := os.MkdirAll(Root+"/apps", 0755); err != nil {
			panic(err)
		}
		repoRoot = Root + "/apps"
	default:
		repoRoot = "/"
	}
	return repoRoot
}

func appData() (appData string) {
	if err := os.MkdirAll(Root+"/apps/appData", 0755); err != nil {
		panic(err)
	}
	appData = Root + "/apps/appData"
	appData = Root + "/apps/appData"
	return appData
}

func CheckArch(repo *Repositories) (err error) {
	if repo.Arch == runtime.GOARCH && repo.OS == runtime.GOOS {
		return nil
	}
	if _, err := fmt.Printf("Warn: You are using %s instead of %s in %s\n", repo.OS+"/"+repo.Arch, runtime.GOOS+"/"+
		runtime.GOARCH, repo.Name); err != nil {
		return err
	}
	fmt.Print("\n" + Indent1 + "Is it mandatory to continue? (This may break your system) [Y/n]")
	var yesOrNo string
	if _, err := fmt.Scanln(&yesOrNo); err != nil {
		return err
	}
	if !(yesOrNo == "Y" || yesOrNo == "y") {
		fmt.Println("The action is canceled by the user.")
		return nil
	}
	return nil
}

func CheckRoot() {
	if runtime.GOOS != "windows" && os.Getuid() != 0 {
		fmt.Println("Error: you cannot perform this operation unless you are root.")
		os.Exit(1)
	}
}
