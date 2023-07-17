.PHONY: proto
proto:
	cd proto && buf generate
	cd proto && buf generate --template buf.gen.openapi-operatorservice.yaml --path necode/operatorservice
	cd proto && buf generate --template buf.gen.openapi-workflowservice.yaml --path necode/workflowservice