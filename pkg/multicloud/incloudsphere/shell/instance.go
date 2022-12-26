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

	"yunion.io/x/cloudmux/pkg/multicloud/incloudsphere"
)

func init() {
	type InstanceListOptions struct {
		HOST_ID string
	}
	shellutils.R(&InstanceListOptions{}, "instance-list", "list instances", func(cli *incloudsphere.SRegion, args *InstanceListOptions) error {
		instances, err := cli.GetInstances(args.HOST_ID)
		if err != nil {
			return err
		}
		printList(instances, 0, 0, 0, []string{})
		return nil
	})

	type InstanceIdOptions struct {
		ID string
	}

	shellutils.R(&InstanceIdOptions{}, "instance-show", "show instance", func(cli *incloudsphere.SRegion, args *InstanceIdOptions) error {
		ret, err := cli.GetInstance(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&InstanceIdOptions{}, "instance-stop", "stop instance", func(cli *incloudsphere.SRegion, args *InstanceIdOptions) error {
		return cli.StopVm(args.ID)
	})

	type InstanceStartOptions struct {
		ID        string
		Password  string
		PublicKey string
	}

	shellutils.R(&InstanceStartOptions{}, "instance-start", "Start instance", func(cli *incloudsphere.SRegion, args *InstanceStartOptions) error {
		return cli.StartVm(args.ID, args.Password, args.PublicKey)
	})

	type InstanceAttachDiskOptions struct {
		ID      string
		DISK_ID string
	}

	shellutils.R(&InstanceAttachDiskOptions{}, "instance-attach-disk", "Attach instance disk", func(cli *incloudsphere.SRegion, args *InstanceAttachDiskOptions) error {
		return cli.AttachDisk(args.ID, args.DISK_ID)
	})

	shellutils.R(&InstanceAttachDiskOptions{}, "instance-detach-disk", "Attach instance disk", func(cli *incloudsphere.SRegion, args *InstanceAttachDiskOptions) error {
		return cli.DetachDisk(args.ID, args.DISK_ID)
	})

	type InstanceChangeConfigOptions struct {
		ID    string
		Cpu   int
		MemMb int
	}

	shellutils.R(&InstanceChangeConfigOptions{}, "instance-change-config", "Change instance config", func(cli *incloudsphere.SRegion, args *InstanceChangeConfigOptions) error {
		return cli.ChangeConfig(args.ID, args.Cpu, args.MemMb)
	})

}
