package wraperr

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %s", msg, err)
}

func WrapIfErr(msg string, err error) error {
	if err != nil {
		return Wrap(msg, err)
	}
	return nil
}