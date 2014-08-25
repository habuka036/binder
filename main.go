// Copyright (c) 2014 Osamu Habuka. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

// This source code is fork from SkyDNS.

package main

import (
    "flag"
    "strings"
    "log"
    "os"

//    "github.com/coreos/go-etcd/etcd"
)

const Version = "0.0.0"

var (
   config = &Config{
      VolumePath: "",
      BindAddress: "",
   }
   etcdhosts = ""
)

func env(key, def string) string {
   if x := os.Getenv(key); x != "" {
      return x
   }
   return def
}

func init() {
   flag.StringVar(
      &config.VolumePath, "volumepath", 
      env("VOLUME_PATH", "/var/lib/binder/volumes"), 
      "volume file path",
   )
   flag.StringVar(
      &config.BindAddress, "bind",
      env("BIND_ADDRESS", "127.0.0.1:8776"),
      "bind address of binder",
   )
   flag.StringVar(
      &etcdhosts, "etcdhosts", 
      env("ETCD_HOSTS", "127.0.0.1:4001"), 
      "ip:port of etcd",
   )
}

func main() {
   flag.Parse()
   hosts := strings.Split(etcdhosts, ",")
   client := NewClient(hosts)
   config, err := loadConfig(client, config)
   if err != nil {
      log.Fatal(err)
   }
   s := NewServer(config, client)
   if err := s.Run(); err != nil {
      log.Fatal(err)
   }
}
