BINARY_NAME=ftbbackup
SERVICE_NAME=ftbbackup.service
INSTALL_PATH=/etc/ftbbackup
ENV_FILE=/etc/ftbbackup/.env

all: build

build:
	go build -o $(BINARY_NAME) main.go

install: build
	sudo mkdir -p $(INSTALL_PATH)
	sudo cp $(BINARY_NAME) $(INSTALL_PATH)/

	sudo cp .env.example $(ENV_FILE)

	sudo chmod 600 $(ENV_FILE)
	sudo chown root:root $(ENV_FILE)

	sudo ln -s $(INSTALL_PATH)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

uninstall:
	sudo rm -rf $(INSTALL_PATH)
	sudo rm -f /usr/local/bin/$(BINARY_NAME)

.PHONY: all build install uninstall
