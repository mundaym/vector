package vector

func SumGeneric(v []int16) int16 {
	res := int16(0)
	for _, val := range v {
		res += val
	}
	return res
}
