package dispatcher

// ID is a unique identifier representing a registered callback function.
type ID int
type callback func(action interface{})

//Dispatcher ??
type Dispatcher struct {
	idCounter ID
	callbacks map[ID]callback
}

//New ??
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
	return id
}

// Unregister unregisters the callback previously registered via a call to
// Register.
func (d *Dispatcher) Unregister(id ID) {
	delete(d.callbacks, id)
}
