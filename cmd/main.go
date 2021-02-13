//   Copyright 2021 Syed Parvez
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/cmd"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
)

func init()  {
	// initialize config and other necessary packages
	_, err := config.GetConfig()
	if err != nil {
		handler.FatalError(err)
	}
}

func main() {
	cmd.Execute()
}
