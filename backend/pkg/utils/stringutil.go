package utils

import "strconv"

func StrToUint32(idStr string) (uint32, error) {
	idU64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}

	id := uint32(idU64)
	return id, nil
}
