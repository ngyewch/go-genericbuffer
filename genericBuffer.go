package genericbuffer

type GenericBuffer[T any] struct {
	data []T
}

func NewGenericBuffer[T any]() *GenericBuffer[T] {
	return &GenericBuffer[T]{
		data: make([]T, 0),
	}
}

func (genericBuffer *GenericBuffer[T]) Len() int {
	return len(genericBuffer.data)
}

func (genericBuffer *GenericBuffer[T]) Append(data []T) {
	if len(data) == 0 {
		return
	}
	genericBuffer.data = append(genericBuffer.data, data...)
}

func (genericBuffer *GenericBuffer[T]) Next(n int) []T {
	if n <= 0 {
		return nil
	}
	if n > len(genericBuffer.data) {
		n = len(genericBuffer.data)
	}
	b := make([]T, n)
	copy(b, genericBuffer.data[0:n])
	remainingDataLen := len(genericBuffer.data) - n
	copy(genericBuffer.data[0:remainingDataLen], genericBuffer.data[n:])
	genericBuffer.data = genericBuffer.data[:remainingDataLen]
	return b
}

func (genericBuffer *GenericBuffer[T]) Skip(n int) {
	if n <= 0 {
		return
	}
	if n > len(genericBuffer.data) {
		n = len(genericBuffer.data)
	}
	remainingDataLen := len(genericBuffer.data) - n
	copy(genericBuffer.data[0:remainingDataLen], genericBuffer.data[n:])
	genericBuffer.data = genericBuffer.data[:remainingDataLen]
}

func (genericBuffer *GenericBuffer[T]) Peek(n int) []T {
	if n <= 0 {
		return nil
	}
	if n > len(genericBuffer.data) {
		n = len(genericBuffer.data)
	}
	return genericBuffer.data[0:n]
}
