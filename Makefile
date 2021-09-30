ifeq ($(OS),Windows_NT)
	main := .\main.go
else
	main := main.go
endif

run:
	--exec go run $(main)

run_watch:
	reflex -r '\.go' -s -- sh -c "go run $(main)"

graph_gen:
	go run github.com/99designs/gqlgen generate