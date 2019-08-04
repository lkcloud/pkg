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

package datastore

import (
	"fmt"
)

type MySQLStore struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Name     string `json:"name"`
}

func NewMySQLStore(username, password, address, name string) DataStore {
	return &MySQLStore{
		Username: username,
		Password: password,
		Address:  address,
		Name:     name,
	}
}

func (ds *MySQLStore) Source() string {
	config := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		ds.Username,
		ds.Password,
		ds.Address,
		ds.Name,
		true,
		"Local")

	return config
}

func (m *MySQLStore) Driver() string {
	return "mysql"
}
