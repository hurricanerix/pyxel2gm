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
	"io/ioutil"
	"regexp"
)

// SplitSpritePath and return the Game Maker Studio project path, and
// sprite filename (excluding the index and extension).
func SplitSpritePath(path string) ([]string, error) {
	var pngPath = regexp.MustCompile(`^(.*)[\\/]sprites[\\/]images[\\/](.*)_[0-9]+.png$`)
	m := pngPath.FindStringSubmatch(path)
	if len(m) != 3 {
		return nil, fmt.Errorf("invalid path: %s", path)
	}
	return []string{m[1], m[2]}, nil
}

// FindAsset in assetsDir and return the full path to the directory containing it.
func FindAsset(projectPath, assetsDir, name string) (string, error) {
	p := fmt.Sprintf("%s\\%s", projectPath, assetsDir)
	filepath, err := findAsset(p, name)
	if err != nil {
		return "", err
	}
	if filepath == "" {
		return "", fmt.Errorf("file not found: %s.pyxel", name)
	}
	return filepath, nil
}

// findAsset by recursivly searching through directories and compairing names.
func findAsset(p, name string) (string, error) {
	fileList, err := ioutil.ReadDir(p)
	if err != nil {
		return "", err
	}
	for _, f := range fileList {
		fullpath := fmt.Sprintf("%s\\%s", p, f.Name())
		if f.IsDir() {
			found, err := findAsset(fullpath, name)
			if err != nil {
				return "", err
			}
			if found != "" {
				return found, nil
			}
		} else if fmt.Sprintf("%s.pyxel", name) == f.Name() {
			return p, nil
		}
	}
	return "", nil
}
