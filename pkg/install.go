package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
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
	if _, err := Download("", "", "nil"); err != nil {
		fmt.Println("Error: Download error!")
		return
	}
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

func Download(URL string, PATH string, MD5string string) (written int64, err error) {
	var (
		resp     *http.Response
		data     *os.File
		checkMD5 hash.Hash
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

	checkMD5 = md5.New()
	if file, err := os.Open(PATH); err != nil {
		return 0, err
	} else {
		if _, err := io.Copy(checkMD5, file); err != nil {
			return 0, err
		}
	}

	if MD5string != "nil" && MD5string != hex.EncodeToString(checkMD5.Sum(nil)) {
		fmt.Println("Error: MD5 error")
	} else {
		fmt.Println("MD5 OK!")
	}

	if err := resp.Body.Close(); err != nil {
		return -5, err
	}

	if err := data.Close(); err != nil {
		return -6, err
	}

	return written, nil
}
