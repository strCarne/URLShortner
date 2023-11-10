package main

import (
	"fmt"

	"github.com/strCarne/URLShortner/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}