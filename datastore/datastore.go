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

	"github.com/jinzhu/gorm"
)

// DebugQueries define the debug
const (
	DebugQueries = iota
)

// Define common vars
var (
	Debug = false
)

type DataStore interface {
	Source() string
	Driver() string
}

type BaseModel struct {
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt Time   `gorm:"column:createdAt;default:'0000-00-00 00:00:00'" json:"createdAt"`
	UpdatedAt Time   `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *Time  `gorm:"column:deletedAt;default:NULL" sql:"index" json:"-"`
}

type ds struct {
	alias *alias
	*gorm.DB
	isTx bool
}

func (o *ds) Using(name string) *ds {
	if o.isTx {
		panic(fmt.Errorf("<ds.Using> transaction has been start, cannot change db"))
	}
	if al, ok := dataBaseCache.get(name); ok {
		o.alias = al
		o.DB = al.DB
		if Debug {
			o.Debug()
		}
	} else {
		panic(fmt.Errorf("<ds.Using> unknown db alias name `%s`", name))
	}

	return o
}

// return current using database Driver
func (o *ds) Driver() string {
	return o.alias.DriverName
}

// New create new ds
func New() *ds {
	o := new(ds)
	o.Using("default")
	return o
}

// NewWithDB create a new ds object with specify *gorm.DB for query
func NewWithDB(driverName, aliasName string, db *gorm.DB) (*ds, error) {
	var al *alias

	if dr, ok := drivers[driverName]; ok {
		al = new(alias)
		al.Driver = dr
	} else {
		return nil, fmt.Errorf("driver name `%s` have not registered", driverName)
	}

	al.Name = aliasName
	al.DriverName = driverName

	o := new(ds)
	o.alias = al

	if Debug {
		o.Debug()
	}

	return o, nil
}

func (o *ds) Index(index ...int64) *ds {
	for i, v := range index {
		switch i {
		case 0:
			if v > 0 {
				o.DB = o.Offset(v)
			}
		case 1:
			if v > 0 {
				o.DB = o.Limit(v)
			}
		}
	}

	return o
}
