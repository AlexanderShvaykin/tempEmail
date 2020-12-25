ARGS = $(filter-out $@,$(MAKECMDGOALS))
%:
	@:
run:
	go run cmd/temp_email/main.go $(ARGS)

build:
	 go build -o $(GOPATH)/bin/temp_email cmd/temp_email/main.go