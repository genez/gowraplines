// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"flag"
	"io"
	"os"
)

var operationFlag = flag.String("operation", "", "wrap or unwrap lines")
var columnWidthFlag = flag.Int("columnWidth", 0, "sets column width for line breaks")
var separatorFlag = flag.String("separator", "\r\n", "string to use as separator")

func main() {
	flag.Parse()

	if *operationFlag == "" {
		flag.Usage()
		return
	}

	if *columnWidthFlag == 0 {
		flag.Usage()
		return
	}

	columnWidth := int64(*columnWidthFlag)
	separator := *separatorFlag
	operation := *operationFlag

	if operation == "wrap" {
		for true {
			written, err := io.CopyN(os.Stdout, os.Stdin, columnWidth)
			if written == columnWidth && err == nil {
				io.WriteString(os.Stdout, separator)
			} else {
				break
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			io.WriteString(os.Stdout, scanner.Text())
		}
	}
}
