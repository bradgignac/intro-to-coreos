default: application

application:
	docker build -t coreos-intro-application application

.PHONY: application
