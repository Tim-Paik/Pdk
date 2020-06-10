package pkg

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	gzip "github.com/klauspost/pgzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Pack() (err error) {
	var (
		pdkgJson []byte
		pdkg     Pdkg
	)

	if pdkgJson, err = ioutil.ReadFile(PdkgJson); err != nil {
		return err
	}

	if err := json.Unmarshal(pdkgJson, &pdkg); err != nil {
		return err
	}

	fmt.Printf("Packing %s version %s\n", pdkg.Name, pdkg.Version)

	fmt.Println("Files to delete when uninstalling:")

	for _, file := range pdkg.Files {
		fmt.Println(file)
	}

	pdkg.Update = time.Now().Unix()

	if pdkgJson, err = json.Marshal(pdkg); err != nil {
		return err
	}

	if err := os.MkdirAll("apps/appData/"+pdkg.Name, 0755); err != nil {
		return err
	}

	if err = ioutil.WriteFile("apps/appData/"+pdkg.Name+"/"+PdkgJson, pdkgJson, 0644); err != nil {
		return err
	}

	if err := Tar("apps/", "packages/"+pdkg.Name+"-"+pdkg.Version+".pdkg"); err != nil {
		return err
	}

	return nil
}

func Tar(src, dst string) (err error) {
	var (
		fw *os.File
		gw *gzip.Writer
		tw *tar.Writer
	)

	if fw, err = os.Create(dst); err != nil {
		return err
	}
	defer func() {
		if err := fw.Close(); err != nil {
			return
		}
	}()

	gw = gzip.NewWriter(fw)
	defer func() {
		if err := gw.Close(); err != nil {
			return
		}
	}()

	tw = tar.NewWriter(gw)

	if err := filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var (
			hdr *tar.Header
			fr  *os.File
		)

		if hdr, err = tar.FileInfoHeader(fi, ""); err != nil {
			return err
		}

		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		if fr, err = os.Open(fileName); err != nil {
			return err
		}
		defer func() {
			if err := fr.Close(); err != nil {
				return
			}
		}()

		if _, err := io.Copy(tw, fr); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
