// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024 Nhi Pham

// This program decodes UEFI boot progress and error codes from EDK2.
// It supports the following formats:
//
// Progress Code: PROGRESS CODE: V03020003 I0
// Error Code:    ERROR: C000000002:V03058002 I0 6D33944A-EC75-4855-A54D-809C75241F6C
//
// The decoder extracts and interprets the hexadecimal codes, providing
// human-readable descriptions for class, subclass, and operation.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nhivp/boot-progress-decoder/pkg/edk2"
)

func handleProgressCode(progressCode string) {
	progressCode = strings.TrimSpace(progressCode)
	statusCodeValue := strings.TrimPrefix(progressCode, "PROGRESS CODE:")
	statusCodeValue = strings.TrimSuffix(statusCodeValue, "I0")

	classDesc, subclassDesc, operationDesc, _ := edk2.DecodeStatusValue(statusCodeValue, false)

	fmt.Println(progressCode)
	fmt.Println("Class     : ", classDesc)
	fmt.Println("Subclass  : ", subclassDesc)
	fmt.Println("Operation : ", operationDesc)
}

func handleErrorCode(statusCode string) {
	statusCode = strings.TrimSpace(statusCode)
	errorCode := strings.TrimPrefix(statusCode, "ERROR: ")

	parts := strings.Split(errorCode, " ")

	if len(parts) != 3 {
		fmt.Println("invalid error code format")
		return
	}

	callerID := parts[len(parts)-1]
	parts = strings.Split(parts[0], ":")
	if len(parts) != 2 {
		fmt.Println("invalid error code format")
		return
	}

	statusCodeType := parts[0]
	statusCodeValue := parts[1]

	classDesc, subclassDesc, operationDesc, _ := edk2.DecodeStatusValue(statusCodeValue, true)

	_, severityDesc, _ := edk2.DecodeStatusType(statusCodeType)

	fmt.Println(errorCode)
	fmt.Println("Severity  : ", severityDesc)
	fmt.Println("Class     : ", classDesc)
	fmt.Println("Subclass  : ", subclassDesc)
	fmt.Println("Operation : ", operationDesc)
	fmt.Println("Module    : ", callerID)
}

func helpString() string {
	return `Usage: boot-progress-decoder <PROGRESS CODE line | ERROR line>

Decodes a single UEFI boot progress or error code line.

The input should be a single line in one of the following formats:
  - Progress codes: PROGRESS CODE: V<hex_code> ...
  - Error codes: ERROR: C<status_code_type>:V<hex_code> ...

Examples:
  boot-progress-decoder "PROGRESS CODE: V03020003 I0"
  boot-progress-decoder "ERROR: C40000002:V010E0005 I0 55E3774A-EB45-4FD2-AAAE-B7DEEB504A0E"

Output:
  The decoded information will be displayed in a formatted table for the input line.`
}

func main() {
	// Check if there are enough command-line arguments
	// If not, print the help message and exit
	if len(os.Args) < 2 {
		fmt.Println(helpString())
		return
	}

	// Get the status code from the first command-line argument
	statusCode := os.Args[1]

	// Check if it's a progress code or error code based on the prefix
	if strings.HasPrefix(statusCode, "PROGRESS CODE:") {
		handleProgressCode(statusCode)
	} else if strings.HasPrefix(statusCode, "ERROR:") {
		handleErrorCode(statusCode)
	} else {
		fmt.Println("Invalid input line. Must start with 'PROGRESS CODE:' or 'ERROR:'.")
	}
}
