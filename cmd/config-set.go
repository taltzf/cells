/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/pydio/cells/common/config"
	"github.com/pydio/cells/common/utils"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "set",
	Short: "Store a configuration",
	Long: `Stores a configuration. Will be stored in both your pydio.json file and in the database.

SYNTAX
======
Configurations are represented by three parameters that you must pass as arguments :
- serviceName: name of the corresponding service
- configName: name of the parameter
- configValue: json-encoded value of the parameter you want to set/change

EXAMPLES
========
Change the port of micro.web service (rest api)
$ config set micro.web port 8083

Json parameter value
$ config set pydio.grpc.yourservice configName '{"key":"value"}'

`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("Requires at least three arguments, please see 'pydio config set --help'")
		}

		// IsValidService ?
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		path := args[1]
		data := args[2]

		config.Set(data, "services", id, path)

		utils.SaveConfigs()
		cmd.Println("Config set")
	},
}

func init() {
	configCmd.AddCommand(updateCmd)
}
