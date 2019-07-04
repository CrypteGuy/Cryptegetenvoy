// Copyright 2019 Tetrate
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

package manifest

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrint(t *testing.T) {
	tests := []struct {
		name           string
		wantOutputFile string
	}{
		{
			name:           "Prints golden output",
			wantOutputFile: "list.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bytes.NewBuffer(nil)
			PrettyPrint(got, goodManifest())
			want, _ := ioutil.ReadFile(filepath.Join("testdata", tt.wantOutputFile))
			assert.Equal(t, string(want), got.String())
		})
	}
}
