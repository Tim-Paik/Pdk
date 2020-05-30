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
	"github.com/spf13/cobra"
	"pdk/pkg"
)

// md5sumCmd represents the md5sum command
var md5sumCmd = &cobra.Command{
	Use:   "md5sum <filePath>",
	Short: "Print MD5 (128-bit) checksums.",
	Long:  `Print MD5 (128-bit) checksums.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.HelpFunc()
			return
		} else {
			for i := 0; i < len(args); i++ {
				if md5, err := pkg.Md5Sum(args[i]); err != nil {
					fmt.Println(err)
					return
				} else {
					fmt.Println(md5 + " " + args[i])
				}
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(md5sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
