.PHONY: build-cpu run-cpu run-cpu-multi \
        build-spin run-spin run-spin-multi \
        clean

build-cpu:
	go build -o cpu cpu.go

run-cpu:
	./cpu A

run-cpu-multi:
	./cpu A & ./cpu B & ./cpu C & ./cpu D &

clean:
	rm -f cpu
