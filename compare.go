package timezone

import "time"

func CompareEQ(timestampFrom, timestampTo string, layout string) (bool, error) {
	from, err := time.ParseInLocation(layout, timestampFrom, zoneLocal)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, zoneLocal)
	if err != nil {
		return false, err
	}
	return from.Equal(to), nil
}

func CompareSecondEQ(timestampFrom, timestampTo string) (bool, error) {
	return CompareEQ(timestampFrom, timestampTo, LayoutSecond)
}

func CompareMicroEQ(timestampFrom, timestampTo string) (bool, error) {
	return CompareEQ(timestampFrom, timestampTo, LayoutMicro)
}

func CompareLT(timestampFrom, timestampTo string, layout string) (bool, error) {
	from, err := time.ParseInLocation(layout, timestampFrom, zoneLocal)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, zoneLocal)
	if err != nil {
		return false, err
	}
	return from.Before(to), nil
}

func CompareSecondLT(timestampFrom, timestampTo string) (bool, error) {
	return CompareLT(timestampFrom, timestampTo, LayoutSecond)
}

func CompareMicroLT(timestampFrom, timestampTo string) (bool, error) {
	return CompareLT(timestampFrom, timestampTo, LayoutMicro)
}

func CompareGT(timestampFrom, timestampTo string, layout string) (bool, error) {
	from, err := time.ParseInLocation(layout, timestampFrom, zoneLocal)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, zoneLocal)
	if err != nil {
		return false, err
	}
	return from.After(to), nil
}

func CompareSecondGT(timestampFrom, timestampTo string) (bool, error) {
	return CompareGT(timestampFrom, timestampTo, LayoutSecond)
}

func CompareMicroGT(timestampFrom, timestampTo string) (bool, error) {
	return CompareGT(timestampFrom, timestampTo, LayoutMicro)
}
