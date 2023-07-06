package rtp

func Ptr[T any](v T) *T {
	return &v
}
func Arr[T any](v T) []*T {
	var arrayOfT []*T
	arrayOfT = append(arrayOfT, &v)
	return arrayOfT
}
