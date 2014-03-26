package notities

type kladblok struct {
	text []byte
}

func (k *kladblok) Write(p []byte) (int, error) {
	k.text = append(k.text, p...)
	return len(p), nil
}
