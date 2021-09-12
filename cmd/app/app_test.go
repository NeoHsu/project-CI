package main

import (
	"testing"

	bundle "github.com/NeoHsu/project-ci/internal/bundle"
	"go.uber.org/fx/fxtest"
)

func TestDependenciesAreSatisfied(t *testing.T) {
	fxtest.New(t, bundle.Module).RequireStart().RequireStop()
}
