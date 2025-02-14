// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/qcloud"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type NatDTableListOptions struct {
		Offset int
		Limit  int    `default:"10"`
		NATID  string `help:"Nat ID"`
	}
	shellutils.R(&NatDTableListOptions{}, "dtable-list", "List nat dtables", func(cli *qcloud.SRegion, args *NatDTableListOptions) error {
		tables, _, err := cli.GetDTables(args.NATID, args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(tables, len(tables), 0, 0, []string{})
		return nil
	})
}
