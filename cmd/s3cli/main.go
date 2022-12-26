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

package main

import (
	"fmt"
	"os"

	"yunion.io/x/pkg/util/shellutils"
	"yunion.io/x/structarg"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/objectstore"
	"yunion.io/x/cloudmux/pkg/multicloud/objectstore/ceph"
	_ "yunion.io/x/cloudmux/pkg/multicloud/objectstore/shell"
	"yunion.io/x/cloudmux/pkg/multicloud/objectstore/xsky"
)

type BaseOptions struct {
	Debug      bool   `help:"debug mode"`
	AccessUrl  string `help:"Access url" default:"$S3_ACCESS_URL" metavar:"S3_ACCESS_URL"`
	AccessKey  string `help:"Access key" default:"$S3_ACCESS_KEY" metavar:"S3_ACCESS_KEY"`
	Secret     string `help:"Secret" default:"$S3_SECRET" metavar:"S3_SECRET"`
	Backend    string `help:"Backend driver" default:"$S3_BACKEND" metavar:"S3_BACKEND"`
	SUBCOMMAND string `help:"s3cli subcommand" subcommand:"true"`
}

func getSubcommandParser() (*structarg.ArgumentParser, error) {
	parse, e := structarg.NewArgumentParserWithHelp(&BaseOptions{},
		"s3cli",
		"Command-line interface to standard S3 API.",
		`See "s3cli COMMAND --help" for help on a specific command.`)

	if e != nil {
		return nil, e
	}

	subcmd := parse.GetSubcommand()
	if subcmd == nil {
		return nil, fmt.Errorf("No subcommand argument.")
	}
	for _, v := range shellutils.CommandTable {
		_, e := subcmd.AddSubParserWithHelp(v.Options, v.Command, v.Desc, v.Callback)
		if e != nil {
			return nil, e
		}
	}
	return parse, nil
}

func showErrorAndExit(e error) {
	fmt.Fprintf(os.Stderr, "%s", e)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func newClient(options *BaseOptions) (cloudprovider.ICloudRegion, error) {
	if len(options.AccessUrl) == 0 {
		return nil, fmt.Errorf("Missing accessUrl")
	}

	if len(options.AccessKey) == 0 {
		return nil, fmt.Errorf("Missing accessKey")
	}

	if len(options.Secret) == 0 {
		return nil, fmt.Errorf("Missing secret")
	}

	if options.Backend == api.CLOUD_PROVIDER_CEPH {
		return ceph.NewCephRados(
			objectstore.NewObjectStoreClientConfig(
				options.AccessUrl, options.AccessKey, options.Secret,
			).Debug(options.Debug),
		)
	} else if options.Backend == api.CLOUD_PROVIDER_XSKY {
		return xsky.NewXskyClient(
			objectstore.NewObjectStoreClientConfig(
				options.AccessUrl, options.AccessKey, options.Secret,
			).Debug(options.Debug),
		)
	}
	return objectstore.NewObjectStoreClient(
		objectstore.NewObjectStoreClientConfig(
			options.AccessUrl, options.AccessKey, options.Secret,
		).Debug(options.Debug),
	)
}

func main() {
	parser, e := getSubcommandParser()
	if e != nil {
		showErrorAndExit(e)
	}
	e = parser.ParseArgs(os.Args[1:], false)
	options := parser.Options().(*BaseOptions)

	if parser.IsHelpSet() {
		fmt.Print(parser.HelpString())
		return
	}
	subcmd := parser.GetSubcommand()
	subparser := subcmd.GetSubParser()
	if e != nil || subparser == nil {
		if subparser != nil {
			fmt.Print(subparser.Usage())
		} else {
			fmt.Print(parser.Usage())
		}
		showErrorAndExit(e)
	}
	suboptions := subparser.Options()
	if subparser.IsHelpSet() {
		fmt.Print(subparser.HelpString())
		return
	}
	var client cloudprovider.ICloudRegion
	client, e = newClient(options)
	if e != nil {
		showErrorAndExit(e)
	}
	e = subcmd.Invoke(client, suboptions)
	if e != nil {
		showErrorAndExit(e)
	}
}
