// Copyright (c) 2014 Osamu Habuka. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

package main

import (
   "fmt"
   "strings"

   "github.com/coreos/go-etcd/etcd"
   "github.com/go-martini/martini"
   "github.com/martini-contrib/render"
)

type server struct {
   client *etcd.Client
   config *Config
}

type Volume struct {
   Key string
   Value string
}

func NewServer(config *Config, client *etcd.Client) *server {
   return &server{
      client: client,
      config: config,
   }
}

func (s *server) Run() error {
   m := martini.Classic()
   m.Use(render.Renderer())
   m.Group("/volume", func(r martini.Router) {
      r.Get("/(:id)*", s.list_volume)
      r.Post("/", s.create_volume)
      r.Delete("/:id", s.delete_volume)
   })
   m.Run()
   return nil
}

func (s *server) list_volume(params martini.Params, render render.Render) {
   println("list_volume() invoked")   
   path := "/binder/volumes/"
   if "" != params["id"] {
      println("list_volume(): id =", params["id"])
      path += params["id"]
   }
   r, err := s.client.Get(path, false, true)
   if err != nil {
      println("list_volume():", err)
   }
   var volumes []Volume
   if r.Node.Dir {
      for _, node := range r.Node.Nodes {
         id := strings.Split(node.Key, "/")[3]
         if !node.Dir {
            volume := Volume{id, node.Value}
            fmt.Println("volume", volume)
            volumes = append(volumes, volume)
         }
      }
   } else {
      id := strings.Split(r.Node.Key, "/")[3]
      if !r.Node.Dir {
         volume := Volume{id, r.Node.Value}
         fmt.Println("volume", volume)
         volumes = append(volumes, volume)
      }
   }
   render.JSON(200, volumes)
}

func (s *server) create_volume(params martini.Params) {
   println("create_volume() invoked")
   
}

func (s *server) delete_volume(params martini.Params) {
   println("delete_volume() invoked")
}

