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
	"flag"
	"fmt"
	"os"

	"github.com/hurricanerix/pyxel2gm/app"
)

var (
	assetsDir   string
	projectDir  string
	ignoreGMX   bool
	dry         bool
	showVersion bool
)

func init() {
	flag.StringVar(&assetsDir, "assets-dir", "assets", "directory to scan for .pyxel files")
	flag.StringVar(&projectDir, "project-dir", ".", "Game Maker project directory to export into")
	flag.BoolVar(&ignoreGMX, "ignore-gmx", false, "don't create or modify .gmx files")
	flag.BoolVar(&dry, "dry", false, "display report of files to be created")
	flag.BoolVar(&showVersion, "version", false, "display the version of this app")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println(app.Version)
		os.Exit(0)
	}

	// Create app context
	a := app.Context{
		AssetsDir:  assetsDir,
		ProjectDir: projectDir,
		IgnoreGMX:  ignoreGMX,
	}

	// Run app
	if err := a.Run(dry); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
