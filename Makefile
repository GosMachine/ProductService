docker-push:
	@docker build -t gosmach1ne/gosboostproduct .
	@docker push gosmach1ne/gosboostproduct:latest