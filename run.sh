#!/bin/bash

# 要运行的Go程序的相对文件路径
GO_PROGRAM_PATH="main.go"

# 通过ps命令和grep查找正在运行的Go程序的进程ID（PID）
PID=$(lsof -ti :9876)
echo "PID: $PID"
# 如果找到正在运行的进程，则终止它
if [ -n "$PID" ]; then
  echo "Stopping existing process with PID: $PID"
  kill -9 $PID
else
  echo "No existing process found"
fi

# 使用nohup启动新的Go程序
echo "Starting new process"
nohup go run $GO_PROGRAM_PATH > output.log 2>&1 &

echo "Process started"

