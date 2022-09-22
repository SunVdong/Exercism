package erratum

import (
	"errors"
)

func Use(opener ResourceOpener, input string) (res_err error) {
	res, err := opener()
	for err != nil && errors.As(err, &TransientError{}) {
		res, err = opener()
	}
	if err != nil {
		res_err = err
		return
	}

	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				res_err = e.inner
			default:
				res_err = r.(error)
			}
		}

		res.Close()

		// msg := recover()
		// er, ok := msg.(FrobError)
		// if ok {
		// 	res.Defrob(er.defrobTag)
		// 	res_err = er
		// } else {
		// 	res_err, _ = msg.(error)
		// }
		// res.Close()
	}()

	res.Frob(input)
	return res_err
}
