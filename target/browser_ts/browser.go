// browser provides
package browser_ts

import (
	"bytes"
	"proto2http/target"
	"strconv"
	"strings"
)

type Browser struct{}

func New() *Browser {
	return &Browser{}
}

func (Browser) FileExtension() string {
	return "ts"
}

func (Browser) Generate(data target.Proto) ([]byte, error) {
	var indent string = "    "
	var writer strings.Builder
	var typesWriter strings.Builder
	var typesMap = make(typeMap)

	for _, srv := range data.Services {
		if len(srv.Comment) > 0 {
			var comments = strings.Split(srv.Comment, "\n")
			writer.WriteString("/**\n")
			for _, c := range comments {
				writer.WriteString(" * " + strings.TrimSpace(c))
				writer.WriteString("\n")
			}
			writer.WriteString(" */\n")
		}

		writer.WriteString("export class ")
		writer.WriteString(srv.Name)
		writer.WriteString("Client")
		writer.WriteString(" {\n")

		// Constructor
		writer.WriteString(indent + "_baseUrl: string\n")
		writer.WriteString(indent + "constructor(baseUrl?: string) {\n")
		writer.WriteString(indent + indent + "if (baseUrl === \"\" || baseUrl == null) {\n")
		writer.WriteString(indent + indent + indent + "this._baseUrl = " + strconv.Quote(data.BaseUrl) + ";\n")
		writer.WriteString(indent + indent + "} else {\n")
		writer.WriteString(indent + indent + indent + "this._baseUrl = baseUrl;\n")
		writer.WriteString(indent + indent + "}\n")
		writer.WriteString(indent + "}\n\n")

		for _, rpc := range srv.RPCs {
			// Do things with the type
			if !typesMap.Exists(rpc.Request.Name) {
				if len(rpc.Request.Comment) > 0 {
					typesWriter.WriteString("/**\n")

					var comments = strings.Split(rpc.Request.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}

						typesWriter.WriteString(" * " + trimmed)
						typesWriter.WriteString("\n")
					}

					typesWriter.WriteString(" */\n")
				}

				typesWriter.WriteString("type " + rpc.Request.Name + " = {\n")

				for _, t := range rpc.Request.Fields {
					if len(t.Comment) > 0 {
						typesWriter.WriteString(indent + "/**\n")

						var comments = strings.Split(t.Comment, "\n")
						for _, c := range comments {
							trimmed := strings.TrimSpace(c)
							if trimmed == "" {
								continue
							}
							typesWriter.WriteString(indent)
							typesWriter.WriteString(" * " + trimmed)
							typesWriter.WriteString("\n")
						}

						typesWriter.WriteString(indent + " */\n")
					}
					typesWriter.WriteString(indent)
					typesWriter.WriteString(t.Name)

					if t.Optional {
						typesWriter.WriteString("?")
					}

					typesWriter.WriteString(": ")
					typesWriter.WriteString(typeToTypescript(t.Type))
					if t.Repeated {
						typesWriter.WriteString("[]")
					}
					typesWriter.WriteString("\n")
				}

				typesWriter.WriteString("}\n\n")

				typesMap.Add(rpc.Request.Name)
			}

			if !typesMap.Exists(rpc.Response.Name) {
				if len(rpc.Response.Comment) > 0 {
					typesWriter.WriteString("/**\n")

					var comments = strings.Split(rpc.Response.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}
						typesWriter.WriteString(" * " + trimmed)
						typesWriter.WriteString("\n")
					}

					typesWriter.WriteString(" */\n")
				}

				typesWriter.WriteString("type " + rpc.Response.Name + " = {\n")

				for _, t := range rpc.Response.Fields {
					if len(t.Comment) > 0 {
						typesWriter.WriteString(indent + "/**\n")

						var comments = strings.Split(t.Comment, "\n")
						for _, c := range comments {
							trimmed := strings.TrimSpace(c)
							if trimmed == "" {
								continue
							}
							typesWriter.WriteString(indent)
							typesWriter.WriteString(" * " + trimmed)
							typesWriter.WriteString("\n")
						}

						typesWriter.WriteString(indent + " */\n")
					}
					typesWriter.WriteString(indent)
					typesWriter.WriteString(t.Name)

					if t.Optional {
						typesWriter.WriteString("?")
					}

					typesWriter.WriteString(": ")
					typesWriter.WriteString(typeToTypescript(t.Type))

					if t.Repeated {
						typesWriter.WriteString("[]")
					}

					typesWriter.WriteString("\n")
				}

				typesWriter.WriteString("}\n\n")

				typesMap.Add(rpc.Response.Name)
			}

			var comments = strings.Split(rpc.Comment, "\n")
			writer.WriteString(indent + "/**\n")
			for _, c := range comments {
				writer.WriteString(indent + " ")
				writer.WriteString("* " + strings.TrimSpace(c))
				writer.WriteString("\n")
			}
			writer.WriteString(indent + " */\n")

			writer.WriteString(indent + "public async ")
			writer.WriteString(rpc.Name)
			writer.WriteString("(input: " + rpc.Request.Name)
			writer.WriteString("): Promise<" + rpc.Response.Name + "> {\n")

			writer.WriteString(indent + indent + "const request = await fetch(\n")
			writer.WriteString(indent + indent + indent + "new URL(" + strconv.Quote(rpc.Name) + ", this._baseUrl).toString(),\n")
			writer.WriteString(indent + indent + indent + "{\n")
			writer.WriteString(indent + indent + indent + indent + "method: \"POST\",\n")
			writer.WriteString(indent + indent + indent + indent + "headers: {\n")
			writer.WriteString(indent + indent + indent + indent + indent + "\"Content-Type\": \"application/json\",\n")
			writer.WriteString(indent + indent + indent + indent + indent + "\"Accept\": \"application/json\"\n")
			writer.WriteString(indent + indent + indent + indent + "},\n")
			writer.WriteString(indent + indent + indent + indent + "body: JSON.stringify(input),\n")
			writer.WriteString(indent + indent + indent + "}\n")
			writer.WriteString(indent + indent + ");\n\n")
			writer.WriteString(indent + indent + "const body = await request.json();\n")
			// TODO: should this be more explicit?
			writer.WriteString(indent + indent + "return body;\n")

			writer.WriteString(indent + "}\n\n")
		}

		writer.WriteString("}")
	}

	for _, msg := range data.Messages {
		if !typesMap.Exists(msg.Name) {
			if len(msg.Comment) > 0 {
				typesWriter.WriteString("/**\n")

				var comments = strings.Split(msg.Comment, "\n")
				for _, c := range comments {
					trimmed := strings.TrimSpace(c)
					if trimmed == "" {
						continue
					}

					typesWriter.WriteString(" * " + trimmed)
					typesWriter.WriteString("\n")
				}

				typesWriter.WriteString(" */\n")
			}

			typesWriter.WriteString("type " + msg.Name + " = {\n")

			for _, t := range msg.Fields {
				if len(t.Comment) > 0 {
					typesWriter.WriteString(indent + "/**\n")

					var comments = strings.Split(t.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}
						typesWriter.WriteString(indent)
						typesWriter.WriteString(" * " + trimmed)
						typesWriter.WriteString("\n")
					}

					typesWriter.WriteString(indent + " */\n")
				}
				typesWriter.WriteString(indent)
				typesWriter.WriteString(t.Name)

				if t.Optional {
					typesWriter.WriteString("?")
				}

				typesWriter.WriteString(": ")
				typesWriter.WriteString(typeToTypescript(t.Type))
				if t.Repeated {
					typesWriter.WriteString("[]")
				}
				typesWriter.WriteString("\n")
			}

			typesWriter.WriteString("}\n\n")

			typesMap.Add(msg.Name)
		}
	}

	for _, enum := range data.Enums {
		if !typesMap.Exists(enum.Name) {
			if len(enum.Comment) > 0 {
				typesWriter.WriteString("/**\n")

				var comments = strings.Split(enum.Comment, "\n")
				for _, c := range comments {
					trimmed := strings.TrimSpace(c)
					if trimmed == "" {
						continue
					}

					typesWriter.WriteString(" * " + trimmed)
					typesWriter.WriteString("\n")
				}

				typesWriter.WriteString(" */\n")
			}

			typesWriter.WriteString("enum " + enum.Name + " {\n")

			for i, t := range enum.Values {
				if len(t.Comment) > 0 {
					typesWriter.WriteString(indent + "/**\n")

					var comments = strings.Split(t.Comment, "\n")
					for _, c := range comments {
						trimmed := strings.TrimSpace(c)
						if trimmed == "" {
							continue
						}
						typesWriter.WriteString(indent)
						typesWriter.WriteString(" * " + trimmed)
						typesWriter.WriteString("\n")
					}

					typesWriter.WriteString(indent + " */\n")
				}
				typesWriter.WriteString(indent)
				typesWriter.WriteString(t.Key)

				typesWriter.WriteString(" = ")
				typesWriter.WriteString(strconv.Itoa(int(t.Integer)))
				if i != len(enum.Values)-1 {
					typesWriter.WriteString(",")
				}
				typesWriter.WriteString("\n")
			}

			typesWriter.WriteString("}\n\n")

			typesMap.Add(enum.Name)
		}
	}

	var finalWriter bytes.Buffer
	finalWriter.WriteString(typesWriter.String())
	finalWriter.WriteString(writer.String())

	return finalWriter.Bytes(), nil
}

type typeMap map[string]bool

func (t typeMap) Exists(key string) bool {
	for k := range t {
		if k == key {
			return true
		}
	}

	return false
}

func (t typeMap) Add(key string) {
	t[key] = true
}

func typeToTypescript(t string) string {
	switch t {
	case "float":
		fallthrough
	case "double":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "sint32":
		fallthrough
	case "sint64":
		fallthrough
	case "fixed32":
		fallthrough
	case "fixed64":
		fallthrough
	case "sfixed32":
		fallthrough
	case "sfixed64":
		return "number"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	default:
		return t
	}
}
