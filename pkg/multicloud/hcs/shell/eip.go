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

	"yunion.io/x/cloudmux/pkg/multicloud/hcs"
)

func init() {
	type EipListOptions struct {
		PortId  string
		Address []string
	}
	shellutils.R(&EipListOptions{}, "eip-list", "List eips", func(cli *hcs.SRegion, args *EipListOptions) error {
		eips, err := cli.GetEips(args.PortId, args.Address)
		if err != nil {
			return nil
		}
		printList(eips, 0, 0, 0, nil)
		return nil
	})

	type EipTypeListOptions struct {
	}

	shellutils.R(&EipTypeListOptions{}, "eip-type-list", "List eip types", func(cli *hcs.SRegion, args *EipTypeListOptions) error {
		types, err := cli.GetEipTypes()
		if err != nil {
			return nil
		}
		printList(types, 0, 0, 0, nil)
		return nil
	})

	type EipAllocateOptions struct {
		NAME       string `help:"eip name"`
		Ip         string
		BW         int    `help:"Bandwidth limit in Mbps"`
		BGP        string `help:"bgp type"`
		SubnetId   string
		ChargeType string `help:"eip charge type" default:"traffic" choices:"traffic|bandwidth"`
		ProjectId  string
	}
	shellutils.R(&EipAllocateOptions{}, "eip-create", "Allocate an EIP", func(cli *hcs.SRegion, args *EipAllocateOptions) error {
		eip, err := cli.AllocateEIP(args.NAME, args.Ip, args.BW, hcs.TInternetChargeType(args.ChargeType), args.BGP, args.SubnetId, args.ProjectId)
		if err != nil {
			return err
		}
		printObject(eip)
		return nil
	})

	type EipReleaseOptions struct {
		ID string `help:"EIP allocation ID"`
	}
	shellutils.R(&EipReleaseOptions{}, "eip-delete", "Release an EIP", func(cli *hcs.SRegion, args *EipReleaseOptions) error {
		err := cli.DeallocateEIP(args.ID)
		return err
	})

	type EipAssociateOptions struct {
		ID       string `help:"EIP allocation ID"`
		INSTANCE string `help:"Instance ID"`
	}
	shellutils.R(&EipAssociateOptions{}, "eip-associate", "Associate an EIP", func(cli *hcs.SRegion, args *EipAssociateOptions) error {
		return cli.AssociateEip(args.ID, args.INSTANCE)
	})

	type EipDissociateOptions struct {
		ID string `help:"EIP allocation ID"`
	}

	shellutils.R(&EipDissociateOptions{}, "eip-dissociate", "Dissociate an EIP", func(cli *hcs.SRegion, args *EipDissociateOptions) error {
		return cli.DissociateEip(args.ID)
	})

}
