# go module 存放路徑
GO_MOD_PATH="$HOME/go/pkg/mod"

# 專案名稱
PROJECT_NAME=${PWD##*/}

echo "PROJECT_NAME=$PROJECT_NAME">.env
echo "GO_MOD_PATH"=$GO_MOD_PATH>>.env

# 啟動容器服務
docker-compose up -d
