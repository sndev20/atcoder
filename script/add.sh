#!/bin/bash

# 引数のチェック
if [ "$#" -ne 1 ]; then
    echo "使用方法: $0 <コンテストID>"
    exit 1
fi

name=$1

# 問題の追加
cd "./src/${name}"
acc add