

## 🔧 Makefile para automação

```makefile
.PHONY: run-dia1 run-dia2 test-all fmt clean

run-dia1:
	go run cmd/dia1-hello/main.go

run-dia2:
	go run cmd/dia2-adivinhacao/main.go

test-all:
	go test -v -race -cover ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

build-all:
	@for dir in cmd/*/; do \
		echo "Building $$dir..."; \
		go build -o bin/$$(basename $$dir) ./$$dir; \
	done

clean:
	rm -rf bin/
	go clean -cache

progress:
	@echo "=== Progresso do Curso ==="
	@echo "Dias completos:" $$(ls -d cmd/dia*/ 2>/dev/null | wc -l) "/ 15"
	@echo "Projetos:" $$(ls -d internal/*/ 2>/dev/null | wc -l)
