.PHONY: gremlin
gremlin:
	@docker-compose up graph
console:
	@docker-compose run console
gen:
	@go generate ./ent