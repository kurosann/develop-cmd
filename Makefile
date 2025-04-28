# 定义变量
BINARY_NAME=devctl

# 默认目标
.PHONY: all
all: build

# 构建所有平台
.PHONY: build
build: build-linux build-darwin build-windows

# 构建 Linux 64位
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux/$(BINARY_NAME) ./cmd

# 构建 macOS
.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/darwin/$(BINARY_NAME) ./cmd
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o bin/darwin-arm/$(BINARY_NAME) ./cmd

# 构建 Windows
.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/windows/$(BINARY_NAME).exe ./cmd

# 开发环境构建（当前平台）
.PHONY: dev
dev:
	CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) ./cmd

install-darwin-arm:
	cp bin/darwin-arm/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# 构建不同平台的安装包
build-run:
	mkdir -p run
	mkdir -p bin/linux bin/darwin bin/darwin-arm

	# 构建 Linux 版本
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux/$(BINARY_NAME) ./cmd
	echo '#!/bin/bash' > run/$(BINARY_NAME)-linux.run
	echo 'if [ $$(id -u) != "0" ]; then' >> run/$(BINARY_NAME)-linux.run
	echo '    echo "Error: You must be root to run this script, please use sudo"' >> run/$(BINARY_NAME)-linux.run
	echo '    exit 1' >> run/$(BINARY_NAME)-linux.run
	echo 'fi' >> run/$(BINARY_NAME)-linux.run
	echo 'if [ "$$(uname)" != "Linux" ] || [ "$$(uname -m)" != "x86_64" ]; then' >> run/$(BINARY_NAME)-linux.run
	echo '    echo "Error: This installer is for Linux x86_64 only"' >> run/$(BINARY_NAME)-linux.run
	echo '    exit 1' >> run/$(BINARY_NAME)-linux.run
	echo 'fi' >> run/$(BINARY_NAME)-linux.run
	echo 'binStart=16' >> run/$(BINARY_NAME)-linux.run
	echo 'tail -n+$$binStart "$$0" > /usr/local/bin/$(BINARY_NAME)' >> run/$(BINARY_NAME)-linux.run
	echo 'chmod +x /usr/local/bin/$(BINARY_NAME)' >> run/$(BINARY_NAME)-linux.run
	echo 'echo "Installation completed. You can now use: $(BINARY_NAME)"' >> run/$(BINARY_NAME)-linux.run
	echo 'exit 0' >> run/$(BINARY_NAME)-linux.run
	echo '' >> run/$(BINARY_NAME)-linux.run
	cat bin/linux/$(BINARY_NAME) >> run/$(BINARY_NAME)-linux.run
	chmod +x run/$(BINARY_NAME)-linux.run

	# 构建 Darwin AMD64 版本
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/darwin/$(BINARY_NAME) ./cmd
	echo '#!/bin/bash' > run/$(BINARY_NAME)-darwin.run
	echo 'if [ $$(id -u) != "0" ]; then' >> run/$(BINARY_NAME)-darwin.run
	echo '    echo "Error: You must be root to run this script, please use sudo"' >> run/$(BINARY_NAME)-darwin.run
	echo '    exit 1' >> run/$(BINARY_NAME)-darwin.run
	echo 'fi' >> run/$(BINARY_NAME)-darwin.run
	echo 'if [ "$$(uname)" != "Darwin" ] || [ "$$(uname -m)" != "x86_64" ]; then' >> run/$(BINARY_NAME)-darwin.run
	echo '    echo "Error: This installer is for macOS Intel (x86_64) only"' >> run/$(BINARY_NAME)-darwin.run
	echo '    exit 1' >> run/$(BINARY_NAME)-darwin.run
	echo 'fi' >> run/$(BINARY_NAME)-darwin.run
	echo 'binStart=16' >> run/$(BINARY_NAME)-darwin.run
	echo 'tail -n+$$binStart "$$0" > /usr/local/bin/$(BINARY_NAME)' >> run/$(BINARY_NAME)-darwin.run
	echo 'chmod +x /usr/local/bin/$(BINARY_NAME)' >> run/$(BINARY_NAME)-darwin.run
	echo 'echo "Installation completed. You can now use: $(BINARY_NAME)"' >> run/$(BINARY_NAME)-darwin.run
	echo 'exit 0' >> run/$(BINARY_NAME)-darwin.run
	echo '' >> run/$(BINARY_NAME)-darwin.run
	cat bin/darwin/$(BINARY_NAME) >> run/$(BINARY_NAME)-darwin.run
	chmod +x run/$(BINARY_NAME)-darwin.run

	# 构建 Darwin ARM64 版本
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o bin/darwin-arm/$(BINARY_NAME) ./cmd
	echo '#!/bin/bash' > run/$(BINARY_NAME)-darwin-arm.run
	echo 'if [ $$(id -u) != "0" ]; then' >> run/$(BINARY_NAME)-darwin-arm.run
	echo '    echo "Error: You must be root to run this script, please use sudo"' >> run/$(BINARY_NAME)-darwin-arm.run
	echo '    exit 1' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'fi' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'if [ "$$(uname)" != "Darwin" ] || [ "$$(uname -m)" != "arm64" ]; then' >> run/$(BINARY_NAME)-darwin-arm.run
	echo '    echo "Error: This installer is for macOS Apple Silicon (M1/M2) only"' >> run/$(BINARY_NAME)-darwin-arm.run
	echo '    exit 1' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'fi' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'binStart=16' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'tail -n+$$binStart "$$0" > /usr/local/bin/$(BINARY_NAME)' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'chmod +x /usr/local/bin/$(BINARY_NAME)' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'echo "Installation completed. You can now use: $(BINARY_NAME)"' >> run/$(BINARY_NAME)-darwin-arm.run
	echo 'exit 0' >> run/$(BINARY_NAME)-darwin-arm.run
	echo '' >> run/$(BINARY_NAME)-darwin-arm.run
	cat bin/darwin-arm/$(BINARY_NAME) >> run/$(BINARY_NAME)-darwin-arm.run
	chmod +x run/$(BINARY_NAME)-darwin-arm.run

	rm -rf bin/linux/$(BINARY_NAME)
	rm -rf bin/darwin/$(BINARY_NAME)
	rm -rf bin/darwin-arm/$(BINARY_NAME)

# 清理构建文件
.PHONY: clean

clean:
	rm -rf run/
	rm -rf bin/
