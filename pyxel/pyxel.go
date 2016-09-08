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

// Package pyxel handles Pyxel Edit specific interactions.
package pyxel

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

// ExportTiles from file at pyxelPath into projectPath.
func ExportTiles(pyxelPath, pyxelName, projectPath string) error {
	fullpath := fmt.Sprintf("%s\\%s.pyxel", pyxelPath, pyxelName)
	z, err := zip.OpenReader(fullpath)
	if err != nil {
		return err
	}
	defer z.Close()

	path := fmt.Sprintf("%s/sprites/images", projectPath)
	// TODO: Make the directory if it does not exist
	// err = os.MkdirAll(path, 0777)
	// if err != nil {
	// 	return err
	// }
	// Iterate through files in the archive.
	var tileName = regexp.MustCompile(`^tile([0-9]+)\.png$`)
	for _, zf := range z.File {
		if tileName.MatchString(zf.Name) {
			i := tileName.FindStringSubmatch(zf.Name)[1]
			name := fmt.Sprintf("%s_%s.png", pyxelName, i)
			filepath := fmt.Sprintf("%s/%s", path, name)

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

	return nil
}
