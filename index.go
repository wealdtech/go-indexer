// Copyright 2020 Weald Technology Trading
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

package indexer

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Index maps between names and IDs.
type Index struct {
	names map[uuid.UUID]string
	ids   map[string]uuid.UUID
}

// New creates a new index.
func New() *Index {
	return &Index{
		names: make(map[uuid.UUID]string),
		ids:   make(map[string]uuid.UUID),
	}
}

// Add adds an entry to this index.
func (i *Index) Add(id uuid.UUID, name string) {
	i.names[id] = name
	i.ids[name] = id
}

// Remove removes an entry from this index.
func (i *Index) Remove(id uuid.UUID, name string) {
	delete(i.names, id)
	delete(i.ids, name)
}

// Name fetches the name of an entry given its ID.
// If present the second return value will be true, otherwise false.
func (i *Index) Name(id uuid.UUID) (string, bool) {
	res, exists := i.names[id]
	return res, exists
}

// NameKnown returns true if this name is known.
func (i *Index) NameKnown(name string) bool {
	_, exists := i.ids[name]
	return exists
}

// ID fetches the ID of an entry given its name.
// If present the second return value will be true, otherwise false.
func (i *Index) ID(name string) (uuid.UUID, bool) {
	res, exists := i.ids[name]
	return res, exists
}

// IDKnown returns true if this ID is known.
func (i *Index) IDKnown(id uuid.UUID) bool {
	_, exists := i.names[id]
	return exists
}

type indexEntry struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

// Deserialize deserializes a serialized index.
func Deserialize(data []byte) (*Index, error) {
	var entries []*indexEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	index := New()
	for _, entry := range entries {
		index.Add(entry.ID, entry.Name)
	}
	return index, nil
}

// Serialize serializes an index.
func (i *Index) Serialize() ([]byte, error) {
	entries := make([]*indexEntry, 0)
	for k, v := range i.names {
		entries = append(entries, &indexEntry{
			ID:   k,
			Name: v,
		})
	}
	return json.Marshal(entries)
}
