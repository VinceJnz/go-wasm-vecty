package dispatcher

import "log"

const debugTag = "dispatcher_"

// ID is a unique identifier representing a registered callback function.
type ID int

// Callback stores a single callback function that can be called by the dispatcher
type callback func(action interface{})

//Dispatcher structure used by the dispatcher for storing a map of callbacks
type Dispatcher struct {
	idCounter ID
	callbacks map[ID]callback
}

// New retuns a new instance of Dispatcher
func New() *Dispatcher {
	d := new(Dispatcher)
	d.callbacks = make(map[ID]callback)
	d.idCounter = 0
	return d
}

// Dispatch dispatches the given action to all registered callbacks.
func (d *Dispatcher) Dispatch(action interface{}) {
	for _, c := range d.callbacks {
		c(action)
	}
}

// Register registers the callback to handle dispatched actions, the returned
// ID may be used to unregister the callback later.
func (d *Dispatcher) Register(callback callback) ID {
	d.idCounter++
	id := d.idCounter
	d.callbacks[id] = callback
	log.Println(debugTag+"Register1 ", "d =", d)
	return id
}

// Unregister unregisters the callback previously registered via a call to
// Register.
func (d *Dispatcher) Unregister(id ID) {
	delete(d.callbacks, id)
}
