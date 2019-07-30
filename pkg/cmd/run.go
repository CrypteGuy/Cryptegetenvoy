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

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tetratelabs/getenvoy/pkg/binary/getenvoy"
)

// NewRunCmd create a command responsible for starting an Envoy process
func NewRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run [binary] -- <envoy-args>",
		Short: "Starts an Envoy process using the binary passed.",
		Long: `
Starts an Envoy process using the binary passed. 
Location can be a manifest reference or local file.`,
		Example: `# Run using a manifest reference. Reference format is <flavor>:<version>.
getenvoy run standard:1.10.1 -- --config-path ./bootstrap.yaml

# Run using a local file.
getenvoy run ./envoy -- --config-path ./bootstrap.yaml

# List available Envoy flags
getenvoy run standard:1.10.1 -- --help
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("missing binary parameter")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			runtime, err := getenvoy.New()
			if err != nil {
				return err
			}
			return runtime.Run(args[0], args[1:])
		},
	}
}
