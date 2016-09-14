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

	/*

		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			t.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("This is a test log entry")
	*/

	// Split the path provided into projectPath and name.
	if len(os.Args) != 2 {
		return
	}
	parts, err := gm.SplitSpritePath(os.Args[1])
	if err != nil {
		panic(err)
	}
	projectPath := parts[0]
	shortName := parts[1]
	logfile := fmt.Sprintf("%s\\pyxel2gm-editor.log", projectPath)
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Find the .pyxel file associated with the provided name.
	assetsDirName := "assets"
	searchPath := fmt.Sprintf("%s\\%s", projectPath, assetsDirName)
	pyxelDir, err := pyxel.FindAsset(searchPath, shortName)
	if _, ok := err.(pyxel.FileNotFound); ok {
		// TODO: GetTiles might need to know if this is a sprite or background.
		tiles, err := gm.GetTiles(projectPath, shortName)
		if err != nil {
			log.Fatal(err)
		}
		pyxelDir = searchPath
		err = pyxel.Create(pyxelDir, shortName, tiles)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	// Open the .pyxel file with the default program (this should be Pyxel Edit).
	pyxelFile := fmt.Sprintf("%s\\%s.pyxel", pyxelDir, shortName)
	log.Println(pyxelFile)
	openCMD := fmt.Sprintf("/K %s", pyxelFile)
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
	err = pyxel.ExportTiles(pyxelDir, shortName, projectPath)
	if err != nil {
		log.Fatal(err)
	}
}
