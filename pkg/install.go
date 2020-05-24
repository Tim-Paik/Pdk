package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Install(packages []string, repoName string) {
	fmt.Println("Install called!")
	var repo Repositories
	var index []int
	if repo, err = Read(RepoRoot + "/" + repoName + ".json"); err != nil {
		fmt.Println(err)
		return
	}
	if index, err = Search(&repo, packages); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(index)
	return
}

func Search(repo *Repositories, packages []string) (index []int, err error) {
	index = make([]int, len(packages))
	for ix := 0; ix < len(packages); ix++ {
		for i := 0; i < len(repo.Pkgs); i++ {
			if repo.Pkgs[i].Name == packages[ix] {
				index[ix] = i
			}
		}
	}
	return index, nil
}

func Download(URL, PATH string) (written int64, err error) {
	var (
		resp *http.Response
		data *os.File
	)

	if resp, err = http.Get(URL); err != nil {
		fmt.Println("Error: Sending request error")
		return -1, err
	}

	if data, err = os.Create(PATH); err != nil {
		return -2, err
	}

	if written, err = io.Copy(data, resp.Body); err != nil {
		return -3, err
	}

	if err := resp.Body.Close(); err != nil {
		return -4, err
	}

	if err := data.Close(); err != nil {
		return -5, err
	}

	return written, nil
}
