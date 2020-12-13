package httpstub

type HttpClient struct {
	Response string
	Args     map[string]string
	Url      string
}

func (c *HttpClient) Get(url string, args map[string]string) ([]byte, error) {
	c.Args = args
	c.Url = url

	return []byte(c.Response), nil
}
