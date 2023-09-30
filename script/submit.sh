#!/bin/bash

# 引数のチェック
if [ "$#" -ne 2 ]; then
    echo "使用方法: $0 <コンテストID> <問題ID>"
    exit 1
fi

name=$1
question=$2

# コマンドの実行
cd "./src/${name}/${question}"
acc submit "main.go" # 提出