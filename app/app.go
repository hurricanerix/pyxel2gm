// Copyright 2016 Richard Hawkins (hurricanerix@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Version of app
const Version = "0.0.1"

// Context of app.
type Context struct {
	AssetsDir  string
	ProjectDir string
}

type FileInfo struct {
	Path string // Path to the file.
	Name string // Name of file, excluding the extension.
}

// Run the app
func (ctx Context) Run() error {
	// Get a list of all pyxel files under the assets dir.
	files, err := getFiles(ctx.AssetsDir, ".pyxel")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Error: %s\n", err)
		}
		return err
	}

	var tileName = regexp.MustCompile(`^tile([0-9]+)\.png$`)

	// For every pyxel file found:
	for _, f := range files {
		filepath := fmt.Sprintf("%s/%s.pyxel", f.Path, f.Name)
		z, err := zip.OpenReader(filepath)
		if err != nil {
			// TODO: handle error better than this
			continue
		}
		defer z.Close()

		path := fmt.Sprintf("%s/sprites/images", ctx.ProjectDir)
		//err = os.MkdirAll(path, 0777)
		// if err != nil {
		// 	return err
		// }
		// Iterate through files in the archive.
		for _, zf := range z.File {
			if tileName.MatchString(zf.Name) {
				i := tileName.FindStringSubmatch(zf.Name)[1]
				name := fmt.Sprintf("%s_%s.png", f.Name, i)
				filepath := fmt.Sprintf("%s/%s", path, name)
				println(filepath)
				gmf, err := os.Create(filepath)
				if err != nil {
					// TODO: better error handling.
					fmt.Println(err)
					continue
				}

				rc, err := zf.Open()
				if err != nil {
					log.Fatal(err)
				}

				_, err = io.Copy(gmf, rc)
				if err != nil {
					log.Fatal(err)
				}
				rc.Close()

				gmf.Close()
			}
		}
	}

	return nil
}

func getFiles(p, ext string) ([]FileInfo, error) {
	var files []FileInfo
	fileList, err := ioutil.ReadDir(p)
	if err != nil {
		return files, err
	}

	for _, f := range fileList {
		fullpath := path.Join(p, f.Name())
		if f.IsDir() {
			subFiles, err := getFiles(fullpath, ext)
			if err != nil {
				return files, err
			}
			if subFiles != nil {
				files = append(subFiles)
			}
		} else if strings.HasSuffix(f.Name(), ext) {
			var extension = filepath.Ext(f.Name())
			var name = f.Name()[0 : len(f.Name())-len(extension)]
			files = append(files, FileInfo{Path: p, Name: name})
		}
	}
	return files, nil
}
