# Proto2http

Proto2http provides a code generation tool to convert your protocol buffers (.proto) files
into invokable HTTP request.

## Usage

```ssh
proto2http -path=your-file.proto -output=../generated_protos/ -baseurl=https://your-api-path.com -target=language-target
```

See `proto2http -help` for more detail.

## Samples

Converted files from gRPC's official sample [HelloWorld.proto](https://github.com/grpc/grpc/blob/be1a2ee5006aa903a46be662fe24b694384bccb3/examples/protos/helloworld.proto)
invoked with command:

```sh
proto2http -path=test.proto -target=browser-ts -baseurl=http://test-api.com/
```

```ts
/**
 * The request message containing the user's name.
 */
type HelloRequest = {
    name: string
}

/**
 * The response message containing the greetings
 */
type HelloReply = {
    message: string
}

/**
 * The greeting service definition.
 */
export class GreeterClient {
    _baseUrl: string
    constructor(baseUrl?: string) {
        if (baseUrl === "" || baseUrl == null) {
            this._baseUrl = "http://test-api.com/";
        } else {
            this._baseUrl = baseUrl;
        }
    }

    /**
     * Sends a greeting
     */
    public async SayHello(input: HelloRequest): Promise<HelloReply> {
        const request = await fetch(
            new URL("SayHello", this._baseUrl).toString(),
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json"
                },
                body: JSON.stringify(input),
            }
        );

        const body = await request.json();
        return body;
    }

}
```

For all client and server samples of [router_guide.proto](./handlers//fixtures/route_guide.proto), see [this gist](https://gist.github.com/aldy505/51493cd3026c6e2be563286e6319532a).

## Installation

### Precompiled binary

See [RELEASES](https://github.com/kodiiing/proto2http/releases) page.

Available systems: MacOS AMD64, Linux 386, Linux AMD64, Linux ARM, Linux ARM64, Windows 386, Windows AMD64, Windows ARM.

### Install it as a Go binary

```sh
go install github.com/kodiiing/proto2http@latest
```

### Build from source

You will need Go 1.17+

```sh
go build -o proto2http cmd/main.go
mv proto2http /usr/local/bin/proto2http
```

## License

[MIT](./LICENSE)
