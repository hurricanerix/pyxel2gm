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
	var spritePath = regexp.MustCompile(`^(.*)[\\/](sprites)[\\/]images[\\/](.*)_[0-9]+.png$`)
	m := spritePath.FindStringSubmatch(path)
	// len should be 3 (full string match, projectPath, shortName)
	if len(m) != 4 {
		// Provided path does not match a sprite, check for a background.
		var bgPath = regexp.MustCompile(`^(.*)[\\/](background)[\\/]images[\\/](.*).png$`)
		m = bgPath.FindStringSubmatch(path)
		if len(m) != 4 {
			return nil, fmt.Errorf("path does not match: %s", path)
		}
	}
	projectPath := m[1]
	imageDir := m[2]
	shortName := m[3]
	return []string{projectPath, imageDir, shortName}, nil
}

// GetImages from sprite returning a slice of Image pointers.
func GetImages(imagesPath, name string) ([]*image.Image, error) {
	tiles := []*image.Image{}
	fileList, err := ioutil.ReadDir(imagesPath)
	if err != nil {
		return tiles, nil
	}
	for _, f := range fileList {
		// If the file name is not prefixed with name, continue to the next file.
		if !strings.HasPrefix(f.Name(), name) {
			continue
		}
		fp := fmt.Sprintf("%s\\%s", imagesPath, f.Name())
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
