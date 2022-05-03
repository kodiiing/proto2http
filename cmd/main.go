package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path"
	"proto2http/handlers"
	"proto2http/target"
	"proto2http/target/browser_ts"

	"github.com/emicklei/proto"
)

type Dependency struct {
	Handlers        *handlers.Dependency
	TargetGenerator target.ITarget
}

func main() {
	var protoPath string
	flag.StringVar(&protoPath, "path", "", "path to proto file")

	var outputDirectory string
	flag.StringVar(&outputDirectory, "output", "", "output directory")

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbose output")

	var languageTarget string
	flag.StringVar(&languageTarget, "target", "", "target language")

	var baseUrl string
	flag.StringVar(&baseUrl, "baseurl", "", "http endpoint base url")

	flag.Parse()

	// Validate targets
	availableTargets := []string{"browser-js", "browser-ts", "go"}
	var targetExist bool
	for _, t := range availableTargets {
		if languageTarget == t {
			targetExist = true
		}
	}

	logger := log.New(os.Stderr, "[proto2http] ", 0)

	if !targetExist {
		logger.Printf("target does not exists")
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
	)

	deps.Handlers.Output.BaseUrl = baseUrl

	logger.Printf("%+v", deps.Handlers.Output)

	err = file.Close()
	if err != nil && !errors.Is(err, os.ErrClosed) {
		log.Printf("error closing file: %v", err)
		os.Exit(11)
	}

	switch languageTarget {
	case "browser-ts":
		deps.TargetGenerator = browser_ts.New()
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
