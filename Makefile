swagger:
	swag init --dir ./cmd/ --output ./cmd/docs --pd --parseInternal --parseDepth 10  && rm cmd/docs/docs.go