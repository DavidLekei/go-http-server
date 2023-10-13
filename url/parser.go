package url

type URLParser interface {
	Create() URLParser
	Validate(url string) bool
	ParsePath(url string) string
	ParseArgs(url string) map[string]string
}

type RESTUrlParser struct {
	pathTree *BSTNode
	url *URL
}

func (parser *RESTUrlParser) Create() URLParser {
	return &RESTUrlParser{
		pathTree: BSTCreate("test data", false),
		url: URLCreate("TODO"),
	}
}

func (parser *RESTUrlParser) Validate(url string) bool {
	return true
}

func (parser *RESTUrlParser) ParsePath(url string) string {
	return "Parsed URL Path"
}

func (parser *RESTUrlParser) ParseArgs(url string) map[string]string {
	return make(map[string]string)
}