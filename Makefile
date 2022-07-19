run:
	@go run -ldflags "-X 'main.version=0.01' -X 'main.commitHash=`git rev-parse HEAD`'" main.go