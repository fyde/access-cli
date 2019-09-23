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

	"github.com/spf13/cobra"
)

// endpointSetCmd represents the endpoint set command
var endpointSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set console endpoint to use",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing endpoint argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO perform some sort of validation of the endpoint

		authViper.Set(ckeyAuthAccessToken, "")
		authViper.Set(ckeyAuthClient, "")
		authViper.Set(ckeyAuthUID, "")
		authViper.Set(ckeyAuthMethod, "")
		authViper.Set(ckeyAuthEndpoint, args[0])
		if global.WriteFiles {
			err := authViper.WriteConfig()
			if err != nil {
				return err
			}
		}
		cmd.Printf("Endpoint changed. Credentials cleared, please login again using `%s login`\n", ApplicationName)
		return nil
	},
}

func init() {
	endpointCmd.AddCommand(endpointSetCmd)
}
