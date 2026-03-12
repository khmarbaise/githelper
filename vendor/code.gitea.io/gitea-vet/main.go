// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"code.gitea.io/gitea-vet/checks"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		checks.Imports,
		checks.License,
		checks.Migrations,
		checks.ModelsSession,
	)
}
