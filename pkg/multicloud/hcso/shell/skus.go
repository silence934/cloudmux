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
	type SkuListOptions struct {
	}
	shellutils.R(&SkuListOptions{}, "sku-list", "List skus", func(cli *huawei.SRegion, args *SkuListOptions) error {
		skus, e := cli.GetISkus()
		if e != nil {
			return e
		}

		printList(skus, 0, 0, 0, []string{})
		return nil
	})
}
