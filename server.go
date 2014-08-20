// Copyright (c) 2014 Osamu Habuka. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

package main

import (
    "github.com/go-martini/martini"
)

func main() {
   m := martini.Classic()
   m.Group("/volume", func(r martini.Router) {
      r.Get("/(:id)*", list_volume)
      r.Post("/", create_volume)
      r.Delete("/:id", delete_volume)
   })
   m.Run()
}

func list_volume(params martini.Params) {
   println("list_volume() invoked")   
   if "" != params["id"] {
      println("list_volume(): id =", params["id"])
   }
}

func create_volume(params martini.Params) {
   println("create_volume() invoked")
   
}

func delete_volume(params martini.Params) {
   println("delete_volume() invoked")
}

