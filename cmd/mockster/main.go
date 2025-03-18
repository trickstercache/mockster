/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
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

// Package main is the main package for the Mockster application
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/trickstercache/mockster/pkg/routes"
)

const (
	applicationName    = "mockster"
	applicationVersion = "1.1.3"
)

func main() {

	port := "8482"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}

	fmt.Println("Starting up", applicationName, applicationVersion, "on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), routes.GetRouter())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
