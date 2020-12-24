package app

import (
	"bst/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testArray = []int{1, 2, 4, 5, 6, 23, 3}

func TestApp_Search(t *testing.T) {
	tree := model.NewTree(testArray)
	app := New(tree)
	_, err := app.Search(22)
	assert.Equal(t, model.ErrNotFound, err)

	key := 3
	node, err := app.Search(key)
	if assert.NoError(t, err) {
		assert.Equal(t, key, node.Key)
	}
}

func TestApp_Insert(t *testing.T) {
	tree := model.NewTree(testArray)
	app := New(tree)
	existedKey := 1
	_, err := app.Insert(existedKey)
	assert.Equal(t, model.ErrAlreadyExists, err)
	newKey := 7
	node, err := app.Insert(newKey)
	if assert.NoError(t, err) {
		assert.Equal(t, newKey, node.Key)
	}
}

func TestApp_Delete(t *testing.T) {
	tree := model.NewTree(testArray)
	app := New(tree)
	key := -4
	_, err := app.Insert(key)
	assert.NoError(t, err)
	err = app.Delete(key)
	assert.NoError(t, err)
	_, err = app.Search(key)
	assert.Equal(t, model.ErrNotFound, err)
}
