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

func WrapOp(op string, msg error, err error) error {
	return fmt.Errorf("PLACE: %s, MSG: %s, ERR: %s", op, msg, err)
}

func WrapOpIfErr(op string, msg error, err error) error {
	if err != nil {
		return WrapOp(op, msg, err)
	}
	return nil
}
