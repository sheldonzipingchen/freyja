BINARY="freyja"
VERSION="0.0.1"

default: help

help: ## 显示帮助信息
	@echo 'usage: make [targets]'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

build: ## 构建应用程序
	go build -o bin/${BINARY}-${VERSION}-local freyja/lcmd/freyja
	@echo '本地应用程序构建完毕' 

build-linux: ## 构建Linux应用程序
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${BINARY}-${VERSION}-linux-amd64 freyja/cmd/freyja
	@echo 'linux应用程序构建完毕'

build-mac: ## 构建mac应用程序
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY}-${VERSION}-darwin-amd64 freyja/cmd/freyja
	@echo 'mac应用程序构建完毕'

build-windows: ## 构建windows应用程序
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${BINARY}-${VERSION}-windows-amd64.exe freyja/cmd/freyja
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${BINARY}-api-${VERSION}-windows-amd64.exe freyja/cmd/api
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${BINARY}-bot-${VERSION}-windows-amd64.exe freyja/cmd/bot
	@echo 'windows应用程序构建完毕'

clean: ## 清除构建文件
	rm bin/*
	@echo '过去构建文件清理完毕'

doc: ## 生成 godocs 文档并开启本地文档服务
	godoc -http=:8085 -index

db-migrate: ## 运行数据库迁移命令
	@echo '还没实现'

swagger-ui: ## 运行本地 swagger-ui
	@echo '还没实现'

