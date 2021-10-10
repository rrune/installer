package installer

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	. "github.com/rrune/installer/util"
)

type installer struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
	Temp     string `json:"temp"`
	Dest     string
}

func (i installer) SetDest(dest string) {
	i.Dest = dest
}

func (i installer) Install() {
	Check(i.download())
	Check(i.unzip())
	os.RemoveAll(i.Temp)
}

func (i installer) download() error {
	var err error
	// Create temp dir
	err = os.Mkdir("temp", 0666)
	if err != nil {
		return err
	}
	// Create the file
	out, err := os.Create(i.Temp + i.Filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(i.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (i installer) unzip() error {

	r, err := zip.OpenReader(i.Temp + i.Filename)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(i.Dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(i.Dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func New() *installer {
	return &installer{}
}
