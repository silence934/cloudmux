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
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/hcs"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type ElbListOptions struct {
	}
	shellutils.R(&ElbListOptions{}, "elb-list", "List loadbalancers", func(cli *hcs.SRegion, args *ElbListOptions) error {
		elbs, err := cli.GetILoadBalancers()
		if err != nil {
			return err
		}
		printList(elbs, len(elbs), 0, 0, []string{})
		return nil
	})

	shellutils.R(&cloudprovider.SLoadbalancerCreateOptions{}, "elb-create", "create loadbalancer", func(cli *hcs.SRegion, args *cloudprovider.SLoadbalancerCreateOptions) error {
		elb, err := cli.CreateLoadBalancer(args)
		if err != nil {
			return err
		}

		printObject(elb)
		return nil
	})

	type ElbIdOptions struct {
		ID string `help:"loadblancer id"`
	}
	shellutils.R(&ElbIdOptions{}, "elb-delete", "delete loadbalancer", func(cli *hcs.SRegion, args *ElbIdOptions) error {
		err := cli.DeleteLoadBalancer(args.ID)
		if err != nil {
			return err
		}

		return nil
	})

	shellutils.R(&ElbIdOptions{}, "elb-show", "show loadbalancer", func(cli *hcs.SRegion, args *ElbIdOptions) error {
		lb, err := cli.GetLoadbalancer(args.ID)
		if err != nil {
			return err
		}
		printObject(lb)
		return nil
	})

	type ElbListenerListOptions struct {
		ElbID string `help:"loadblancer id"`
	}
	shellutils.R(&ElbListenerListOptions{}, "elb-listener-list", "list loadbalancer listener", func(cli *hcs.SRegion, args *ElbListenerListOptions) error {
		listeners, err := cli.GetLoadBalancerListeners(args.ElbID)
		if err != nil {
			return err
		}

		printList(listeners, len(listeners), 0, 0, []string{})
		return nil
	})

	type ElbListenerCreateOptions struct {
		Name          string `help:"listener name"`
		Desc          string `help:"listener Description"`
		Http2         bool   `help:"http2 enable status"`
		PoolID        string `help:"default backend group id"`
		CertId        string `help:"default certification id"`
		XForwardedFor bool   `help:"XForwardedFor enable status"`
		LISTENER_TYPE string `help:"listener type" choices:"tcp|udp|http|https"`
		LISTENER_PORT int    `help:"listener port"`
		ELB_ID        string `help:"loadbalancer id"`
	}
	shellutils.R(&ElbListenerCreateOptions{}, "elb-listener-create", "create loadbalancer listener", func(cli *hcs.SRegion, args *ElbListenerCreateOptions) error {
		input := &cloudprovider.SLoadbalancerListenerCreateOptions{
			Name:           args.Name,
			ListenerType:   args.LISTENER_TYPE,
			ListenerPort:   args.LISTENER_PORT,
			BackendGroupId: args.PoolID,
			EnableHTTP2:    args.Http2,
			CertificateId:  args.CertId,
			Description:    args.Desc,
			XForwardedFor:  args.XForwardedFor,
		}
		listener, err := cli.CreateLoadBalancerListener(args.ELB_ID, input)
		if err != nil {
			return err
		}

		printObject(listener)
		return nil
	})

	type ElbListenerUpdateOptions struct {
		Name          string `help:"listener name"`
		Desc          string `help:"listener Description"`
		Http2         bool   `help:"http2 enable status"`
		PoolID        string `help:"default backend group id"`
		CertId        string `help:"default certification id"`
		XForwardedFor bool   `help:"XForwardedFor enable status"`
		LISTENER_ID   string `help:"listener id"`
	}
	shellutils.R(&ElbListenerUpdateOptions{}, "elb-listener-update", "update loadbalancer listener", func(cli *hcs.SRegion, args *ElbListenerUpdateOptions) error {
		input := &cloudprovider.SLoadbalancerListenerCreateOptions{
			Name:           args.Name,
			BackendGroupId: args.PoolID,
			EnableHTTP2:    args.Http2,
			CertificateId:  args.CertId,
			Description:    args.Desc,
			XForwardedFor:  args.XForwardedFor,
		}

		err := cli.UpdateLoadBalancerListener(args.LISTENER_ID, input)
		if err != nil {
			return err
		}

		return nil
	})

	type ElbBackendGroupListOptions struct {
		ElbID string `help:"loadbalancer id"`
	}
	shellutils.R(&ElbBackendGroupListOptions{}, "elb-backend-group-list", "List backend groups", func(cli *hcs.SRegion, args *ElbBackendGroupListOptions) error {
		elbbg, err := cli.GetLoadBalancerBackendGroups(args.ElbID)
		if err != nil {
			return err
		}

		printList(elbbg, len(elbbg), 0, 0, []string{})
		return nil
	})

	type ElbBackendGroupCreateOptions struct {
		Name      string `help:"backend group name"`
		Desc      string `help:"backend group description"`
		PROTOCOL  string `help:"backend group protocol" choices:"tcp|udp|http"`
		ALGORITHM string `help:"backend group algorithm" choices:"rr|wlc|sch"`
		ElbID     string `help:"loadbalancer id belong to"`
	}
	shellutils.R(&ElbBackendGroupCreateOptions{}, "elb-backend-group-create", "Create backend groups", func(cli *hcs.SRegion, args *ElbBackendGroupCreateOptions) error {
		opts := &cloudprovider.SLoadbalancerBackendGroup{
			Name:      args.Name,
			Protocol:  args.PROTOCOL,
			Scheduler: args.ALGORITHM,
		}
		elbbg, err := cli.CreateLoadBalancerBackendGroup(args.ElbID, opts)
		if err != nil {
			return err
		}
		printObject(elbbg)
		return nil
	})

	type ElbBackendGroupDeleteOptions struct {
		POOL_ID string `help:"backend group id"`
	}
	shellutils.R(&ElbBackendGroupDeleteOptions{}, "elb-backend-group-delete", "Delete backend group", func(cli *hcs.SRegion, args *ElbBackendGroupDeleteOptions) error {
		err := cli.DeleteLoadBalancerBackendGroup(args.POOL_ID)
		if err != nil {
			return err
		}

		return nil
	})

	type ElbBackendAddOptions struct {
		Name      string `help:"backend name"`
		POOL_ID   string `help:"backend group id"`
		SUBNET_ID string `help:"instance subnet id"`
		ADDRESS   string `help:"instance ip address"`
		PORT      int    `help:"backend protocol port  [1，65535]"`
		Weight    int    `help:"backend weight [0，100]" default:"1"`
	}
	shellutils.R(&ElbBackendAddOptions{}, "elb-backend-add", "Add backend to backendgroup", func(cli *hcs.SRegion, args *ElbBackendAddOptions) error {
		elbb, err := cli.AddLoadBalancerBackend(args.POOL_ID, args.SUBNET_ID, args.ADDRESS, args.PORT, args.Weight)
		if err != nil {
			return err
		}

		printObject(elbb)
		return nil
	})

	type ElbBackendListOptions struct {
		POOL_ID string `help:"backend group id"`
	}
	shellutils.R(&ElbBackendListOptions{}, "elb-backend-list", "list backend", func(cli *hcs.SRegion, args *ElbBackendListOptions) error {
		elbb, err := cli.GetLoadBalancerBackends(args.POOL_ID)
		if err != nil {
			return err
		}

		printList(elbb, len(elbb), 0, 0, []string{})
		return nil
	})

	type ElbListenerPolicyListOptions struct {
		ListenerID string `help:"listener id"`
	}
	shellutils.R(&ElbListenerPolicyListOptions{}, "elb-listener-policy-list", "List listener policies", func(cli *hcs.SRegion, args *ElbListenerPolicyListOptions) error {
		elblp, err := cli.GetLoadBalancerPolicies(args.ListenerID)
		if err != nil {
			return err
		}

		printList(elblp, len(elblp), 0, 0, []string{})
		return nil
	})

	type ElbListenerPolicyCreateOptions struct {
		LISTENER_ID string `help:"listener id"`
		Name        string `help:"policy name"`
		Domain      string `help:"policy domain"`
		Path        string `help:"policy path"`
		PoolID      string `help:"backend group name"`
	}
	shellutils.R(&ElbListenerPolicyCreateOptions{}, "elb-listener-policy-create", "Create listener policy", func(cli *hcs.SRegion, args *ElbListenerPolicyCreateOptions) error {
		rule := &cloudprovider.SLoadbalancerListenerRule{
			Name:           args.Name,
			Domain:         args.Domain,
			Path:           args.Path,
			BackendGroupId: args.PoolID,
		}

		elblp, err := cli.CreateLoadBalancerPolicy(args.LISTENER_ID, rule)
		if err != nil {
			return err
		}

		printObject(elblp)
		return nil
	})

	type ElbListenerPolicyDeleteOptions struct {
		POLICY_ID string `help:"policy id"`
	}
	shellutils.R(&ElbListenerPolicyDeleteOptions{}, "elb-listener-policy-delete", "Delete listener policy", func(cli *hcs.SRegion, args *ElbListenerPolicyDeleteOptions) error {
		err := cli.DeleteLoadBalancerPolicy(args.POLICY_ID)
		if err != nil {
			return err
		}

		return nil
	})

	type ElbListenerPolicyRuleListOptions struct {
		POLICY_ID string `help:"policy id"`
	}
	shellutils.R(&ElbListenerPolicyRuleListOptions{}, "elb-listener-policyrule-list", "List listener policy rules", func(cli *hcs.SRegion, args *ElbListenerPolicyRuleListOptions) error {
		elblpr, err := cli.GetLoadBalancerPolicyRules(args.POLICY_ID)
		if err != nil {
			return err
		}

		printList(elblpr, len(elblpr), 0, 0, []string{})
		return nil
	})
}
