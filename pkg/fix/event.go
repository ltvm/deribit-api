package fix

import "github.com/chuckpreslar/emission"

// On adds a listener to a specific event.
func (c *Client) On(event interface{}, listener interface{}) *emission.Emitter {
	return c.emitter.On(event, listener)
}

// Emit emits an event.
func (c *Client) Emit(event interface{}, args ...interface{}) *emission.Emitter {
	return c.emitter.Emit(event, args...)
}

// Off removes a listener for an event.
func (c *Client) Off(event interface{}, listener interface{}) *emission.Emitter {
	return c.emitter.Off(event, listener)
}
