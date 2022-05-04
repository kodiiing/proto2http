package target

type ITarget interface {
	Generate(data Proto) ([]byte, error)
	FileExtension() string
}

type Proto struct {
	Name     string
	Comment  string
	BaseUrl  string
	Services []Service
	Enums    []Enum
	Messages []Message
}

type Service struct {
	Name    string
	Comment string
	RPCs    []RPC
}

type RPC struct {
	Name     string
	Comment  string
	Request  Message
	Response Message
}

type Message struct {
	Name    string
	Comment string
	Fields  []Field
}

type Field struct {
	Name     string
	Comment  string
	Type     string
	Sequence int16
	Repeated bool
	Optional bool
	Required bool
}

type Enum struct {
	Name    string
	Comment string
	Values  []EnumValue
}

type EnumValue struct {
	Key     string
	Comment string
	Integer int16
}
