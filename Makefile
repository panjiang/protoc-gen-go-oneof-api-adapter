test_example:
	go install
	protoc -I ./example/api \
		--go_out ./example/api --go_opt=paths=source_relative \
		--go-oneof-api-adapter_out ./example/api --go-oneof-api-adapter_opt=paths=source_relative,api=Api,request=Request/body,response=Response/body \
		example/api/product/app/v1/v1.proto