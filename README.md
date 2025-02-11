# protoc-gen-go-oneof-api-adapter

When building an HTTP server based on protobuf message interactions, we needed a code generation tool to implement interface dispatch and avoid writing repetitive switch case logic by hand. Therefore, this protoc compilation plugin was developed.

See [example](./example/main.go) for complete code.

**Proto:**

```proto3
message Request {
  RequestHead head = 1;

  oneof body {
    LoginWithOpenIDRequest login_with_openid = 2;
    LoginWithAccountRequest login_with_account = 3;
  }
}

message Response {
  ResponseHead head = 1;

  oneof body {
    LoginWithOpenIDResponse login_with_openid = 2;
    LoginWithAccountResponse login_with_account = 3;
  }
}

...
```

**Install:**

```sh
go install github.com/panjiang/protoc-gen-go-oneof-api-adapter@latest
```

**Compile:**

```sh
protoc -I ./example/api \
  --go_out ./example/api --go_opt=paths=source_relative \
  --go-oneof-api-adapter_out ./example/api --go-oneof-api-adapter_opt=paths=source_relative,request=Request/body,response=Response/body \
  example/api/product/app/v1/v1.proto
```

**Options:**

- `request` - The oneof path of request messages attachment at.
- `response` - The oneof path of response messages attachment at.
