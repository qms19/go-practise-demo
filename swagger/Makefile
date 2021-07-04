.PHONY: swagger.run
swagger.run:
	@echo "Generating swagger API docs"
	@swagger generate spec --scan-models -o swagger.yaml


.PHONY: swagger.serv
swagger.serv:
	@echo "Start swagger server"
	@swagger serve --no-open -F=swagger --port 36666 swagger.yaml