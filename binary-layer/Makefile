clean:
	rm -Rf dist/extension*

build_x86: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/extensions/async-httpApi-extension main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/forwardEvent forward/main.go
	cd dist; zip extension.zip -r extensions/ forwardEvent

build_arm64: clean
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dist/extensions/async-httpApi-extension main.go
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dist/forwardEvent forward/main.go
	cd dist; zip extension.zip -r extensions/ forwardEvent

deploy_x86: build_x86
	aws lambda publish-layer-version --layer-name async-httpApi-extension-x86 --description "This layer will enable async behavior on an httpApi" --zip-file fileb://dist/extension.zip

deploy_arm64: build_arm64
	aws lambda publish-layer-version --layer-name async-httpApi-extension-arm --description "This layer will enable async behavior on an httpApi" --zip-file fileb://dist/extension.zip
