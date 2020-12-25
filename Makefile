ARGS = $(filter-out $@,$(MAKECMDGOALS))
%:
	@:
run:
	go run cmd/temp_email/main.go $(ARGS)