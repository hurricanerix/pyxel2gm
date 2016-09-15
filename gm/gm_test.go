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

package gm

import (
	"fmt"
	"testing"
)

func TestSplitPath(t *testing.T) {
	files := []string{
		"background\\images\\%s.png",
		"sprites\\images\\%s_0.png",
		"sprites\\images\\%s_10.png",
	}

	for _, f := range files {
		projectPath := "C:\\Users\\foo\\Documents\\GameMaker\\Projects\\Test.gmx"
		shortName := "spr_test"
		filename := fmt.Sprintf(f, shortName)
		path := fmt.Sprintf("%s\\%s", projectPath, filename)

		parts, err := SplitSpritePath(path)
		if err != nil {
			t.Errorf("(%s) error returned: %s", filename, err.Error())
		}
		if parts == nil {
			t.Errorf("(%s) SplitPath returned parts as nil", filename)
			return
		}
		if len(parts) != 2 {
			t.Errorf("(%s) invalid length of parts: %d, expected 2", filename, len(parts))
			return
		}
		if parts[0] != projectPath {
			t.Errorf("(%s) projectPath returned \"%s\", expected \"%s\"", filename, parts[0], projectPath)
		}
		if parts[1] != shortName {
			t.Errorf("(%s) filename returned \"%s\", expected \"%s\"", filename, parts[1], shortName)

		}
	}
}
