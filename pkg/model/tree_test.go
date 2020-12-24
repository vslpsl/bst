package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testArray = []int{1, 2, 3, 4, 5, 10, 20, 30}

func TestTree_Search(t *testing.T) {
	key := 3
	tree := NewTree(testArray)
	node, err := tree.Search(key)
	if assert.NoError(t, err) {
		assert.Equal(t, key, node.Key)
	}
}

func TestTree_Insert(t *testing.T) {
	keyToInsert1 := 5
	keyToInsert2 := 25
	tree := NewTree(testArray)
	_, err := tree.Insert(keyToInsert1)
	assert.Equal(t, ErrAlreadyExists, err)

	node, err := tree.Insert(keyToInsert2)
	if assert.NoError(t, err) {
		assert.Equal(t, keyToInsert2, node.Key)
	}
}

func TestTree_Delete(t *testing.T) {
	keyToDelete := 22
	tree := NewTree(testArray)
	_, err := tree.Insert(keyToDelete)
	assert.NoError(t, err)
	err = tree.Delete(keyToDelete)
	assert.NoError(t, err)
}
