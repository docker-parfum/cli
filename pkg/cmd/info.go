// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nitrictech/cli/pkg/ghissue"
	"github.com/nitrictech/cli/pkg/output"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gather information about Nitric and the environment",
	Long:  `Gather information about Nitric and the environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		d := ghissue.Gather()

		s, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			output.Print(d)
		} else {
			fmt.Println(string(s))
		}
	},
}
