package uint32sets

type set interface {
	add(uint32)
	contains(uint32) bool
	length() int
}
