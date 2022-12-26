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
	"time"

	"yunion.io/x/log"
	"yunion.io/x/pkg/util/shellutils"
	"yunion.io/x/pkg/util/timeutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type NamespaceListOptions struct {
	}
	shellutils.R(&NamespaceListOptions{}, "namespace-list", "List monbitor metric namespaces", func(cli *aliyun.SRegion, args *NamespaceListOptions) error {
		nslist, err := cli.FetchNamespaces()
		if err != nil {
			return err
		}
		printList(nslist, 0, 0, 0, nil)
		return nil
	})

	type MetricListOptions struct {
		NAMESPACE string `help:"namespace"`
	}
	shellutils.R(&MetricListOptions{}, "metrics-list", "List metrics in a namespace", func(cli *aliyun.SRegion, args *MetricListOptions) error {
		metrics, err := cli.FetchMetrics(args.NAMESPACE)
		if err != nil {
			return err
		}
		printList(metrics, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&cloudprovider.MetricListOptions{}, "metric-list", "List metrics in a namespace", func(cli *aliyun.SRegion, args *cloudprovider.MetricListOptions) error {
		metrics, err := cli.GetClient().GetMetrics(args)
		if err != nil {
			return err
		}
		for i := range metrics {
			log.Infof("metric: %s %s %s", metrics[i].Id, metrics[i].MetricType, metrics[i].Unit)
			printList(metrics[i].Values, len(metrics[i].Values), 0, 0, []string{})
		}
		return nil
	})

	type DescribeMetricListOptions struct {
		METRIC    string `help:"metric name"`
		NAMESPACE string `help:"name space"`
		Since     string `help:"since, 2019-11-29T11:22:00Z"`
		Until     string `help:"since, 2019-11-30T11:22:00Z"`
	}
	shellutils.R(&DescribeMetricListOptions{}, "metric-data-list", "DescribeMetricList", func(cli *aliyun.SRegion, args *DescribeMetricListOptions) error {
		var since time.Time
		var err error
		if len(args.Since) > 0 {
			since, err = timeutils.ParseTimeStr(args.Since)
			if err != nil {
				return err
			}
		}
		var until time.Time
		if len(args.Until) > 0 {
			until, err = timeutils.ParseTimeStr(args.Until)
			if err != nil {
				return err
			}
		}
		data, err := cli.FetchMetricData(args.METRIC, args.NAMESPACE, since, until)
		if err != nil {
			return err
		}
		printList(data, 0, 0, 0, nil)
		return nil
	})
}
