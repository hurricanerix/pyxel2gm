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
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Version of app
const Version = "0.0.1"

// Context of app.
type Context struct {
	AssetsDir  string
	ProjectDir string
	IgnoreGMX  bool
}

// Run the app
func (ctx Context) Run(dry bool) error {
	// Get a list of all pyxel files under the assets dir.
	files, err := getFiles(ctx.AssetsDir, ".pyxel")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Error: %s\n", err)
		}
		return err
	}

	// For every pyxel file found:
	for _, f := range files {
		//   * Open pyxel/docData.json to get name
		//     - expName = ['settings']['ExportTilesetPanel_prefFileName']
		//     - gmxFilename = './sprites/$expName.sprite.gmx'
		//   * Check for a corresponding gmx file
		//     - If found, read and modify if needed
		//     - Otherwise, write a new file
		//   * List files in pyxel
		//     - for every file prefixed with 'title'
		//       * re.compile('tile([0-9]+)\.png')
		//       * copy to sprites/images/$expName_$num.png
		fmt.Println(f)
	}

	return nil
}

func getFiles(p, ext string) ([]string, error) {
	var files []string

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
			files = append(subFiles)
		} else if strings.HasSuffix(f.Name(), ext) {
			files = append(files, fullpath)
		}
	}

	return files, nil
}
