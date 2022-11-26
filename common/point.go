package common

// NewPoint 创建指针
func NewPoint[T any](v T) *T {
	return &v
}
