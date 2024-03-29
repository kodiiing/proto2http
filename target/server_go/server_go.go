package servergo

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/kodiiing/proto2http/target"
)

type ServerGo struct{}

func New() *ServerGo {
	return &ServerGo{}
}

func (ServerGo) FileExtension() string {
	return "go"
}

func (ServerGo) Generate(data target.Proto) ([]byte, error) {
	var indent = "\t"

	var headersWriter strings.Builder
	var serverWriter strings.Builder
	var interfaceWriter strings.Builder
	var typesWriter strings.Builder
	var typesMap = make(typeMap)

	if len(data.Comment) > 0 {
		var comments = strings.Split(data.Comment, "\n")
		for _, c := range comments {
			headersWriter.WriteString("// " + strings.TrimSpace(c) + "\n")
		}
	}

	headersWriter.WriteString("package " + strings.ToLower(data.Name) + "\n\n")
	headersWriter.WriteString("import (\n")
	headersWriter.WriteString(indent + "\"context\"\n")
	headersWriter.WriteString(indent + "\"encoding/json\"\n")
	headersWriter.WriteString(indent + "\"net/http\"\n")
	headersWriter.WriteString(indent + "\"log\"\n")
	headersWriter.WriteString(")\n\n")

	for _, srv := range data.Services {
		if len(srv.Comment) > 0 {
			var comments = strings.Split(srv.Comment, "\n")
			for _, c := range comments {
				serverWriter.WriteString("// " + strings.TrimSpace(c))
				serverWriter.WriteString("\n")
			}
		}

		capitalizedServiceName := strings.ToUpper(string(srv.Name[0])) + string(srv.Name[1:])

		typesWriter.WriteString("type " + capitalizedServiceName + "Error struct {\n")
		typesWriter.WriteString(indent + "StatusCode int\n")
		typesWriter.WriteString(indent + "Error error\n")
		typesWriter.WriteString("}\n\n")

		serverWriter.WriteString("func New" + capitalizedServiceName + "Server(implementation " + capitalizedServiceName + "Server) http.Handler {\n")
		serverWriter.WriteString(indent + "mux := http.NewServeMux()\n")

		interfaceWriter.WriteString("type " + capitalizedServiceName + "Server interface {\n")

		for _, rpc := range srv.RPCs {
			// Do things with the type
			if !typesMap.Exists(rpc.Request.Name) {
				if len(rpc.Request.Comment) > 0 {
					var comments = strings.Split(rpc.Request.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}

						typesWriter.WriteString("// " + trimmed)
						typesWriter.WriteString("\n")
					}
				}

				capitalizedStructName := strings.ToUpper(string(rpc.Request.Name[0])) + string(rpc.Request.Name[1:])
				typesWriter.WriteString("type " + capitalizedStructName + " struct {\n")

				for _, t := range rpc.Request.Fields {
					if len(t.Comment) > 0 {
						var comments = strings.Split(t.Comment, "\n")
						for _, c := range comments {
							trimmed := strings.TrimSpace(c)
							if trimmed == "" {
								continue
							}
							typesWriter.WriteString(indent)
							typesWriter.WriteString("// " + trimmed)
							typesWriter.WriteString("\n")
						}
					}

					typesWriter.WriteString(indent)
					typesWriter.WriteString(convertToCamel(t.Name))

					typesWriter.WriteString(" ")
					if t.Repeated {
						typesWriter.WriteString("[]")
					}
					typesWriter.WriteString(typeToGo(t.Type))
					typesWriter.WriteString(" `json:\"" + t.Name + "\"`\n")
				}

				typesWriter.WriteString("}\n\n")

				typesMap.Add(rpc.Request.Name)
			}

			if !typesMap.Exists(rpc.Response.Name) {
				if len(rpc.Response.Comment) > 0 {
					var comments = strings.Split(rpc.Response.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}
						typesWriter.WriteString("// " + trimmed)
						typesWriter.WriteString("\n")
					}
				}

				capitalizedStructName := strings.ToUpper(string(rpc.Response.Name[0])) + string(rpc.Response.Name[1:])
				typesWriter.WriteString("type " + capitalizedStructName + " struct {\n")

				for _, t := range rpc.Response.Fields {
					if len(t.Comment) > 0 {
						var comments = strings.Split(t.Comment, "\n")
						for _, c := range comments {
							trimmed := strings.TrimSpace(c)
							if trimmed == "" {
								continue
							}
							typesWriter.WriteString(indent)
							typesWriter.WriteString("// " + trimmed)
							typesWriter.WriteString("\n")
						}
					}
					typesWriter.WriteString(indent)
					typesWriter.WriteString(convertToCamel(t.Name))

					typesWriter.WriteString(" ")
					if t.Repeated {
						typesWriter.WriteString("[]")
					}
					typesWriter.WriteString(typeToGo(t.Type))
					typesWriter.WriteString(" `json:\"" + t.Name + "\"`\n")
				}

				typesWriter.WriteString("}\n\n")

				typesMap.Add(rpc.Response.Name)
			}

			capitalizedRpcName := strings.ToUpper(string(rpc.Name[0])) + string(rpc.Name[1:])
			capitalizedRequestName := strings.ToUpper(string(rpc.Request.Name[0])) + string(rpc.Request.Name[1:])
			capitalizedResponseName := strings.ToUpper(string(rpc.Response.Name[0])) + string(rpc.Response.Name[1:])

			serverWriter.WriteString(indent + "mux.HandleFunc(\"/" + rpc.Name + "\", func(w http.ResponseWriter, r *http.Request) {\n")
			serverWriter.WriteString(indent + indent + "if r.Method != http.MethodPost {\n")
			serverWriter.WriteString(indent + indent + indent + "w.WriteHeader(http.StatusMethodNotAllowed)\n")
			serverWriter.WriteString(indent + indent + indent + "return\n")
			serverWriter.WriteString(indent + indent + "}\n")
			serverWriter.WriteString(indent + indent + "var req " + capitalizedRequestName + "\n")
			serverWriter.WriteString(indent + indent + "e := json.NewDecoder(r.Body).Decode(&req)\n")
			serverWriter.WriteString(indent + indent + "if e != nil {\n")
			serverWriter.WriteString(indent + indent + indent + "w.Header().Set(\"Content-Type\", \"application/json\")\n")
			serverWriter.WriteString(indent + indent + indent + "w.WriteHeader(400)\n")
			serverWriter.WriteString(indent + indent + indent + "e := json.NewEncoder(w).Encode(map[string]string{\n")
			serverWriter.WriteString(indent + indent + indent + indent + "\"message\": e.Error(),\n")
			serverWriter.WriteString(indent + indent + indent + "})\n")
			serverWriter.WriteString(indent + indent + indent + "if e != nil {\n")
			serverWriter.WriteString(indent + indent + indent + indent + "log.Printf(\"[" + capitalizedServiceName + " - " + capitalizedRpcName + "error] writing to response stream: %s\", e.Error())\n")
			serverWriter.WriteString(indent + indent + indent + "}\n")
			serverWriter.WriteString(indent + indent + indent + "return\n")
			serverWriter.WriteString(indent + indent + "}\n")
			serverWriter.WriteString(indent + indent + "resp, err := implementation." + capitalizedRpcName + "(r.Context(), &req)\n")
			serverWriter.WriteString(indent + indent + "if err != nil {\n")
			serverWriter.WriteString(indent + indent + indent + "w.Header().Set(\"Content-Type\", \"application/json\")\n")
			serverWriter.WriteString(indent + indent + indent + "w.WriteHeader(err.StatusCode)\n")
			serverWriter.WriteString(indent + indent + indent + "e := json.NewEncoder(w).Encode(map[string]string{\n")
			serverWriter.WriteString(indent + indent + indent + indent + "\"message\": err.Error.Error(),\n")
			serverWriter.WriteString(indent + indent + indent + "})\n")
			serverWriter.WriteString(indent + indent + indent + "if e != nil {\n")
			serverWriter.WriteString(indent + indent + indent + indent + "log.Printf(\"[" + capitalizedServiceName + " - " + capitalizedRpcName + "error] writing to response stream: %s\", e.Error())\n")
			serverWriter.WriteString(indent + indent + indent + "}\n")
			serverWriter.WriteString(indent + indent + indent + "return\n")
			serverWriter.WriteString(indent + indent + "}\n")
			serverWriter.WriteString(indent + indent + "w.Header().Set(\"Content-Type\", \"application/json\")\n")
			serverWriter.WriteString(indent + indent + "w.WriteHeader(http.StatusOK)\n")
			serverWriter.WriteString(indent + indent + "e = json.NewEncoder(w).Encode(resp)\n")
			serverWriter.WriteString(indent + indent + "if e != nil {\n")
			serverWriter.WriteString(indent + indent + indent + "log.Printf(\"[" + capitalizedServiceName + " - " + capitalizedRpcName + "error] writing to response stream: %s\", e.Error())\n")
			serverWriter.WriteString(indent + indent + "}\n")
			serverWriter.WriteString(indent + "})\n\n")

			var comments = strings.Split(rpc.Comment, "\n")
			for _, c := range comments {
				if strings.TrimSpace(c) == "" {
					continue
				}
				interfaceWriter.WriteString(indent)
				interfaceWriter.WriteString("// " + strings.TrimSpace(c))
				interfaceWriter.WriteString("\n")
			}
			interfaceWriter.WriteString(indent + capitalizedRpcName + "(ctx context.Context, req *" + capitalizedRequestName + ") (*" + capitalizedResponseName + ", *" + capitalizedServiceName + "Error)\n")
		}

		serverWriter.WriteString(indent + "return mux\n")
		serverWriter.WriteString("}\n")

		interfaceWriter.WriteString("}\n")
	}

	for _, msg := range data.Messages {
		if !typesMap.Exists(msg.Name) {
			if len(msg.Comment) > 0 {
				var comments = strings.Split(msg.Comment, "\n")
				for _, c := range comments {
					trimmed := strings.TrimSpace(c)
					if trimmed == "" {
						continue
					}

					typesWriter.WriteString("// " + trimmed)
					typesWriter.WriteString("\n")
				}
			}

			capitalizedStructName := strings.ToUpper(string(msg.Name[0])) + string(msg.Name[1:])
			typesWriter.WriteString("type " + capitalizedStructName + " struct {\n")

			for _, t := range msg.Fields {
				if len(t.Comment) > 0 {
					var comments = strings.Split(t.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}
						typesWriter.WriteString(indent)
						typesWriter.WriteString("// " + trimmed)
						typesWriter.WriteString("\n")
					}
				}
				typesWriter.WriteString(indent)
				typesWriter.WriteString(convertToCamel(t.Name))

				typesWriter.WriteString(" ")
				if t.Repeated {
					typesWriter.WriteString("[]")
				}
				typesWriter.WriteString(typeToGo(t.Type))
				typesWriter.WriteString(" `json:\"" + t.Name + "\"`\n")
			}

			typesWriter.WriteString("}\n\n")

			typesMap.Add(msg.Name)
		}
	}

	for _, enum := range data.Enums {
		if !typesMap.Exists(enum.Name) {
			if len(enum.Comment) > 0 {
				var comments = strings.Split(enum.Comment, "\n")
				for _, c := range comments {
					trimmed := strings.TrimSpace(c)
					if trimmed == "" {
						continue
					}

					typesWriter.WriteString("// " + trimmed)
					typesWriter.WriteString("\n")
				}
			}

			capitalizedEnumName := strings.ToUpper(string(enum.Name[0])) + string(enum.Name[1:])

			typesWriter.WriteString("type " + capitalizedEnumName + " uint32\n")
			typesWriter.WriteString("const (\n")
			for i, t := range enum.Values {
				if len(t.Comment) > 0 {
					var comments = strings.Split(t.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}
						typesWriter.WriteString(indent)
						typesWriter.WriteString("// " + trimmed)
						typesWriter.WriteString("\n")
					}
				}

				typesWriter.WriteString(indent + t.Key + " " + capitalizedEnumName + " = " + strconv.Itoa(i) + "\n")
			}

			typesWriter.WriteString(")\n\n")

			typesMap.Add(enum.Name)
		}
	}

	var finalWriter bytes.Buffer
	finalWriter.WriteString(headersWriter.String())
	finalWriter.WriteString(typesWriter.String())
	finalWriter.WriteString(interfaceWriter.String())
	finalWriter.WriteString(serverWriter.String())

	return finalWriter.Bytes(), nil
}

type typeMap map[string]bool

func (t typeMap) Exists(key string) bool {
	_, ok := t[key]
	return ok
}

func (t typeMap) Add(key string) {
	t[key] = true
}

func typeToGo(t string) string {
	switch t {
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "sint32":
		fallthrough
	case "sfixed32":
		fallthrough
	case "int32":
		return "int32"
	case "sint64":
		fallthrough
	case "sfixed64":
		fallthrough
	case "int64":
		return "int64"
	case "fixed32":
		fallthrough
	case "uint32":
		return "uint32"
	case "fixed64":
		fallthrough
	case "uint64":
		return "uint64"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "bytes":
		return "[]byte"
	default:
		return t
	}
}

func convertToCamel(input string) string {
	var splitted = strings.Split(input, "_")

	var veryClean []string
	for _, s := range splitted {
		veryClean = append(veryClean, strings.Split(s, "-")...)
	}

	var camel []string
	for _, s := range veryClean {
		camel = append(camel, strings.ToUpper(string(s[0]))+string(s[1:]))
	}

	return strings.Join(camel, "")
}
