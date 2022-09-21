package erratum

import (
	"errors"
)

func Use(opener ResourceOpener, input string) (a error) {
	var res Resource
	var err error

	res, err = opener()
	for err != nil && errors.As(err, &TransientError{}) {
		res, err = opener()
	}
	if err != nil {
		return err
	}
	defer res.Close()

	defer func() {
		if e := recover(); e != nil {
			if er, ok := e.(*FrobError); ok {
				res.Defrob(er.defrobTag)
				a = er
				return
			} else {
				a = errors.New("aaa")
				return
			}
		} else {
			return
		}
	}()

	res.Frob(input)
	return
}
