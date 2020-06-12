package pkg

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Pack() (err error) {
	var (
		pdkgJson []byte
		pdkg     Pdkg
		pkgJson  []byte
		pkg      Pkg
	)

	if pdkgJson, err = ioutil.ReadFile(PdkgJson); err != nil {
		return err
	}

	if err := json.Unmarshal(pdkgJson, &pdkg); err != nil {
		return err
	}

	fmt.Printf("Packing %s version %s\n", pdkg.Name, pdkg.Version)

	fmt.Println(Indent1 + "Files to delete when uninstalling:")

	for _, file := range pdkg.Files {
		fmt.Println(Indent2 + file)
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

	if err := os.MkdirAll("packages/", 0755); err != nil {
		return err
	}

	tarName := "packages/" + pdkg.Name + "-" + pdkg.Version + ".tar"
	if tarName, err = filepath.Abs(tarName); err != nil {
		return err
	}

	fmt.Println(Indent1 + "Packing...")

	if err := Tar("apps/", tarName); err != nil {
		return err
	}

	if err := os.RemoveAll(tarName + ".zst"); err != nil {
		return err
	}
	var outInfo bytes.Buffer
	cmd := exec.Command("zstd", tarName)
	cmd.Stdout = &outInfo
	fmt.Println(Indent1 + "Compress package...")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: Please install ZSTD to package the software")
		fmt.Println("Run: pdk install zstd")
		return err
	}
	if err := os.RemoveAll(tarName); err != nil {
		return err
	}

	fmt.Println(Indent1 + "Format JSON...")
	pkg.Name = pdkg.Name
	pkg.Version = pdkg.Version
	pkg.Description = pdkg.Description
	pkg.Update = pdkg.Update
	if pkg.Md5, err = Md5Sum(tarName + ".zst"); err != nil {
		return err
	}
	pkg.URL = ""

	if pkgJson, err = json.Marshal(pkg); err != nil {
		return err
	}

	if err = ioutil.WriteFile("packages/"+pkg.Name+".json", pkgJson, 0644); err != nil {
		return err
	}

	fmt.Println("Done!")
	return nil
}

func Tar(src, dst string) (err error) {
	var (
		fw *os.File
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

	tw = tar.NewWriter(fw)

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

		hdr.Name = strings.TrimPrefix(strings.TrimPrefix(fileName, "apps"), string(filepath.Separator))

		if hdr.Name == "/" {
			return nil
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		fmt.Println(Indent2 + "Added file: " + hdr.Name + " " + strconv.FormatInt(hdr.Size, 10))
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
