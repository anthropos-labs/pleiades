/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package runtime

import (
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, message string) {
	errorStr := fmt.Sprintf("[ Failed ] %v\n", message)
	fmt.Printf(errorStr)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errorStr))
}
