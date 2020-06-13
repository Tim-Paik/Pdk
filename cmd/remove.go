/*
Copyright © 2020 Tim_Paik <timpaik@163.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"pdk/pkg"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove <package(s)>",
	Short:   "Remove package(s)",
	Long:    `Remove package(s) from local data`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		pkg.CheckRoot()
		if err := pkg.Remove(args); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
