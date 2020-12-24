package app

import (
	"bst/pkg/log"
	"bst/pkg/model"
	"github.com/sirupsen/logrus"
	"sync"
)

type Tree interface {
	Search(key int) (*model.Node, error)
	Insert(key int) (*model.Node, error)
	Delete(key int) error
}

type App struct {
	tree Tree
	log  *logrus.Logger
	mu   sync.RWMutex
}

func New(tree Tree) *App {
	return &App{tree: tree, log: log.New()}
}

func (app *App) Search(value int) (*model.Node, error) {
	app.mu.RLock()
	defer app.mu.RUnlock()
	app.log.WithField("key", value).Infof("search")
	return app.tree.Search(value)
}

func (app *App) Insert(value int) (*model.Node, error) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.log.WithField("key", value).Infof("insert")
	return app.tree.Insert(value)
}

func (app *App) Delete(value int) error {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.log.WithField("key", value).Infof("delete")
	return app.tree.Delete(value)
}
