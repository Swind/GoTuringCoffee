DEST_IP = 192.168.1.77

GENERATE = internal/hardware/max31856/max31856.go\
		   internal/hardware/max31865/max31865.go\
		   internal/hardware/hardware.go

UTILS = max31856 max31865 db process

ARCH ?= amd64

.PHONY: generate clean

all: generate app utils

dep:
	dep ensure

app: generate
	GOOS=linux GOARCH=$(ARCH) go build -o GoTuringCoffee_$(ARCH)

utils: generate
	for util in $(UTILS); do \
		GOOS=linux GOARCH=$(ARCH) go build -o ./bin/$$util\_$(ARCH) ./utils/$$util.go ;\
	done

generate: $(GENERATE)
	for file in $(GENERATE) ; do \
		go generate $$file ; \
	done

clean:
	go clean
	rm GoTuringCoffee_*
	rm -rf ./bin

copy:
	sshpass scp config.yml bin/*_arm ./GoTuringCoffee_arm root@$(DEST_IP):/home/root
