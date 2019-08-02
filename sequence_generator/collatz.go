package sequence_generator

import "errors"

func Next(seed int) (int, error) {

	if seed <= 0 {
		return 0, errors.New("invalid value")
	}

	if seed == 1 {
		return 0, errors.New("already finished")
	}

	result := 1

	if seed%2 != 0 {
		result = seed*3 + 1
	}

	if seed%2 == 0 {
		result = seed / 2
	}

	return result, nil
}

func Sequence(seed int) ([]int, error) {
	result := make([]int, 0)

	for nextValue, _ := Next(seed); nextValue != 1; nextValue, _ = Next(nextValue) {
		result = append(result, nextValue)
	}

	result = append(result, 1)

	return result, nil
}
