package main

import (
	"fmt"

	bundle "github.com/NeoHsu/project-ci/internal/bundle"
	"go.uber.org/fx"
)

var (
	CommitHash string
	BuildTime  string
)

func main() {
	fmt.Printf("CommitHash: %s, BuildTime: %s \n", CommitHash, BuildTime)
	fx.New(
		bundle.Module,
	).Run()
}
