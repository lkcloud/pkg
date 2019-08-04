/*
 * Copyright © 2019 Lingfei Kong <466701708@qq.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const DefaultConfigurationPath = "./conf"

// RunFunc defines the alias of the command entry function
type RunFunc func(cmd *cobra.Command, args []string)
type RunFuncE func(cmd *cobra.Command, args []string) error

// FormatBaseName is formatted as an executable file name under different
// operating systems according to the given name
func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}
	return basename
}

// AddConfigFlag defines the configuration path
func AddConfFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("config", "c", DefaultConfigurationPath, "config file path.")
}
