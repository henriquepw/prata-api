package apid

import "github.com/nrednav/cuid2"

func createID(size int) string {
	generate, err := cuid2.Init(
		cuid2.WithLength(size),
		cuid2.WithFingerprint("pobrin-api"),
	)
	if err != nil {
		panic(err)
	}

	return generate()
}

func New() string {
	return createID(24)
}

func NewTiny() string {
	return createID(6)
}
