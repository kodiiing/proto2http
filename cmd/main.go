package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path"
	"proto2http/handlers"
	"proto2http/target"
	browserts "proto2http/target/browser_ts"
	servergo "proto2http/target/server_go"
	servergochi "proto2http/target/server_go_chi"

	"github.com/emicklei/proto"
)

type Dependency struct {
	Handlers        *handlers.Dependency
	TargetGenerator target.ITarget
}

func main() {
	var protoPath string
	flag.StringVar(&protoPath, "path", "", "Path to the proto file (required)")

	var outputDirectory string
	flag.StringVar(&outputDirectory, "output", "", "Output directory (optional, default: current path)")

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Verbose output (optional, default: false)")

	var languageTarget string
	flag.StringVar(&languageTarget, "target", "", "Target language (required). Available values: browser-ts")

	var baseUrl string
	flag.StringVar(&baseUrl, "baseurl", "", "Base URL for HTTP endpoint (optional, default: empty string)")

	flag.Parse()

	// Validate targets
	availableTargets := []string{"browser-ts", "server-go", "server-go-chi"}
	var targetExist bool
	for _, t := range availableTargets {
		if languageTarget == t {
			targetExist = true
		}
	}

	logger := log.New(os.Stderr, "[proto2http] ", 0)

	if !targetExist {
		logger.Printf("target does not available")
		os.Exit(1)
		return
	}

	if protoPath == "" {
		logger.Printf("proto path cannot be empty")
		os.Exit(1)
		return
	}

	file, err := os.Open(protoPath)
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			logger.Fatalf("error during closing file: %v", err)
		}
	}()

	var targetOutput = &target.Proto{}
	var deps = &Dependency{
		Handlers: &handlers.Dependency{
			Verbose: verbose,
			Log:     logger,
			Output:  targetOutput,
		},
	}

	protoFile := proto.NewParser(file)

	parsedProtoFile, err := protoFile.Parse()
	if err != nil {
		logger.Printf("error parsing proto file: %v", err)
		os.Exit(10)
	}

	proto.Walk(
		parsedProtoFile,
		proto.WithPackage(deps.Handlers.PackageHandler),
		proto.WithService(deps.Handlers.ServiceHandler),
		proto.WithMessage(deps.Handlers.MessageHandler),
		proto.WithEnum(deps.Handlers.EnumHandler),
	)

	deps.Handlers.Output.BaseUrl = baseUrl

	err = file.Close()
	if err != nil && !errors.Is(err, os.ErrClosed) {
		log.Printf("error closing file: %v", err)
		os.Exit(11)
	}

	switch languageTarget {
	case "browser-ts":
		deps.TargetGenerator = browserts.New()
	case "server-go":
		deps.TargetGenerator = servergo.New()
	case "server-go-chi":
		deps.TargetGenerator = servergochi.New()
	default:
		logger.Printf("Unrecognizable target")
		os.Exit(12)
	}

	outputBytes, err := deps.TargetGenerator.Generate(*deps.Handlers.Output)
	if err != nil {
		logger.Printf("error generating to target: %v", err)
		os.Exit(13)
	}

	fileExtension := deps.TargetGenerator.FileExtension()
	createdFile, err := os.Create(path.Join(outputDirectory, deps.Handlers.Output.Name+"."+fileExtension))
	if err != nil {
		logger.Printf("error creating a file: %v", err)
		os.Exit(14)
	}
	defer createdFile.Close()

	_, err = createdFile.Write(outputBytes)
	if err != nil {
		logger.Printf("error writing to file: %v", err)
		os.Exit(15)
	}

	logger.Printf("file is written")
}
