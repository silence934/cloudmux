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

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type FirewallListOptions struct {
		Network    string
		MaxResults int
		PageToken  string
	}
	shellutils.R(&FirewallListOptions{}, "firewall-list", "List firewalls", func(cli *google.SRegion, args *FirewallListOptions) error {
		firewalls, err := cli.GetClient().GetFirewalls(args.Network, args.MaxResults, args.PageToken)
		if err != nil {
			return err
		}
		printList(firewalls, 0, 0, 0, nil)
		return nil
	})

	type FirewallShowOptions struct {
		ID string
	}
	shellutils.R(&FirewallShowOptions{}, "firewall-show", "Show firewall", func(cli *google.SRegion, args *FirewallShowOptions) error {
		firewall, err := cli.GetClient().GetFirewall(args.ID)
		if err != nil {
			return err
		}
		printObject(firewall)
		return nil
	})

}
