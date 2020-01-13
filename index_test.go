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

package indexer_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/go-indexer"
)

func TestIndex(t *testing.T) {
	index := indexer.New()

	id := uuid.New()
	name := "Test name"

	assert.False(t, index.IDKnown(id))
	assert.False(t, index.NameKnown(name))
	index.Add(id, name)
	assert.True(t, index.IDKnown(id))
	assert.True(t, index.NameKnown(name))

	foundID, exists := index.ID(name)
	assert.True(t, exists)
	assert.Equal(t, id, foundID)

	foundName, exists := index.Name(id)
	assert.True(t, exists)
	assert.Equal(t, name, foundName)

	index.Remove(id, name)

	assert.False(t, index.IDKnown(id))
	assert.False(t, index.NameKnown(name))
}

func TestIndexSerDeser(t *testing.T) {
	index := indexer.New()

	id1 := uuid.New()
	name1 := "Test name 1"
	index.Add(id1, name1)
	id2 := uuid.New()
	name2 := "Test name 2"
	index.Add(id2, name2)

	ser, err := index.Serialize()
	require.Nil(t, err)

	index, err = indexer.Deserialize(ser)
	require.Nil(t, err)

	foundName, exists := index.Name(id1)
	assert.True(t, exists)
	assert.Equal(t, foundName, name1)
	foundID, exists := index.ID(name1)
	assert.True(t, exists)
	assert.Equal(t, foundID, id1)

	foundName, exists = index.Name(id2)
	assert.True(t, exists)
	assert.Equal(t, foundName, name2)
	foundID, exists = index.ID(name2)
	assert.True(t, exists)
	assert.Equal(t, foundID, id2)
}
