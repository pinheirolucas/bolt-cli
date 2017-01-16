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

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/pinheirolucas/bolt-cli/cache"
	"github.com/pinheirolucas/bolt-cli/utils"
	"github.com/pinheirolucas/bolt-cli/validator"
)

var setCmdKey string
var setCmdValue string
var setCmdBucket string
var setCmdJSON string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Creates bucket with a given name and value.",
	Long:  "",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := validator.IsBoltDBValid(args); err != nil {
			er(err)
		}

		if setCmdBucket == "" {
			er(errors.New("a bucket name must be provided"))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		bp := args[0]

		err := cache.CreateBoltDB(bp)
		if err != nil {
			er(err)
		}
		defer cache.CloseBoltDB()

		if setCmdKey != "" && setCmdValue != "" {
			err := cache.InsertSimpleValue(setCmdBucket, setCmdKey, setCmdValue)
			if err != nil {
				er(err)
			}
		} else if setCmdKey == "" && setCmdValue != "" {
			err := cache.InsertComplexValue(setCmdBucket, setCmdValue)
			if err != nil {
				er(err)
			}
		} else if setCmdJSON != "" {
			err = validator.JSONPathValid(setCmdJSON)
			if err != nil {
				er(err)
			}

			data, err := utils.RetrieveJSONData(setCmdJSON)
			if err != nil {
				er(err)
			}

			err = cache.InsertBucketValue(setCmdBucket, data)
			if err != nil {
				er(err)
			}
		} else {
			er(errors.New("invalid combination of arguments"))
		}

		fmt.Println("The provided values were successfully inserted.")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(
		&setCmdBucket,
		"bucket", "b", "",
		"Creates a bucket with the provided name.",
	)
	setCmd.Flags().StringVarP(&setCmdKey, "key", "k", "", "The key to be created inside of a bucket.")
	setCmd.Flags().StringVarP(&setCmdValue, "value", "v", "", "Creates a simple register with the provided value.")
	setCmd.Flags().StringVarP(&setCmdJSON, "json", "j", "", "Path to a JSON file to insert into a bucket.")
}
