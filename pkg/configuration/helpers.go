
/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package configuration

import (
	"net"
	"time"
)

// portChecker just verifies that specific ports are open, used for tests
// ref: https://stackoverflow.com/a/56336811/4949938
func portChecker(host string, ports ...string) error {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			return err
		}
		if conn != nil {
			err := conn.Close()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
