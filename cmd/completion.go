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
	"github.com/spf13/cobra"
	"os"
)

var long = fmt.Sprintf(`Output completion script for Bash and Zsh.

Sourcing the completion script in your shell enables completion for %s.
How to do this depends on your shell and OS:

Bash on Linux
=============
  The completion script depends on bash-completion. Install it with:

    $ apt-get install bash-completion

  Then make sure the completion script gets sourced in all your shell sessions:

    $ echo 'source <(%s completion bash)' >>~/.bashrc
    # or
    $ %s completion bash >/etc/bash_completion.d/%s

Bash on macOS
=============
  The completion script depends on bash-completion v2 which depens on Bash 4.1+.
  By default macOS contains Bash 3.2. To use this completion script, you need to
  install Bash 4.1+ (see https://itnext.io/upgrading-bash-on-macos-7138bd1066ba).

  After upgrading to Bash 4.1+, install bash-completion v2:

    $ brew install bash-completion@2

  Now make sure the completion script gets sourced in all your shell sessions:

    $ echo 'source <(%s completion bash)' >>~/.bashrc
    # or
    $ %s completion bash >/usr/local/etc/bash_completion.d/%s

Zsh
===
  The completion script doesn't have any dependencies for Zsh! Simply make sure
  it gets sourced in your .zshrc file:

    $ echo 'source <(%s completion zsh)' >>~/.zshrc`,
	rootCmd.Use, rootCmd.Use, rootCmd.Use, rootCmd.Use, rootCmd.Use, rootCmd.Use,
	rootCmd.Use, rootCmd.Use)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:                   "completion bash|zsh",
	Short:                 "Output completion script for Bash and Zsh",
	Long:                  long,
	Args:                  cobra.ExactValidArgs(1),
	ValidArgs:             []string{"bash", "zsh"},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "bash" {
			rootCmd.GenBashCompletion(os.Stdout)
		} else {
			rootCmd.GenZshCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
