package conv

import "strconv"

func StringToUint(v string) (uint, error) {
	id, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func UintToString(v uint) string {
	return strconv.Itoa(int(v))
}
