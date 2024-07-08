package transformer

func Noop(content []byte) ([]byte, error) {
	return content, nil
}
