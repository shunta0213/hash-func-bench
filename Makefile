.PHONY: bench
bench:
	@go test -bench=. -benchtime=0.1s > data.txt

.PHONY: plot
plot:
	@python3 graph/graph.py