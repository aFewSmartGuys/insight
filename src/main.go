package main

import (
  "fmt"
  "net/http"
)

type Plugin interface {
  Run() bool
  Name() string
}

type Http struct {
  domain string
  name string
}

func (h Http) Run() bool {
  fmt.Println("Running http plugin.")
  resp, err := http.Get(h.domain)

  if err != nil {
    return false
  }

  return resp.StatusCode == 200
}

func (h Http) Name() string {
  return h.name
}

func runPlugins(p []Plugin) {
  for i := 0; i < len(p); i++ {
    success := p[i].Run()
    if success {
      fmt.Println(p[i].Name(), "is up.")
    } else {
      fmt.Println(p[i].Name(), "is down.")
    }
  }
}

func main() {
  plugins := []Plugin{Http{"https://archlinux.org", "Arch Linux"}}
  runPlugins(plugins)
}

