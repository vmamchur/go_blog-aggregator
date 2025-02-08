package main

import (
    "fmt"
    "log"

	"github.com/vmamchur/go_blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
    if err != nil {
        log.Fatalf("Error reading config: %v", err)
    }
    fmt.Printf("Read config: %+v\n", cfg)

    err = cfg.SetUser("bugsworld")

    cfg, err = config.Read()
    if err != nil {
        log.Fatalf("Error reading config: %v", err)
    }
    fmt.Printf("Read config again: %+v\n", cfg)
}
