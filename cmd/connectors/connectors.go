// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectors

import (
	"github.com/spf13/cobra"
)

// Cmd to manage preferences
var Cmd = &cobra.Command{
	Use:   "connectors",
	Short: "Manage Connectors for Integration",
	Long:  "Manage Connectors for Integration",
}

var region, project, name string

func init() {
	Cmd.AddCommand(ListCmd)
	Cmd.AddCommand(IamCmd)
}
