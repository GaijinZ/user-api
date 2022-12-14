check-swagger:
	where swagger || go get github.com/go-swagger/go-swagger/cmd/swagger

swagger: check-swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml
	