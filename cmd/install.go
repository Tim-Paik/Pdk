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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install <package(s)>",
	Short:   "Install packages",
	Long:    `Download and install the latest package from your local repository (default)`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		var repoName string
		if repoName, err = cmd.Flags().GetString("repoName"); err != nil {
			fmt.Println(err)
			return
		}
		pkg.Install(args, repoName)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")
	installCmd.PersistentFlags().String("repoName", pkg.DefaultRepo, "Install from specified repo (repo by default)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
