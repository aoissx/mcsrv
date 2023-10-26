#!/bin/bash

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
APP_NAME="mcsrv"
DIST_DIR="$PROJECT_DIR/dist"

# プロジェクトのルートディレクトリに移動する
cd "$PROJECT_DIR"

# フラグを解析する
while getopts ":cfd" opt; do
  case $opt in
    c)
      # -cフラグが指定された場合、distフォルダ内のファイルを全て削除する
      rm -rf "$DIST_DIR"/*
      ;;
    f)
      # -fフラグが指定された場合、ビルド時に強制的に上書きする
      FORCE_BUILD=true
      ;;
    d)
      # -dフラグが指定された場合、デバッグビルドを行う
      DEBUG_BUILD=true
      ;;
    \?)
      # 不正なフラグが指定された場合、エラーメッセージを表示する
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

# ビルドする
if [ "$FORCE_BUILD" = true ]; then
  # -fフラグが指定された場合、強制的にビルドする
  go build -o "$DIST_DIR/$APP_NAME"
else
  # -fフラグが指定されていない場合、ビルド済みのバイナリがある場合はビルドしない
  if [ ! -f "$DIST_DIR/$APP_NAME" ] || [ "$DEBUG_BUILD" = true ]; then
    go build -o "$DIST_DIR/$APP_NAME"
  fi
fi

# フラグが指定されていない場合、ビルドした成果物をdistフォルダに出力する
if [ "$OPTIND" = 1 ]; then
  echo "Build successful. Output: $DIST_DIR/$APP_NAME"
fi