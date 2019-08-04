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

package version

import (
	"bytes"
	"text/template"

	"github.com/spf13/cobra"
)

const versionTemplate = `GitVersion:		{{.GitVersion}}
BuildDate:		{{.BuildDate}}
GoVersion:		{{.GoVersion}}
Compiler:		{{.Compiler}}
Platform:		{{.Platform}}
`

// AddCommand sets version template to be used.
func AddCommand(cmd *cobra.Command) {
	cmd.Version = String()
	cmd.SetVersionTemplate("{{.Version}}")
}

// String gets the formatted version string
func String() string {
	tpl, err := template.New("").Parse(versionTemplate)
	if err == nil {
		v := Get()
		var b bytes.Buffer
		if err := tpl.Execute(&b, v); err == nil {
			return b.String()
		}
	}
	return Get().String()
}
