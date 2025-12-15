# 项目配置
BINARY_NAME=server
UI_DIR=ui
CMD_ENTRY=cmd/main.go

# 🟢 定义本地工具路径 (避免污染全局环境，解决找不到命令的问题)
LOCAL_BIN:=$(CURDIR)/bin
AIR:=$(LOCAL_BIN)/air
SWAG:=$(LOCAL_BIN)/swag

# ==============================================================================
# 开发命令 (Development)
# ==============================================================================

dev:
	@make -j2 dev-backend dev-ui

dev-backend:
	@echo "Starting Backend (Air)..."
	# 🟢 使用本地 bin 目录下的 air
	@$(AIR) -c .air.toml

dev-ui:
	@echo "Starting Frontend (Vite)..."
	@cd $(UI_DIR) && pnpm run dev

# ==============================================================================
# 构建命令 (Build)
# ==============================================================================

build: build-ui build-go
	@echo "✅ Build complete! Run ./$(BINARY_NAME) to start."

build-ui:
	@echo "Building Frontend..."
	@cd $(UI_DIR) && pnpm install && pnpm run build

build-go:
	@echo "Building Go Binary..."
	@go build -o $(BINARY_NAME) $(CMD_ENTRY)

# ==============================================================================
# 工具命令 (Utils)
# ==============================================================================

# 🛠 初始化项目 (关键修改)
init:
	@echo "Creating local bin directory..."
	@mkdir -p $(LOCAL_BIN)
	@echo "Installing tools to $(LOCAL_BIN)..."
	# 🟢 强制将工具安装到项目的 bin 目录下
	@GOBIN=$(LOCAL_BIN) go install github.com/air-verse/air@latest
	@GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Installing dependencies..."
	@go mod download
	@cd $(UI_DIR) && pnpm install
	@echo "✅ Initialization complete!"

docs:
	@$(SWAG) init -g $(CMD_ENTRY) --output ./api

clean:
	@rm -f $(BINARY_NAME)
	@rm -rf $(UI_DIR)/build
	@rm -rf $(UI_DIR)/node_modules
	@rm -rf $(LOCAL_BIN) # 清理工具