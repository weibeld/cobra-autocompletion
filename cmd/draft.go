// Copyright © 2019 Daniel Weibel <daniel@weibeld.net>
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

	"github.com/weibeld/cobra"
)

// draftCmd represents the createDraft command
var draftCmd = &cobra.Command{
	Use:     "draft",
	Aliases: []string{"drft", "dr"},
	Short:   "Create a draft article in your Medium account.",
	Long:    `Create a draft article in your Medium account.`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("integration-token")
		title, _ := cmd.Flags().GetString("title")
		url, _ := cmd.Flags().GetString("canoncial-url")
		fmt.Println(token, title, url)
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)
	draftCmd.Flags().StringP("title", "t", "", "title of the article (required)")
	draftCmd.MarkFlagRequired("title")
	draftCmd.Flags().StringP("integration-token", "i", "", "your Medium integration token (required)")
	draftCmd.MarkFlagRequired("integration-token")
	draftCmd.Flags().StringP("canonical-url", "u", "", "canonical URL to be set for the article")
	draftCmd.Flags().BoolP("boolP", "b", false, "Description")
	draftCmd.Flags().Bool("bool", false, "Description")
	draftCmd.Flags().String("string", "foo", "your Medium integration token (required)")
	draftCmd.PersistentFlags().Bool("persistentBool", false, "Description")
	draftCmd.PersistentFlags().Bool("persistentBoolTrue", true, "Description")
}
