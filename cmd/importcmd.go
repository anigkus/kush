// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/anigkus/kush/entity"
	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// ImportCmdRun export command cmd entry
func ImportCmdRun(option entity.Option) {
	if len(option.Output) <= 0 {
		option.Output = sys.DEFAULT_OUTPUT
	}

	bufbytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		util.ExitPrintln(fmt.Sprintf("error: %v", "stdin err"))
	}
	err = util.ImportDataPathFile(bufbytes, option.Output)
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	}
}
