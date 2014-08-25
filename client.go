// Copyright (c) 2014 Osamu Habuka. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

// This source code is fork from SkyDNS.

package main

import (
//   "net/url"
//   "strings"

   "github.com/coreos/go-etcd/etcd"
)

func NewClient(hosts []string) (client *etcd.Client) {
   if 1 == len(hosts) && "" == hosts[0] {
      hosts[0] = "http://127.0.0.1:4001"
   }
   client = etcd.NewClient(hosts)
   client.SyncCluster()
   return client
}

func (s *server) UpdateClient(resp *etcd.Response) {
   
}

