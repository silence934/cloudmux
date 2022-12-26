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
	"strconv"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/qcloud"
)

func init() {
	type DomianListOptions struct {
		Offset  int
		Limit   int
		Keyword string
	}
	shellutils.R(&DomianListOptions{}, "domain-list", "List domains", func(cli *qcloud.SRegion, args *DomianListOptions) error {
		domains, total, e := cli.GetClient().GetDomains(args.Keyword, args.Offset, args.Limit)
		if e != nil {
			return e
		}
		printList(domains, total, args.Offset, args.Limit, []string{})
		return nil
	})

	type DomianCreateOptions struct {
		DOMAIN string
	}
	shellutils.R(&DomianCreateOptions{}, "domain-create", "create domain", func(cli *qcloud.SRegion, args *DomianCreateOptions) error {
		domain, e := cli.GetClient().CreateDomian(args.DOMAIN)
		if e != nil {
			return e
		}
		printObject(domain)
		return nil
	})

	type DomianDeleteOptions struct {
		DOMAIN string
	}
	shellutils.R(&DomianDeleteOptions{}, "domain-delete", "delete domains", func(cli *qcloud.SRegion, args *DomianDeleteOptions) error {
		e := cli.GetClient().DeleteDomian(args.DOMAIN)
		if e != nil {
			return e
		}
		return nil
	})

	type DomianShowOptions struct {
		DOMAIN string
	}
	shellutils.R(&DomianShowOptions{}, "domain-show", "show domains", func(cli *qcloud.SRegion, args *DomianShowOptions) error {
		domain, err := cli.GetClient().GetDomain(args.DOMAIN)
		if err != nil {
			return err
		}
		printObject(domain)
		return nil
	})

	type DnsRecordListOptions struct {
		DOMAIN string
		Offset int
		Limit  int
	}
	shellutils.R(&DnsRecordListOptions{}, "dnsrecord-list", "List dndrecord", func(cli *qcloud.SRegion, args *DnsRecordListOptions) error {
		records, total, e := cli.GetClient().GetDnsRecords(args.DOMAIN, args.Offset, args.Limit)
		if e != nil {
			return e
		}
		printList(records, total, args.Offset, args.Limit, []string{})
		return nil
	})

	type DnsRecordCreateOptions struct {
		DOMAIN string
		NAME   string
		VALUE  string //joined by '*'
		TTL    int64
		TYPE   string
	}
	shellutils.R(&DnsRecordCreateOptions{}, "dnsrecord-create", "create dndrecord", func(cli *qcloud.SRegion, args *DnsRecordCreateOptions) error {
		change := cloudprovider.DnsRecordSet{}
		change.DnsName = args.NAME
		change.DnsValue = args.VALUE
		change.Ttl = args.TTL
		change.DnsType = cloudprovider.TDnsType(args.TYPE)
		_, e := cli.GetClient().CreateDnsRecord(&change, args.DOMAIN)
		if e != nil {
			return e
		}
		return nil
	})

	type DnsRecordUpdateOptions struct {
		DOMAIN   string
		RECORDID int
		NAME     string
		VALUE    string //joined by '*'
		TTL      int64
		TYPE     string
	}
	shellutils.R(&DnsRecordUpdateOptions{}, "dnsrecord-update", "update dndrecord", func(cli *qcloud.SRegion, args *DnsRecordUpdateOptions) error {
		change := cloudprovider.DnsRecordSet{}
		change.DnsName = args.NAME
		change.ExternalId = strconv.Itoa(args.RECORDID)
		change.DnsValue = args.VALUE
		change.Ttl = args.TTL
		change.DnsType = cloudprovider.TDnsType(args.TYPE)
		e := cli.GetClient().ModifyDnsRecord(&change, args.DOMAIN)
		if e != nil {
			return e
		}
		return nil
	})

	type DnsRecordUpdateStatusOptions struct {
		DOMAIN   string
		RECORDID int
		STATUS   string `choices:"DISABLE|ENABLE"`
	}
	shellutils.R(&DnsRecordUpdateStatusOptions{}, "dnsrecord-updatestatus", "update dndrecord", func(cli *qcloud.SRegion, args *DnsRecordUpdateStatusOptions) error {
		e := cli.GetClient().ModifyRecordStatus(args.STATUS, strconv.Itoa(args.RECORDID), args.DOMAIN)
		if e != nil {
			return e
		}
		return nil
	})

	type DnsRecordRemoveOptions struct {
		DOMAIN   string
		RECORDID string
	}
	shellutils.R(&DnsRecordRemoveOptions{}, "dnsrecord-delete", "delete dndrecord", func(cli *qcloud.SRegion, args *DnsRecordRemoveOptions) error {
		e := cli.GetClient().DeleteDnsRecord(args.RECORDID, args.DOMAIN)
		if e != nil {
			return e
		}
		return nil
	})
}
