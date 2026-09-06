package jaccl

import "errors"

// closeStack releases resources in reverse creation order.
//
// It is used by the native backend so a failed setup follows the same resource
// lifetime discipline as a normally closed group.
type closeStack struct {
	closers []func() error
}

func (s *closeStack) add(close func() error) {
	if close != nil {
		s.closers = append(s.closers, close)
	}
}

func (s *closeStack) close() error {
	if s == nil {
		return nil
	}
	var errs []error
	for i := len(s.closers) - 1; i >= 0; i-- {
		if err := s.closers[i](); err != nil {
			errs = append(errs, err)
		}
	}
	s.closers = nil
	return errors.Join(errs...)
}
