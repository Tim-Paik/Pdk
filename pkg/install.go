package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Install(packages []string, repoName string) {
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
	for _, i := range index {
		PATH := PackageRoot + "/" + repo.Pkgs[i].Name + "-" + repo.Pkgs[i].Version
		if !Exists(PATH) {
			if _, err := Download(repo.Pkgs[i].URL, PATH); err != nil {
				fmt.Println(err)
				return
			}
		}
		if MD5, err := Md5Sum(PATH); err != nil {
			fmt.Println(err)
			return
		} else if MD5 != repo.Pkgs[i].Md5 {
			fmt.Println("Error: Md5 error")
			fmt.Println(MD5)
			fmt.Println(repo.Pkgs[i].Md5)
			if err := os.RemoveAll(PATH); err != nil {
				fmt.Println(err)
				return
			}
			if _, err := Download(repo.Pkgs[i].URL, PATH); err != nil {
				fmt.Println(err)
				return
			}
			if MD5, err := Md5Sum(PATH); err != nil {
				fmt.Println(err)
				return
			} else if MD5 != repo.Pkgs[i].Md5 {
				fmt.Println("Error: Serious md5 error")
				return
			}
		}
	}
	return
}

func Search(repo *Repositories, packages []string) (index []int, err error) {
	for _, packagesName := range packages {
		var pkgIndex []int
		for i, finish := 0, false; finish == false; i++ {
			if repo.Pkgs[i].Name == packagesName {
				pkgIndex = append(pkgIndex, i)
				finish = true
			} else if i == (len(repo.Pkgs) - 1) {
				return nil, fmt.Errorf("Error: Package %s not found\n", packagesName)
			}
		}
		index = append(index, pkgIndex...)
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

func Md5Sum(filePath string) (MD5string string, err error) {
	checkMD5 := md5.New()

	if file, err := os.Open(filePath); err != nil {
		return "", err
	} else {
		if _, err := io.Copy(checkMD5, file); err != nil {
			return "", err
		}
	}

	MD5string = hex.EncodeToString(checkMD5.Sum(nil))
	return MD5string, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
