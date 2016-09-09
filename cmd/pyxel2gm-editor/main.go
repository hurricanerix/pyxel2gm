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

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hurricanerix/pyxel2gm/gm"
	"github.com/hurricanerix/pyxel2gm/pyxel"
)

func main() {
	// Split the path provided into projectPath and name.
	if len(os.Args) != 2 {
		return
	}
	parts, err := gm.SplitSpritePath(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	projectPath := parts[0]
	name := parts[1]

	// Find the .pyxel file associated with the provided name.
	assetsDir := "assets"
	// TODO: change FindAssets to take assetsPath instead of projectPath + assetsDir.
	filepath, err := pyxel.FindAsset(projectPath, assetsDir, name)
	if _, ok := err.(pyxel.FileNotFound); ok {
		// TODO: GetTiles might need to know if this is a sprite or background.
		tiles, err := gm.GetTiles(projectPath, name)
		if err != nil {
			log.Fatal(err)
		}
		// TODO: only pass path to desired pyxel file and tiles.
		pyxel.Create(projectPath, assetsDir, name, tiles)
	} else if err != nil {
		log.Fatal(err)
	}

	// Open the .pyxel file with the default program (this should be Pyxel Edit).
	fullpath := fmt.Sprintf("%s\\%s.pyxel", filepath, name)
	openCMD := fmt.Sprintf("/K %s", fullpath)
	cmd := exec.Command("cmd", openCMD)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// Export tiles in the .pyxel file into the project as .png.
	err = pyxel.ExportTiles(filepath, name, projectPath)
	if err != nil {
		log.Fatal(err)
	}
}
