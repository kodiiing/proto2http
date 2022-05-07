package browserts_test

import (
	"proto2http/target"
	browserts "proto2http/target/browser_ts"
	"testing"
)

func TestFileExtension(t *testing.T) {
	b := browserts.New()

	if b.FileExtension() != "ts" {
		t.Errorf("Expected file extension to be 'ts', got '%s'", b.FileExtension())
	}
}

func TestGenerate_RouteGuide(t *testing.T) {
	b := browserts.New()

	data := target.Proto{
		Name:    "routeguide",
		Comment: "",
		BaseUrl: "",
		Services: []target.Service{
			{
				Name:    "RouteGuide",
				Comment: `Interface exported by the server.`,
				RPCs: []target.RPC{
					{
						Name: "GetFeature",
						Comment: `A simple RPC.

						Obtains the feature at a given position.

						A feature with an empty name is returned if there's no feature at the given
						position.`,
						Request: target.Message{
							Name: "Point",
							Comment: `Points are represented as latitude-longitude pairs in the E7 representation
						(degrees multiplied by 10**7 and rounded to the nearest integer).
						Latitudes should be in the range +/- 90 degrees and longitude should be in
						the range +/- 180 degrees (inclusive).`,
							Fields: []target.Field{
								{
									Name:     "latitude",
									Comment:  "",
									Type:     "int32",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "longitude",
									Comment:  "",
									Type:     "int32",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
						Response: target.Message{
							Name: "Feature",
							Comment: `A feature names something at a given point.

							If a feature could not be named, the name is empty.`,
							Fields: []target.Field{
								{
									Name:     "name",
									Comment:  "The name of the feature.",
									Type:     "string",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "location",
									Comment:  "The point where the feature is detected.",
									Type:     "Point",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
					},
					{
						Name: "ListFeatures",
						Comment: `A server-to-client streaming RPC.

						Obtains the Features available within the given Rectangle.  Results are
						streamed rather than returned at once (e.g. in a response message with a
						repeated field), as the rectangle may cover a large area and contain a
						huge number of features.`,
						Request: target.Message{
							Name: "Rectangle",
							Comment: `A latitude-longitude rectangle, represented as two diagonally opposite
									points "lo" and "hi".`,
							Fields: []target.Field{
								{
									Name:     "lo",
									Comment:  "One corner of the rectangle.",
									Type:     "Point",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "hi",
									Comment:  "The other corner of the rectangle.",
									Type:     "Point",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
						Response: target.Message{
							Name: "Feature",
							Comment: `A feature names something at a given point.

							If a feature could not be named, the name is empty.`,
							Fields: []target.Field{
								{
									Name:     "name",
									Comment:  " The name of the feature.",
									Type:     "string",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "location",
									Comment:  "The point where the feature is detected.",
									Type:     "Point",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
					},
					{
						Name: "RecordRoute",
						Comment: `A client-to-server streaming RPC.

						Accepts a stream of Points on a route being traversed, returning a
						RouteSummary when traversal is completed.`,
						Request: target.Message{
							Name: "Point",
							Comment: `Points are represented as latitude-longitude pairs in the E7 representation
							(degrees multiplied by 10**7 and rounded to the nearest integer).
							Latitudes should be in the range +/- 90 degrees and longitude should be in
							the range +/- 180 degrees (inclusive).`,
							Fields: []target.Field{
								{
									Name:     "latitude",
									Comment:  "",
									Type:     "int32",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "longitude",
									Comment:  "",
									Type:     "int32",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
						Response: target.Message{
							Name: "RouteSummary",
							Comment: `A RouteSummary is received in response to a RecordRoute rpc.

							It contains the number of individual points received, the number of
							detected features, and the total distance covered as the cumulative sum of
							the distance between each point.`,
							Fields: []target.Field{
								{
									Name:     "point_count",
									Comment:  "The number of points received.",
									Type:     "int32",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "feature_count",
									Comment:  "The number of known features passed while traversing the route.",
									Type:     "int32",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "distance",
									Comment:  "The distance covered in metres.",
									Type:     "int32",
									Sequence: 3,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "elapsed_time",
									Comment:  "The duration of the traversal in seconds.",
									Type:     "int32",
									Sequence: 4,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
					},
					{
						Name: "RouteChat",
						Comment: `A Bidirectional streaming RPC.

							Accepts a stream of RouteNotes sent while a route is being traversed,
							while receiving other RouteNotes (e.g. from other users).`,
						Request: target.Message{
							Name:    "RouteNote",
							Comment: "A RouteNote is a message sent while at a given point.",
							Fields: []target.Field{
								{
									Name:     "location",
									Comment:  "The location from which the message is sent.",
									Type:     "Point",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "message",
									Comment:  "The message to be sent.",
									Type:     "string",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
						Response: target.Message{
							Name:    "RouteNote",
							Comment: "A RouteNote is a message sent while at a given point.",
							Fields: []target.Field{
								{
									Name:     "location",
									Comment:  "The location from which the message is sent.",
									Type:     "Point",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "message",
									Comment:  "The message to be sent.",
									Type:     "string",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
					},
				},
			},
		},
		Enums: []target.Enum{},
		Messages: []target.Message{
			{
				Name: "Point",
				Comment: `Points are represented as latitude-longitude pairs in the E7 representation
					(degrees multiplied by 10**7 and rounded to the nearest integer).
					Latitudes should be in the range +/- 90 degrees and longitude should be in
					the range +/- 180 degrees (inclusive).`,
				Fields: []target.Field{
					{
						Name:     "latitude",
						Comment:  "",
						Type:     "int32",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "longitude",
						Comment:  "",
						Type:     "int32",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name: "Rectangle",
				Comment: `A latitude-longitude rectangle, represented as two diagonally opposite
					points "lo" and "hi".`,
				Fields: []target.Field{
					{
						Name:     "lo",
						Comment:  "One corner of the rectangle.",
						Type:     "Point",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "hi",
						Comment:  "The other corner of the rectangle.",
						Type:     "Point",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name: "Feature",
				Comment: `A feature names something at a given point.

					If a feature could not be named, the name is empty.`,
				Fields: []target.Field{
					{
						Name:     "name",
						Comment:  "The name of the feature.",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "location",
						Comment:  "The point where the feature is detected.",
						Type:     "Point",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "RouteNote",
				Comment: "A RouteNote is a message sent while at a given point.",
				Fields: []target.Field{
					{
						Name:     "location",
						Comment:  "The location from which the message is sent.",
						Type:     "Point",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "message",
						Comment:  "The message to be sent.",
						Type:     "string",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name: "RouteSummary",
				Comment: `A RouteSummary is received in response to a RecordRoute rpc.

					It contains the number of individual points received, the number of
					detected features, and the total distance covered as the cumulative sum of
					the distance between each point.`,
				Fields: []target.Field{
					{
						Name:     "point_count",
						Comment:  "The number of points received.",
						Type:     "int32",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "feature_count",
						Comment:  "The number of known features passed while traversing the route.",
						Type:     "int32",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "distance",
						Comment:  "The distance covered in metres.",
						Type:     "int32",
						Sequence: 3,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "elapsed_time",
						Comment:  "The duration of the traversal in seconds.",
						Type:     "int32",
						Sequence: 4,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
		},
	}

	out, err := b.Generate(data)
	if err != nil {
		t.Errorf("Generate() failed: %v", err)
	}

	if string(out) == "" {
		t.Errorf("Generate() returned empty output")
	}
}

func TestGenerate_Rce(t *testing.T) {
	b := browserts.New()

	data := target.Proto{
		Name:    "rce",
		Comment: "",
		BaseUrl: "",
		Services: []target.Service{
			{
				Name:    "CodeExecutionEngineService",
				Comment: "",
				RPCs: []target.RPC{
					{
						Name:    "ListRuntimes",
						Comment: "",
						Request: target.Message{
							Name:    "EmptyRequest",
							Comment: "",
							Fields:  []target.Field{},
						},
						Response: target.Message{
							Name:    "Runtimes",
							Comment: "",
							Fields: []target.Field{
								{
									Name:     "runtime",
									Comment:  "",
									Type:     "Runtime",
									Sequence: 1,
									Repeated: true,
									Optional: false,
									Required: false,
								},
							},
						},
					},
					{
						Name:    "Execute",
						Comment: "",
						Request: target.Message{
							Name:    "CodeRequest",
							Comment: "",
							Fields: []target.Field{
								{
									Name:     "language",
									Comment:  "",
									Type:     "string",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "version",
									Comment:  "",
									Type:     "string",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "code",
									Comment:  "",
									Type:     "string",
									Sequence: 3,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "compile_timeout",
									Comment:  "",
									Type:     "int32",
									Sequence: 4,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "run_timeout",
									Comment:  "",
									Type:     "int32",
									Sequence: 5,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "memory_limit",
									Comment:  "",
									Type:     "int32",
									Sequence: 6,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
						Response: target.Message{
							Name:    "CodeResponse",
							Comment: "",
							Fields: []target.Field{
								{
									Name:     "language",
									Comment:  "",
									Type:     "string",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "version",
									Comment:  "",
									Type:     "string",
									Sequence: 2,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "compile",
									Comment:  "",
									Type:     "Output",
									Sequence: 3,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "runtime",
									Comment:  "",
									Type:     "Output",
									Sequence: 4,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
					},
					{
						Name:    "Ping",
						Comment: "",
						Request: target.Message{
							Name:    "EmptyRequest",
							Comment: "",
							Fields:  []target.Field{},
						},
						Response: target.Message{
							Name:    "PingResponse",
							Comment: "", Fields: []target.Field{
								{
									Name:     "message",
									Comment:  "",
									Type:     "string",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
							},
						},
					},
				},
			},
			{
				Name:    "LoggingService",
				Comment: "Logging service provides a way to log messages to the server.",
				RPCs: []target.RPC{
					{
						Name:    "SendLog",
						Comment: "",
						Request: target.Message{
							Name:    "SendLogRequest",
							Comment: "",
							Fields: []target.Field{
								{
									Name:     "message",
									Comment:  "",
									Type:     "string",
									Sequence: 1,
									Repeated: false,
									Optional: false,
									Required: false,
								},
								{
									Name:     "level",
									Comment:  "The log level, whether it be an error, warning, info, or debug message.",
									Type:     "LogLevel",
									Sequence: 2,
									Repeated: false,
									Optional: true,
									Required: false,
								},
								{
									Name:     "additional_data",
									Comment:  "Additional data to be logged.",
									Type:     "AdditionalData",
									Sequence: 3,
									Repeated: true,
									Optional: true,
									Required: false,
								},
							},
						},
						Response: target.Message{
							Name:    "SendLogResponse",
							Comment: "",
							Fields:  []target.Field{},
						},
					},
				},
			},
		},
		Enums: []target.Enum{
			{
				Name:    "LogLevel",
				Comment: "LogLevel provides an enum for handling logging levels",
				Values: []target.EnumValue{
					{
						Key:     "Error",
						Comment: "Means an error occured",
						Integer: 0,
					},
					{
						Key:     "Warning",
						Comment: "Means a warning occured",
						Integer: 1,
					},
					{
						Key:     "Info",
						Comment: "Means an informational message",
						Integer: 2,
					},
					{
						Key:     "Debug",
						Comment: "Means a debug message",
						Integer: 3,
					},
				},
			},
		},
		Messages: []target.Message{
			{
				Name:    "AdditionalData",
				Comment: "AdditionalData provides additional data for the logs",
				Fields: []target.Field{
					{
						Name:     "key",
						Comment:  "The key for the data",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "value",
						Comment:  "The value for the data",
						Type:     "string",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "EmptyRequest",
				Comment: "",
				Fields:  []target.Field{},
			},
			{
				Name:    "Runtimes",
				Comment: "",
				Fields: []target.Field{
					{
						Name:     "runtime",
						Comment:  "",
						Type:     "Runtime",
						Sequence: 1,
						Repeated: true,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "Runtime",
				Comment: "",
				Fields: []target.Field{
					{
						Name:     "language",
						Comment:  "",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "version",
						Comment:  "",
						Type:     "string",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "aliases",
						Comment:  "",
						Type:     "string",
						Sequence: 3,
						Repeated: true,
						Optional: false,
						Required: false,
					},
					{
						Name:     "compiled",
						Comment:  "",
						Type:     "bool",
						Sequence: 4,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "PingResponse",
				Comment: "",
				Fields: []target.Field{
					{
						Name:     "message",
						Comment:  "",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "CodeRequest",
				Comment: "",
				Fields: []target.Field{
					{
						Name:     "language",
						Comment:  "",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "version",
						Comment:  "",
						Type:     "string",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "code",
						Comment:  "",
						Type:     "string",
						Sequence: 3,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "compile_timeout",
						Comment:  "",
						Type:     "int32",
						Sequence: 4,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "run_timeout",
						Comment:  "",
						Type:     "int32",
						Sequence: 5,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "memory_limit",
						Comment:  "",
						Type:     "int32",
						Sequence: 6,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "CodeResponse",
				Comment: "",
				Fields: []target.Field{
					{
						Name:     "language",
						Comment:  "",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "version",
						Comment:  "",
						Type:     "string",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "compile",
						Comment:  "",
						Type:     "Output",
						Sequence: 3,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "runtime",
						Comment:  "",
						Type:     "Output",
						Sequence: 4,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
			{
				Name:    "Output",
				Comment: "",
				Fields: []target.Field{
					{
						Name:     "stdout",
						Comment:  "",
						Type:     "string",
						Sequence: 1,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "stderr",
						Comment:  "",
						Type:     "string",
						Sequence: 2,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "output",
						Comment:  "",
						Type:     "string",
						Sequence: 3,
						Repeated: false,
						Optional: false,
						Required: false,
					},
					{
						Name:     "exitCode",
						Comment:  "",
						Type:     "int32",
						Sequence: 4,
						Repeated: false,
						Optional: false,
						Required: false,
					},
				},
			},
		},
	}

	out, err := b.Generate(data)
	if err != nil {
		t.Errorf("Generate() failed: %v", err)
	}

	if string(out) == "" {
		t.Errorf("Generate() returned empty output")
	}
}
