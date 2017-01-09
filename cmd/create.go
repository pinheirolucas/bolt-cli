// Copyright © 2017 Lucas Pinheiro <lucas.contato1996@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const defaultDBName = "bolt.db"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create [path to create the database]",
	Aliases: []string{"new"},
	Short:   "Create a new empty BoltDB file",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			er(errors.New("create needs the filepath of the new BoltDB"))
		}

		p := args[0]

		absp, err := filepath.Abs(p)
		if err != nil {
			er(err)
		}

		e := filepath.Ext(absp)
		switch e {
		case "":
			isDir, err := dirExists(absp)
			if err != nil {
				er(err)
			} else if !isDir {
				er(errors.Errorf("%s does not exists", absp))
			}

			absp = filepath.Join(absp, defaultDBName)
		case ".db":
			dp := filepath.Dir(absp)

			isDir, err := dirExists(dp)
			if err != nil {
				er(err)
			} else if !isDir {
				er(errors.Errorf("%s does not exists", dp))
			}
		default:
			er(errors.New("the name of DB must have the extension \".db\""))
		}

		_, err = bolt.Open(absp, 0600, bolt.DefaultOptions)
		if err != nil {
			er(err)
		}

		fmt.Println(fmt.Sprintf("New BoltDB created at: %s", absp))
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
