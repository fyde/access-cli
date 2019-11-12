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
	"strconv"

	"github.com/spf13/cobra"

	apiassets "github.com/fyde/fyde-cli/client/assets"
)

// domainDeleteCmd represents the delete command
var domainDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "rm"},
	Short:   "Delete domains",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := preRunCheckAuth(cmd, args)
		if err != nil {
			return err
		}

		err = preRunFlagChecks(cmd, args)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return fmt.Errorf("missing domain ID argument")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		assetIDs := make([]int64, len(args))
		var err error
		for i, arg := range args {
			assetIDs[i], err = strconv.ParseInt(arg, 10, 64)
			if err != nil {
				return err
			}
		}

		delete := func(ids []int64) error {
			params := apiassets.NewDeleteAssetParams()
			params.SetID(ids)

			_, err = global.Client.Assets.DeleteAsset(params, global.AuthWriter)
			if err != nil {
				return processErrorResponse(err)
			}
			return nil
		}

		if loopControlContinueOnError(cmd) {
			// then we must delete individually, because on a request for multiple deletions,
			// the server does nothing if one fails
			i := 0
			for _, assetID := range assetIDs {
				err = delete([]int64{assetID})
				if err != nil {
					cmd.PrintErrln(err)
				} else {
					// only keep successful deletions in list of assetIDs
					// this rewrites the array in place and lets us "delete" as we iterate
					// (junk is removed after the loop)
					assetIDs[i] = assetID
					i++
				}
			}
			// remove junk left at end of slice
			assetIDs = assetIDs[:i]
		} else {
			err = delete(assetIDs)
			if err != nil {
				return err
			}
		}

		printMultiOpOutput(cmd, "Domain", assetIDs, "deleted")
		return nil
	},
}

func init() {
	domainsCmd.AddCommand(domainDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	initLoopControlFlags(domainDeleteCmd)
}