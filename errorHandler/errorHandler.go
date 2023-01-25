package errorhandler

func HandleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func HandleErrorWithOutError(msg ...string) {
	if len(msg) != 0 {
		panic(msg[0])
	}
}

// HandleError(err)
