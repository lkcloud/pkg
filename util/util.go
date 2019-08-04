/*
 * Copyright © 2019 Lingfei Kong <466701708@qq.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
)

// EnsureDirExists make sure directory is exist
func EnsureDirExists(dir string) (err error) {
	f, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dir, os.FileMode(0755))
		} else {
			return err
		}
	}
	if !f.IsDir() {
		return fmt.Errorf("path %s is exist,but not dir", dir)
	}

	return nil
}

// AggregateGoRequestErrors combine gorequest.Response, body, errors and return an representative error
func AggregateGoRequestErrors(resp gorequest.Response, body string, errs []error) error {
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}

	if resp == nil {
		return fmt.Errorf("gorequest response is nil.")
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("gorequest return HTTP code >= 400, body: `%s`", body)
	}

	return nil
}

// DelFromSlice delete elements from a given string slice
func DelFromSlice(slice []string, elems ...string) []string {
	isInElems := make(map[string]bool)
	for _, elem := range elems {
		isInElems[elem] = true
	}

	w := 0
	for _, elem := range slice {
		if !isInElems[elem] {
			slice[w] = elem
			w += 1
		}
	}

	return slice[:w]
}

// GetLocalAddress return the ip address of the host
func GetLocalAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", nil
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", nil
}

func EnsureGetLocalAddress() (string, error) {
	// get available network interfaces for this machine
	interfaces, err := net.Interfaces()

	if err != nil {
		return GetLocalAddress()
	}

	ips := make(map[string]string)
	for _, i := range interfaces {
		byNameInterface, err := net.InterfaceByName(i.Name)
		if err != nil {
			return GetLocalAddress()
		}

		addresses, err := byNameInterface.Addrs()
		if len(addresses) > 0 {
			ips[i.Name] = strings.Split(addresses[0].String(), "/")[0]
		}
	}
	for _, ifname := range []string{"br1", "bond1", "eth1", "eth0"} {
		if value, ok := ips[ifname]; ok {
			return value, nil
		}
	}

	return GetLocalAddress()
}

func SingleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
