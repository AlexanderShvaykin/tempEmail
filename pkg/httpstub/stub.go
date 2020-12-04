package httpstub

type HttpClient struct {
	Response string
}

func (c HttpClient) Get(_ string, _ map[string]string) ([]byte, error) {
	return []byte(c.Response), nil
}
