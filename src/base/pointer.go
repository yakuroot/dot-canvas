package base

func Pointer[T any](elem T) *T { return &elem }
