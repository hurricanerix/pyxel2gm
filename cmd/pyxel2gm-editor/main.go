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
	"image"
	"log"
	"os"
	"os/exec"

	"github.com/hurricanerix/pyxel2gm/gm"
	"github.com/hurricanerix/pyxel2gm/pyxel"
)

func main() {
	// Split the path provided into projectPath and name.
	if len(os.Args) != 2 {
		log.Fatal(fmt.Sprintf("Invalid number of arguments"))
		return
	}
	parts, err := gm.SplitSpritePath(os.Args[1])
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	projectPath := parts[0]
	imageDir := parts[1]
	shortName := parts[2]

	logfile := fmt.Sprintf("%s\\pyxel2gm-editor.log", projectPath)
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Find the .pyxel file associated with the provided name.
	assetsDir := fmt.Sprintf("assets\\%s", imageDir)
	searchPath := fmt.Sprintf("%s\\%s", projectPath, assetsDir)
	pyxelDir, err := pyxel.FindAsset(searchPath, shortName)
	imagePath := fmt.Sprintf("%s\\%s\\images", projectPath, imageDir)

	if _, notFound := err.(pyxel.FileNotFound); notFound {
		// A pyxel file was not found, create one from GM images.
		var tiles []*image.Image
		switch imageDir {
		case "sprites":
			// Look for tiles that fit the format '{shortName}_{N}'
			tiles, err = gm.GetImages(imagePath, fmt.Sprintf("%s_", shortName))
		case "background":
			tiles, err = gm.GetImages(imagePath, shortName)
		default:
			err = fmt.Errorf("invalid imageDir: %s", imageDir)
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(tiles) <= 0 {
			log.Fatal("no images found")
		}
		pyxelDir = searchPath
		err = pyxel.Create(pyxelDir, shortName, tiles)
		if err != nil {
			log.Fatal(err)
		}

	}

	// Open the .pyxel file with the default program (this should be Pyxel Edit).
	pyxelFile := fmt.Sprintf("%s\\%s.pyxel", pyxelDir, shortName)
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

	// Export images in the .pyxel file into the project as .png.
	switch imageDir {
	case "sprites":
		// Sprites should export the tileset as images.
		err = pyxel.ExportTiles(pyxelDir, imagePath, shortName)
	case "background":
		// Backgrounds should export the layers as a single image.
		err = pyxel.ExportLayers(pyxelDir, imagePath, shortName)
	default:
		err = fmt.Errorf("invalid imageDir: %s", imageDir)
	}
	if err != nil {
		log.Fatal(err)
	}
}
