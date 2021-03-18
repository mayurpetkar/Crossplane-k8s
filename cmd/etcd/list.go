/*
Copyright 2021 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package etcd

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/k0sproject/k0s/pkg/etcd"
)

func etcdListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member-list",
		Short: "Returns etcd cluster members list",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := getCmdOpts()
			ctx := context.Background()
			etcdClient, err := etcd.NewClient(c.K0sVars.CertRootDir, k0sVars.EtcdCertDir)
			if err != nil {
				return fmt.Errorf("can't list etcd cluster members: %v", err)
			}
			members, err := etcdClient.ListMembers(ctx)
			if err != nil {
				return fmt.Errorf("can't list etcd cluster members: %v", err)
			}
			l := logrus.New()
			l.SetFormatter(&logrus.JSONFormatter{})

			l.WithField("members", members).
				Info("done")
			return nil
		},
	}
	cmd.Flags().AddFlagSet(getPersistentFlagSet())
	return cmd
}
