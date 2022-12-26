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
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/huawei"
)

func init() {
	type RouteTableListOption struct {
		VpcId string `help:"Vpc ID"`
	}
	shellutils.R(&RouteTableListOption{}, "routetable-list", "List vpc route tables", func(cli *huawei.SRegion, args *RouteTableListOption) error {
		routetables, err := cli.GetRouteTables(args.VpcId)
		if err != nil {
			return err
		}
		printList(routetables, 0, 0, 0, nil)
		return nil
	})

	type RouteTableIdOptions struct {
		ID string `help:"RouteTable ID"`
	}
	shellutils.R(&RouteTableIdOptions{}, "routetable-show", "Show vpc route table", func(cli *huawei.SRegion, args *RouteTableIdOptions) error {
		routetable, err := cli.GetRouteTable(args.ID)
		if err != nil {
			return err
		}
		printObject(routetable)
		return nil
	})

}
