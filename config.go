// Copyright (c) 2014 Osamu Habuka. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

// This source code is fork from SkyDNS.

package main

import (
    "encoding/json"

    "github.com/coreos/go-etcd/etcd"
)

type Config struct {
   VolumePath string `json:"volume_path,omitempty"`
   BindAddress string `json:"bind_address,omitempty"`
}

func loadConfig(client *etcd.Client, config *Config) (*Config, error) {
   n, err := client.Get("/bind/config", false, false)
   if err != nil {
      if err := setDefaults(config); err != nil {
         return nil, err
      }
      return config, nil
   }
   if err := json.Unmarshal([]byte(n.Node.Value), &config); err != nil {
      return nil, err
   }
   if err := setDefaults(config); err != nil {
      return nil, err
   }
   return config, nil
}

func setDefaults(config *Config) error {
   if config.VolumePath == "" {
      config.VolumePath = "/var/lib/binder/volumes/"
   }
   if config.BindAddress == "" {
      config.BindAddress = "127.0.0.1:8776"
   }
   return nil
}

