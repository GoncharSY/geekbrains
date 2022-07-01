package main

import (
	"fmt"
	"go-level-1/lesson-8/config"
	"os"
)

func init() {
	os.Setenv("WEBAPP_SOME_ID", "ENV_ID")
	os.Setenv("WEBAPP_SOME_KEY", "ENV_KEY")
}

func main() {
	var cfg = config.New()
	var flg = cfg.GetFlagConfig()
	var env = cfg.GetEnvConfig()
	var dft = cfg.GetDefaultConfig(false)
	var rst = cfg.GetResultConfig(false)

	if errs := cfg.Validate(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("config error:", err)
		}

		panic("incorrect config")
	}

	fmt.Println()
	fmt.Println("    Default config:", dft)
	fmt.Println("Environment config:", env)
	fmt.Println("       Flag config:", flg)
	fmt.Println("     Result config:", rst)
	fmt.Println()
}
