// Kiebitz - Privacy-Friendly Appointment Scheduling
// Copyright (C) 2021-2021 The Kiebitz Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package services

import (
	"time"
)

type DatabaseDefinition struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Maker             DatabaseMaker     `json:"-"`
	SettingsValidator SettingsValidator `json:"-"`
}

type SettingsValidator func(settings map[string]interface{}) (interface{}, error)
type DatabaseDefinitions map[string]DatabaseDefinition
type DatabaseMaker func(settings interface{}) (Database, error)

type DatabaseOps interface {
	Set(table string, key []byte) Set
	SortedSet(table string, key []byte) SortedSet
	List(table string, key []byte) List
	Map(table string, key []byte) Map
	Value(table string, key []byte) Value
}

// A database can deliver and accept message
type Database interface {
	Close() error
	Open() error
	Begin() (Transaction, error)

	DatabaseOps
}

type Transaction interface {
	// Watch keys for changes while in a transaction
	Watch(...string) error

	Commit() error
	Rollback() error

	DatabaseOps
}

type Object interface {
}

type Set interface {
	Add([]byte) error
	Has([]byte) (bool, error)
	Del(key []byte) error
	Object
}

type SortedSetEntry struct {
	Score int64
	Data  []byte
}

type SortedSet interface {
	Object
	Del([]byte) error
	Add([]byte, int64) error
	PopMin(int64) ([]*SortedSetEntry, error)
}

type List interface {
	Object
}

type Map interface {
	GetAll() (map[string][]byte, error)
	Get(key []byte) ([]byte, error)
	Del(key []byte) error
	Set(key []byte, value []byte) error
	Object
}

type Value interface {
	Object
	Set(value []byte, ttl time.Duration) error
	Get() ([]byte, error)
	Del() error
}

type BaseDatabase struct {
}
