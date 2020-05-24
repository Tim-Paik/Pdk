package pkg

import (
	"os"
	"path/filepath"
)

type Repositories struct {
	Name    string `json:"name"`
	Version int64  `json:"version"`
	Update  int64  `json:"update"`
	URL     string `json:"URL"`
	Pkgs    []struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description"`
		Update      int64  `json:"update"`
		URL         string `json:"URL"`
	} `json:"pkgs"`
}

var (
	Root              = root()
	RepoRoot          = repoRoot()
	PackageRoot       = packageRoot()
	AppRoot           = appRoot()
	DefaultRepo       = "repo"
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
	if err := os.MkdirAll(Root+"/apps", 0755); err != nil {
		panic(err)
	}
	return Root + "/apps"
}
