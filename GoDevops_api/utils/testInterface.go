package utils

type Iwktest interface {
	Create()
	Update()
}

type wktest struct {
}

func (w *wktest) Error() string {
	panic("not implemented") // TODO: Implement
}
