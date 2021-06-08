//  Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package main

import (
	"github.com/kris-nova/logger"
	photoprism "github.com/kris-nova/photoprism-client-go"
)

func main() {
	logger.Level = 4
	client := photoprism.New("http://localhost:8080")
	err := client.Auth(photoprism.NewClientAuthLogin("admin", "missy"))
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	logger.Always("Login Success!")
}
