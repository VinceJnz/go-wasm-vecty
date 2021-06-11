package store

import (
	"github.com/VinceJnz/go-wasm-vecty/actions"
	"github.com/VinceJnz/go-wasm-vecty/dispatcher"
	"github.com/VinceJnz/go-wasm-vecty/store/model"
	"github.com/VinceJnz/go-wasm-vecty/store/storeutil"
)

const debugTag = "store_"

// Store is the data and utilities used for managing and accessing the store
type Store struct {
	// Items represents all of the TODO items in the store.
	Items []*model.Item

	// Filter represents the active viewing filter for items.
	Filter model.FilterState

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners *storeutil.ListenerRegistry

	//Dispatcher is a pointer to the dispatcher
	Dispatcher *dispatcher.Dispatcher
}

// New returns a new instance of a Store
func New() *Store {
	s := new(Store)
	s.Filter = model.All
	s.Listeners = storeutil.NewListenerRegistry()
	s.Dispatcher = dispatcher.New()
	s.Dispatcher.Register(s.onAction)
	return s
}

// ActiveItemCount returns the current number of items that are not completed.
func (s *Store) ActiveItemCount() int {
	return s.count(false)
}

// CompletedItemCount returns the current number of items that are completed.
func (s *Store) CompletedItemCount() int {
	return s.count(true)
}

func (s *Store) count(completed bool) int {
	count := 0
	for _, item := range s.Items {
		if item.Completed == completed {
			count++
		}
	}
	return count
}

func (s *Store) onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.ReplaceItems:
		s.Items = a.Items

	case *actions.AddItem:
		s.Items = append(s.Items, &model.Item{Title: a.Title, Completed: false})

	case *actions.DestroyItem:
		copy(s.Items[a.Index:], s.Items[a.Index+1:])
		s.Items = s.Items[:len(s.Items)-1]

	case *actions.SetTitle:
		s.Items[a.Index].Title = a.Title

	case *actions.SetCompleted:
		s.Items[a.Index].Completed = a.Completed

	case *actions.SetAllCompleted:
		for _, item := range s.Items {
			item.Completed = a.Completed
		}

	case *actions.ClearCompleted:
		var activeItems []*model.Item
		for _, item := range s.Items {
			if !item.Completed {
				activeItems = append(activeItems, item)
			}
		}
		s.Items = activeItems

	case *actions.SetFilter:
		s.Filter = a.Filter

	default:
		return // don't fire listeners
	}

	s.Listeners.Fire()
}
