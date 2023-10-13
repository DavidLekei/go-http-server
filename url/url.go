package url

type URL struct {
	protocol string
	domain string
	params  map[string]string
	port int
	path string
}

func URLCreate(url string) *URL{
	return &URL{
		protocol: "TODO",
		domain: "TODO",
		port: 3456,
		params: make(map[string]string),
		path: "TODO",
	}
}