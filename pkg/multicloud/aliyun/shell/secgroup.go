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
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type SecurityGroupListOptions struct {
		VpcId            string   `help:"VPC ID"`
		Name             string   `help:"Secgroup Name"`
		SecurityGroupIds []string `help:"SecurityGroup ids"`
		Limit            int      `help:"page size"`
		Offset           int      `help:"page offset"`
	}
	shellutils.R(&SecurityGroupListOptions{}, "security-group-list", "List security group", func(cli *aliyun.SRegion, args *SecurityGroupListOptions) error {
		secgrps, total, e := cli.GetSecurityGroups(args.VpcId, args.Name, args.SecurityGroupIds, args.Offset, args.Limit)
		if e != nil {
			return e
		}
		printList(secgrps, total, args.Offset, args.Limit, []string{})
		return nil
	})

	type SecurityGroupIdOptions struct {
		ID string `help:"ID or name of security group"`
	}
	shellutils.R(&SecurityGroupIdOptions{}, "security-group-show", "Show details of a security group", func(cli *aliyun.SRegion, args *SecurityGroupIdOptions) error {
		secgrp, err := cli.GetSecurityGroupDetails(args.ID)
		if err != nil {
			return err
		}
		printObject(secgrp)
		return nil
	})

	shellutils.R(&SecurityGroupIdOptions{}, "security-group-references", "Show references of a security group", func(cli *aliyun.SRegion, args *SecurityGroupIdOptions) error {
		references, err := cli.DescribeSecurityGroupReferences(args.ID)
		if err != nil {
			return err
		}
		printList(references, 0, 0, 0, nil)
		return nil
	})

	type SecurityGroupCreateOptions struct {
		NAME  string `help:"SecurityGroup name"`
		VpcId string `help:"VPC ID"`
		Desc  string `help:"SecurityGroup description"`
	}

	shellutils.R(&SecurityGroupCreateOptions{}, "security-group-create", "Create details of a security group", func(cli *aliyun.SRegion, args *SecurityGroupCreateOptions) error {
		secgroupId, err := cli.CreateSecurityGroup(args.VpcId, args.NAME, args.Desc)
		if err != nil {
			return err
		}
		fmt.Printf("secgroupId: %s", secgroupId)
		return nil
	})

}
