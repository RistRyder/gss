package io

func GetSplitter(tokenLength int) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if len(data) < tokenLength {
			if atEOF {
				return len(data), data, nil
			}

			return 0, nil, nil
		}

		return tokenLength, data[0:tokenLength], nil
	}
}
