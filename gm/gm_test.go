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
	projectPath := "C:\\Users\\foo\\Documents\\GameMaker\\Projects\\Test.gmx"
	filename := "spr_test"
	path := fmt.Sprintf("%s\\sprites\\images\\%s_0.png", projectPath, filename)

	parts, err := SplitSpritePath(path)
	if err != nil {
		t.Errorf("error returned")
	}
	if parts == nil {
		t.Errorf("SplitPath returned parts as nil")
		return
	}
	if len(parts) != 2 {
		t.Errorf("invalid length of parts: %d, expected 2", len(parts))
		return
	}
	if parts[0] != projectPath {
		t.Errorf("projectPath returned \"%s\", expected \"%s\"", parts[0], projectPath)
	}
	if parts[1] != filename {
		t.Errorf("filename returned \"%s\", expected \"%s\"", parts[1], filename)

	}
}
