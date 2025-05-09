# Todo API 部署說明

本文件說明如何將已打包的 `todo-api` Go 應用從開發機傳至 Linux 主機，並在 Docker 容器中啟動，掛載環境變數，進行服務驗證與容器／映像管理。

---

## 前置條件

- 本地已產出 `todo-api.tar` Docker 映像檔與 `.env` 配置檔。  
- 目標 Linux 主機已安裝 Docker 並啟動 Docker daemon。  
- 目標主機可透過 SSH 存取。

---

## 1. 檔案傳輸

1. 在本地（Windows PowerShell）執行：
   ```powershell
   scp "C:\Code\Demo\GolangDemo\todo-api\todo-api.tar" nero2@192.168.131.130:~/image/
   scp "C:\Code\Demo\GolangDemo\todo-api\.env"     nero2@192.168.131.130:~/image/
   ```
## 2. SSH 登入目標主機
   ```powershell
   ssh nero2@192.168.131.130
   ```
## 3. 確認檔案存在

   ```powershell
   ls ~/image/
   # todo-api.tar
   # .env
   ```

## 4. 解壓縮映像檔
   ```powershell
   # 進入映像檔所在目錄
   cd ~/image
   sudo docker load -i todo-api.tar
   # 檢查映像是否載入成功：
   sudo docker images
   # 應看到 todo-api:v3
   # 這裡的 v3 是版本號，可能會隨著更新而改變
   ```

## 5.啟動容器並掛載 .env
   ```powershell
   sudo docker run -d \
     --name todo-api \
     -p 8080:8080 \
     --env-file .env \
     todo-api:v3
   # 檢查容器狀態：
   sudo docker ps
   # 應看到 todo-api 容器正在運行
   # 這裡的 -d 是讓容器在背景運行
   # 這裡的 --env-file 是讓容器讀取 .env 檔案中的環境變數
   # 這裡的 -p 是讓容器的 8080 埠對應到主機的 8080 埠
   # 這裡的 todo-api:v3 是映像的名字和版本號
   # 這裡的 todo-api:v3 是映像的名字和版本號
   ```
## 6. 驗證服務
   ```powershell
   sudo docker logs -f todo-api
   # 應看到類似以下的日誌：
   # ✅ Server running at http://localhost:8080
   ```

## 7. 容器強制停止與刪除
   ```powershell
   # 1.取得容器對應的 PID：
   CONTAINER=容器名稱
   PID=$(sudo docker inspect --format '{{.State.Pid}}' $CONTAINER)
   echo "Container PID: $PID"

   # 2.強制停止容器：
   sudo kill -9 $PID

   # 3.刪除容器：
   sudo docker rm -f $CONTAINER
   ```
## 8. 一鍵批次清理（可選）
   ```powershell
   # 停止並移除所有容器
   sudo docker stop $(sudo docker ps -q)
   sudo docker rm   $(sudo docker ps -aq)

   # 清理未掛載卷
   sudo docker volume prune -f

   # 刪除指定舊映像
   sudo docker rmi skmapi:v1 test:v1 nero_image:latest 74d1dab78be9

   # 清理 dangling 映像
   sudo docker image prune -f
   ```



