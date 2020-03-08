package errorhelper

import "fmt"

func WrapError(e error, wrap error) error {
	return Wrap(e.Error(), wrap)
}

func Wrap(msg string, wrap error) error {
	return fmt.Errorf("%v: %w", msg, wrap)
}

func Wrapf(wrap error, msg string, args ...interface{}) error {
	f := fmt.Sprintf(msg, args...)
	return fmt.Errorf("%v: %w", f, wrap)
}
