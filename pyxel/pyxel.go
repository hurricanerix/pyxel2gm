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
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// FileNotFound when searching for pyxel file.
type FileNotFound struct {
	MSG        string
	SearchPath string
	ShortName  string
}

func (f FileNotFound) Error() string {
	return fmt.Sprintf("pyxel2gm: could not locate %s under %s", f.ShortName, f.SearchPath)
}

// Create pyxel file from tiles.
func Create(createPath, shortName string, tiles []*image.Image) error {
	if len(tiles) < 1 {
		return fmt.Errorf("can not create image with no tiles")
	}

	// Create the archive.
	// TODO: use path utils for this.
	filepath := fmt.Sprintf("%s\\%s.pyxel", createPath, shortName)
	fd, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fd.Close()
	w := zip.NewWriter(fd)

	// Add Tiles to archive.
	tileCount, tileWidth, tileHeight, err := createTiles(w, tiles)
	if err != nil {
		return err
	}

	// Add Layers to archive.
	layerCount, layerWidth, layerHeight, err := createLayers(w, tiles)
	if err != nil {
		return err
	}

	// Add docData.json to archive.
	err = createDocData(w, tileCount, tileWidth, tileHeight, layerCount, layerWidth, layerHeight)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

// createLayers by adding a blank png of widthxheight to the zip archive.
func createLayers(w *zip.Writer, tiles []*image.Image) (int, int, int, error) {
	filename := fmt.Sprintf("layer0.png")
	count := 1
	tileWidth := 0
	tileHeight := 0
	width := 0
	height := 0
	for i, img := range tiles {
		if i == 0 {
			b := (*img).Bounds()
			tileWidth = b.Min.X + b.Max.X
			tileHeight = b.Min.Y + b.Max.Y
		}

		if len(tiles) == 1 {
			width = tileWidth
			height = tileHeight
		} else {
			width = tileWidth * 5
			height = tileHeight * 5
		}

		f, err := w.Create(filename)
		if err != nil {
			return 0, 0, 0, err
		}
		imgRect := image.Rect(0, 0, width, height)
		img := image.NewRGBA(imgRect)
		err = png.Encode(f, img)
		if err != nil {
			return 0, 0, 0, err
		}
	}
	return count, width, height, nil
}

func createDocData(w *zip.Writer, tileCount, tileWidth, tileHeight, layerCount, layerWidth, layerHeight int) error {
	filename := fmt.Sprintf("docData.json")
	f, err := w.Create(filename)
	if err != nil {
		return err
	}
	data := fmt.Sprintf(`{
	  "animations": {
	    "0": {
	      "name": "Animation 1",
	      "length": 3,
	      "baseTile": 0,
	      "frameDurationMultipliers": [
	        100,
	        100,
	        100
	      ],
	      "frameDuration": 100
	    }
	  },
	  "settings": {
	    "ExportTilesetPanel_prefFormat": "0",
	    "ExportTilesetPanel_prefTranspMatteColor": "4278190080",
	    "ExportTilesetPanel_prefPath": "C:\\Users\\hurricanerix\\Desktop\\melee",
	    "ExportTilesetPanel_prefOverwrite": "false",
	    "ExportTilesetPanel_prefSeparateFiles": "true",
	    "ExportTilesetPanel_prefFileName": "Untitled",
	    "ExportTilesetPanel_prefTilePadding": "0",
	    "ExportTilesetPanel_prefScaling": "1"
	  },
	  "canvas": {
	    "width": %d,
	    "layers": {
	      "0": {
	        "name": "Layer 0",
	        "alpha": 255,
	        "hidden": false,
	        "blendMode": "normal",
	        "tileRefs": {
	          "0": {
	            "rot": 0,
	            "index": 0,
	            "flipX": false
	          },
	          "1": {
	            "rot": 0,
	            "index": 1,
	            "flipX": false
	          },
	          "2": {
	            "rot": 0,
	            "index": 2,
	            "flipX": false
	          },
	          "8": {
	            "rot": 0,
	            "index": 0,
	            "flipX": false
	          },
	          "9": {
	            "rot": 0,
	            "index": 1,
	            "flipX": false
	          },
	          "10": {
	            "rot": 0,
	            "index": 2,
	            "flipX": false
	          },
	          "17": {
	            "rot": 0,
	            "index": 0,
	            "flipX": false
	          }
	        }
	      }
	    },
	    "height": %d,
	    "tileWidth": %d,
	    "tileHeight": %d,
	    "numLayers": %d
	  },
	  "name": "Untitled",
	  "tileset": {
	    "numTiles": %d,
	    "tileHeight": %d,
	    "tilesWide": 8,
	    "fixedWidth": true,
	    "tileWidth": %d
	  },
	  "palette": {
	    "colors": {
	      "0": "ff000000",
	      "1": "ffffffff",
	      "2": "ff9d9d9d",
	      "3": "ffe06f8b",
	      "4": "ffbe2633",
	      "5": "ff493c2b",
	      "6": "ffa46422",
	      "7": "ffeb8931",
	      "8": null,
	      "9": null,
	      "10": null,
	      "11": null,
	      "12": "fff7e26b",
	      "13": "ffa3ce27",
	      "14": "ff44891a",
	      "15": "ff2f484e",
	      "16": "ff1b2632",
	      "17": "ff005784",
	      "18": "ff31a2f2",
	      "19": "ffb2dcef"
	    },
	    "numColors": 20,
	    "width": 12,
	    "height": 5
	  },
	  "version": "0.4.2"
	}`, layerWidth, layerHeight, tileWidth, tileHeight, layerCount, tileCount, tileHeight, tileWidth)
	f.Write([]byte(data))
	return nil
}

// createTiles by adding them to the zip archive and return tile count.
func createTiles(w *zip.Writer, tiles []*image.Image) (int, int, int, error) {
	count := 0
	width := 0
	height := 0
	for i, img := range tiles {
		if i == 0 {
			b := (*img).Bounds()
			width = b.Min.X + b.Max.X
			height = b.Min.Y + b.Max.Y
		}
		filename := fmt.Sprintf("tile%d.png", i)
		f, err := w.Create(filename)
		if err != nil {
			return 0, 0, 0, err
		}
		err = png.Encode(f, *img)
		if err != nil {
			return 0, 0, 0, nil
		}
		count++
	}
	return count, width, height, nil
}

// FindAsset shortName in searchPath and return the full path to the directory
// containing it.
func FindAsset(searchPath, shortName string) (string, error) {
	// List files in searchPath
	fileList, err := ioutil.ReadDir(searchPath)
	if err != nil {
		return "", err
	}
	for _, f := range fileList {
		// TODO: use path utils for this.
		filepath := fmt.Sprintf("%s\\%s", searchPath, f.Name())
		if !f.IsDir() && fmt.Sprintf("%s.pyxel", shortName) == f.Name() {
			// File matches, return searchPath
			return searchPath, nil
		} else if f.IsDir() {
			found, err := FindAsset(filepath, shortName)
			if _, ok := err.(FileNotFound); ok {
				// Sub directory does not contain a match, continue on to the next file.
				continue
			} else if err != nil {
				// Non-Search related error returned, return it.
				return "", err
			}
			if found != "" {
				// A match was found, return it.
				return found, nil
			}
		}
	}
	return "", FileNotFound{SearchPath: searchPath, ShortName: shortName}
}

// ExportTiles from file at pyxelPath into projectPath.
func ExportTiles(pyxelPath, imagePath, shortName string) error {
	// TODO: pass assetPath to this.
	fullpath := fmt.Sprintf("%s\\%s.pyxel", pyxelPath, shortName)
	z, err := zip.OpenReader(fullpath)
	if err != nil {
		return err
	}
	defer z.Close()

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
			name := fmt.Sprintf("%s_%s.png", shortName, i)
			filepath := fmt.Sprintf("%s\\%s", imagePath, name)

			gmf, err := os.Create(filepath)
			if err != nil {
				log.Println(err)
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

// ExportLayers from file at pyxelPath into projectPath.
func ExportLayers(pyxelPath, imagePath, shortName string) error {
	// TODO: pass assetPath to this.
	fullpath := fmt.Sprintf("%s\\%s.pyxel", pyxelPath, shortName)
	z, err := zip.OpenReader(fullpath)
	if err != nil {
		return err
	}
	defer z.Close()

	// TODO: Make the directory if it does not exist
	// err = os.MkdirAll(path, 0777)
	// if err != nil {
	// 	return err
	// }

	// Iterate through files in the archive.=
	var tileName = regexp.MustCompile(`^layer0.png$`)
	for _, zf := range z.File {
		if tileName.MatchString(zf.Name) {
			name := fmt.Sprintf("%s.png", shortName)
			filepath := fmt.Sprintf("%s\\%s", imagePath, name)

			gmf, err := os.Create(filepath)
			if err != nil {
				log.Println(err)
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
