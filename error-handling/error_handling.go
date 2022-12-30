package erratum

func Use(opener ResourceOpener, input string) (final error) {
	var (
		resource Resource
	)
	// close non-nil resource in all cases
	defer func() {
		if resource != nil {
			resource.Close()
		}
	}()

	// retry until error is not a TransientError
	for {
		resource, final = opener()
		if _, ok := final.(TransientError); ok {
			continue
		}
		break
	}

	// if there's still an error, return
	if final != nil {
		return
	}

	// recover from panics that can come from the call to Frob
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(FrobError); ok {
				// defrob if it's a frob error
				resource.Defrob(err.defrobTag)
				final = err
			}
			// return any other error
			final = err.(error)
		}
	}()
	resource.Frob(input)

	return
}
