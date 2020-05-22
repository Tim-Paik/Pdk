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
	DefaultRepo       = "repo.json"
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
