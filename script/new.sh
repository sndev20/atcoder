#!/bin/bash

# 引数のチェック
if [ "$#" -ne 1 ]; then
    echo "使用方法: $0 <コンテストID>"
    exit 1
fi

name=$1

# コマンドの実行
cd ./src
acc new "${name}"