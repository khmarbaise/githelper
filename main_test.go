// Copyright 2020 The GitHelper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main_test

import (
	"github.com/khmarbaise/githelper/modules/execute"
	"testing"
)

//Test_Main_first Integration test to execute our own executable within a test.
func Test_Main_first(t *testing.T) {
	t.Run("Execute githelper with --help", func(t *testing.T) {
		//Execute our own produced executable for testing purposes.
		execute.ExternalCommand("./githelper", "--help")
	})
}
