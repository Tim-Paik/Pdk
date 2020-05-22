package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Read(path string) (repo Repositories, err error) {
	var repoJson []byte

	repoJson, err = ioutil.ReadFile(path)
	if err != nil {
		return repo, err
	}

	err = json.Unmarshal(repoJson, &repo)
	if err != nil {
		return repo, err
	}

	return repo, err
}

func UpdateLocal(repoName string) (err error) {
	var repo Repositories
	if repo, err = Read(RepoRoot + "/" + repoName + ".json"); err != nil {
		fmt.Println(repoName + ".json was not found!")
		return err
	}
	if err := Update(repo.URL, repo.Name); err != nil {
		return err
	}
	return nil
}

func Update(url, repoName string) (err error) {
	var (
		repoJson []byte
		repo     Repositories
		data     []byte
		resp     *http.Response
	)

	if resp, err = http.Get(url); err != nil {
		fmt.Println("Error: URL parsing error")
		return err
	}

	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	if err = resp.Body.Close(); err != nil {
		return err
	}

	if err = json.Unmarshal(data, &repo); err != nil {
		return err
	}

	repo.Update = time.Now().Unix()

	if repoJson, err = json.Marshal(repo); err != nil {
		return err
	}

	if err = ioutil.WriteFile(RepoRoot+"/"+repoName+".json", repoJson, 0644); err != nil {
		return err
	}

	return nil
}
