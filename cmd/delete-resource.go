// Package cmd implements fyde-cli commands
package cmd

/*
Copyright © 2019 Fyde, Inc. <hello@fyde.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"

	apiresources "github.com/fyde/fyde-cli/client/access_resources"
)

// resourceDeleteCmd represents the delete command
var resourceDeleteCmd = &cobra.Command{
	Use:     "delete [resource ID]...",
	Aliases: []string{"remove", "rm"},
	Short:   "Delete resources",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := preRunCheckAuth(cmd, args)
		if err != nil {
			return err
		}

		err = preRunFlagChecks(cmd, args)
		if err != nil {
			return err
		}

		if !multiOpCheckArgsPresent(cmd, args) {
			return fmt.Errorf("missing resource ID argument")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceIDs, err := multiOpParseUUIDArgs(cmd, args, "id")
		if err != nil {
			return err
		}

		delete := func(ids []strfmt.UUID) error {
			params := apiresources.NewDeleteResourceParams()
			params.SetID(ids)

			_, err := global.Client.AccessResources.DeleteResource(params, global.AuthWriter)
			if err != nil {
				return processErrorResponse(err)
			}
			return nil
		}

		tw, j := multiOpBuildTableWriter()

		if loopControlContinueOnError(cmd) {
			// then we must delete individually, because on a request for multiple deletions,
			// the server does nothing if one fails

			for _, id := range resourceIDs {
				err = delete([]strfmt.UUID{id})
				var result interface{}
				result = "success"
				if err != nil {
					result = err
				}
				multiOpTableWriterAppend(tw, &j, id, result)
			}
			err = nil
		} else {
			err = delete(resourceIDs)
			var result interface{}
			result = "success"
			if err != nil {
				result = err
			}
			multiOpTableWriterAppend(tw, &j, "*", result)
		}

		return printListOutputAndError(cmd, j, tw, len(resourceIDs), err)
	},
}

func init() {
	resourcesCmd.AddCommand(resourceDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourceDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourceDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	initMultiOpArgFlags(resourceDeleteCmd, "resource", "delete", "id", "[]strfmt.UUID")
	initOutputFlags(resourceDeleteCmd)
	initLoopControlFlags(resourceDeleteCmd)
}
