package actuator

import "net/url"

type Actuator interface {
	Signal(origin string, values url.Values) error
}

type HttpActuator struct{}

func (HttpActuator) Signal(origin string, values url.Values) error {
	return nil
}
