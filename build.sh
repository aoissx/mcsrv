#!/bin/bash

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
APP_NAME="mcsrv"
DIST_DIR="$PROJECT_DIR/dist"

# プロジェクトのルートディレクトリに移動する
cd "$PROJECT_DIR"

# フラグを解析する
parse_flags() {
  while getopts ":cfd:rah" opt; do
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
      r)
        # -rフラグが指定された場合、distフォルダに移動して成果物を実行する
        RUN_FLAG=true
        ;;
      h)
        # -hフラグが指定された場合、ヘルプを表示する
        show_help
        ;;
      \?)
        # 不正なフラグが指定された場合、エラーメッセージを表示する
        echo "Invalid option: -$OPTARG" >&2
        exit 1
        ;;
    esac
  done
}

# ヘルプを表示する
show_help() {
  echo "Usage: build.sh [-c] [-f] [-d] [-r] [-a arg1 arg2 ...] [-h]"
  echo "  -c: Clear the dist directory before building"
  echo "  -f: Force a rebuild even if a binary already exists"
  echo "  -d: Build a debug version of the binary"
  echo "  -r: Build and run the binary from the dist directory"
  echo "  -a: Build and run the binary from the dist directory with arguments"
  echo "  -h: Show this help message"
  exit 0
}

# ビルドする
build() {
  if [ "$FORCE_BUILD" = true ]; then
    # -fフラグが指定された場合、強制的にビルドする
    go build -o "$DIST_DIR/$APP_NAME"
    echo "Build successful. Output: $DIST_DIR/$APP_NAME"
  else
    # -fフラグが指定されていない場合、ビルド済みのバイナリがある場合はビルドしない
    if [ ! -f "$DIST_DIR/$APP_NAME" ] || [ "$DEBUG_BUILD" = true ]; then
      go build -o "$DIST_DIR/$APP_NAME"
        echo "Build successful. Output: $DIST_DIR/$APP_NAME"
    fi
  fi
}

# フラグが指定されていない場合、ビルドした成果物をdistフォルダに出力する
output_build() {
  if [ "$OPTIND" = 1 ]; then
    echo "Build successful. Output: $DIST_DIR/$APP_NAME"
  fi
}

# 実行する
run() {
    # -rフラグが指定された場合、distフォルダに移動して成果物を実行する
    if [ "$RUN_FLAG" = true ]; then
        cd "$DIST_DIR"
        ./"$APP_NAME" "${@:OPTIND}"
        exit 0
    fi
}

# フラグを解析する
parse_flags "$@"

# ビルドする
build

# 実行する
run "$@"