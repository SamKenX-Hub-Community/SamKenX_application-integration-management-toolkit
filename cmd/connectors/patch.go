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
	"os"

	"internal/apiclient"

	"github.com/GoogleCloudPlatform/application-integration-management-toolkit/client/connections"

	"github.com/spf13/cobra"
)

// PatchCmd to create a new connection
var PatchCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing connection",
	Long:  "Update an existing connection in a region",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if _, err := os.Stat(connectionFile); os.IsNotExist(err) {
			return err
		}

		content, err := os.ReadFile(connectionFile)
		if err != nil {
			return err
		}

		_, err = connections.Patch(name, content, updateMask)
		return
	},
}

var updateMask []string

func init() {
	PatchCmd.Flags().StringVarP(&name, "name", "n",
		"", "Connection name")
	PatchCmd.Flags().StringVarP(&connectionFile, "file", "f",
		"", "Connection details JSON file path")
	PatchCmd.Flags().StringArrayVarP(&updateMask, "update-mask", "",
		nil, "Update mask: A list of comma separates values to update")

	_ = PatchCmd.MarkFlagRequired("updateMask")
}
