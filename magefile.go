// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Build() error {
	return gopherJs("build", "-o", "app.js", ".")
}

func InstallDeps() error {
	if err := goMod("vendor"); err != nil {
		return err
	}

	if err := goGet("-u", "github.com/gopherjs/gopherjs"); err != nil {
		return err
	}

	return npm("i")
}

func npm(args ...string) error {
	return sh.Run("npm", args...)
}

func goMod(args ...string) error {
	args = append(append([]string{}, "mod"), args...)
	return sh.RunWith(map[string]string{"GO111MODULE": "on"}, "go", args...)
}

func gopherJs(args ...string) error {
	return sh.RunWith(map[string]string{"GO111MODULE": "on"}, "gopherjs", args...)
}

func goGet(args ...string) error {
	args = append(append([]string{}, "get"), args...)
	return sh.Run( "go", args...)
}
