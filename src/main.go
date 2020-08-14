package main

import (
  "fmt"
)

type Plugin interface {
  Run() bool
}

type Http struct {
  domain string
  method string
  port int
}

func (h Http) Run() bool {
  fmt.Println("running http plugin.")
  return true
}

func runPlugins(p [1]Plugin) {
  for i := 0; i < len(p); i++ {
    p[i].Run()
  }
}

func main() {
  plugins := [1]Plugin{Http{"archlinux.org", "GET", 80}}
  runPlugins(plugins)
}

