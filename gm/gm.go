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

// Package gm handles Game Maker Studio specific interactions.
package gm

import (
	"fmt"
	"image"
	_ "image/png" // register PNG format
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// SplitSpritePath and return the Game Maker Studio project path, and
// sprite filename (excluding the index and extension).
func SplitSpritePath(path string) ([]string, error) {
	// TODO: fix this to work with backgrounds.
	var pngPath = regexp.MustCompile(`^(.*)[\\/]sprites[\\/]images[\\/](.*)_[0-9]+.png$`)
	m := pngPath.FindStringSubmatch(path)
	// len should be 3 (full string match, path to project, short name of sprite)
	if len(m) != 3 {
		return nil, fmt.Errorf("invalid path: %s", path)
	}
	return []string{m[1], m[2]}, nil
}

// GetTiles from sprite returning a slice of Image pointers.
func GetTiles(projectPath, name string) ([]*image.Image, error) {
	tiles := []*image.Image{}
	// TODO: make this work for sprites and backgrounds
	p := fmt.Sprintf("%s\\sprites\\images", projectPath)
	fileList, err := ioutil.ReadDir(p)
	if err != nil {
		return tiles, nil
	}
	for _, f := range fileList {
		// If the file name is not prefixed with name, continue to the next file.
		if !strings.HasPrefix(f.Name(), name) {
			continue
		}
		fp := fmt.Sprintf("%s\\%s", p, f.Name())
		img, err := readImage(fp)
		if err != nil {
			return nil, err
		}
		tiles = append(tiles, img)
	}
	return tiles, nil
}

// readImage from filepath and return Image pointer.
func readImage(filepath string) (*image.Image, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}
	return &img, nil
}
