package util

import "strconv"

func StringToUInt64(str string) (uint64, error) {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
