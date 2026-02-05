// config/constants.go
package config

const DebugIsOn = true

var (
	FileURLs               = "config/urls.txt"
	FileURLsForTesting     = "config/test-urls.txt"
	FilePaths              = "config/paths.txt"
	FilePathsForTesting    = "config/test-paths.txt"
	YAMLFileName           = "config/config.yaml"
	YAMLFileNameForTesting = "config/test-config.yaml"
)

type ContentType int

const (
	ContentTypeTextPlain ContentType = iota
	ContentTypeXml
	ContentTypeJson
	ContentTypeForm
)

func (c ContentType) String() string {
	switch c {
	case ContentTypeXml:
		return "application/xml"
	case ContentTypeJson:
		return "application/json"
	case ContentTypeForm:
		return "application/x-www-form-urlencoded"
	default:
		return "text/plain"
	}
}

type HTTPMethod int

const (
	HTTPMethodGet HTTPMethod = iota
	HTTPMethodPost
	HTTPMethodPut
)

// TODO: other HTTP types
func (c HTTPMethod) String() string {
	switch c {
	case HTTPMethodPost:
		return "POST"
	case HTTPMethodPut:
		return "PUT"
	default:
		return "GET"
	}
}

type PrintOutputMode int

const (
	PrintRespHeaders PrintOutputMode = iota
	PrintRespBody
	PrintRespWhole
	PrintNone
)

type UserCredential struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type YamlConfig struct {
	ProxyURLString string         `yaml:"proxy_url"`
	ValidCred      UserCredential `yaml:"auth_valid"`
	InvalidCred    UserCredential `yaml:"auth_invalid"`
	BodyXml        string         `yaml:"body_xml"`
	BodyJson       string         `yaml:"body_json"`
	BodyxForm      string         `yaml:"body_xform"`
	Debug          bool           `yaml:"debug"`
	HTTPServerPort int            `yaml:"http_server_port"`
}

type FileRow struct {
	Row  int
	Text string
}

type Job func()

// TODO: not connected to the rest of the code yet
var RequestPerSecond = 5
