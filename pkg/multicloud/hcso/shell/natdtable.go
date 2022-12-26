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

	huawei "yunion.io/x/cloudmux/pkg/multicloud/hcso"
)

func init() {
	type SNatDTableOptions struct {
		NatGatewayID string `help:"Nat Gateway ID" positional:"true"`
	}
	shellutils.R(&SNatDTableOptions{}, "dnat-table-list", "List dnat table", func(region *huawei.SRegion, args *SNatDTableOptions) error {
		DNatTable, err := region.GetNatDTable(args.NatGatewayID)
		if err != nil {
			return err
		}
		printList(DNatTable, 0, 0, 0, nil)
		return nil
	})
}
