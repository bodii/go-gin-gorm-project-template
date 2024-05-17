#!/bin/bash

# 命令行帮助文档
usage(){
cat <<'EOF'
help doc

Usage:
  ./init_project.sh -n example_project

Options:
  -n|--name     input project name.
  -h|--help     show this message.
EOF
}

PROJECT_NAME=""

# 参数解析
function doParse() {
    while [ True ]; do
        if [ "$1" = "--name" -o "$1" = "-n" ]; then
            PROJECT_NAME=$2
            if [ -z $PROJECT_NAME  ]; then 
                usage
                exit 1
            else 
                break
            fi
        elif [ "$1" = "--help" -o "$1" = "-h" ]; then
            usage
            exit 1
        else 
            usage
            exit 1
        fi
    done
}





function projectInit() {
    find ./ -type f -name "*.*" -exec sed -i "s/\"template-project-name/\"$PROJECT_NAME/g" {} +
    # 创建项目
    go mod init $PROJECT_NAME
    # 下载插件
    go mod tidy

    # 执行默认env文件复制
    cp .env.example .env

    # 执行默认config项文件复制
    initConifgFile
}

function initConifgFile() {
    configPath="internal/config/"
    filepaths=`ls $configPath*default.toml`

    # echo "${filepaths[*]}"

    for default_filepath in $filepaths
    do
        default_filename=$(basename $default_filepath)
        # echo $default_filename
        filename=`echo "${default_filename%%_default.toml*}"`
        # echo $filename
        # echo "${configPath}${filename}.toml"
        cat $default_filepath > "${configPath}${filename}.toml"
    done
}

function createDir() {
    mkdir -p files/upload files/assets logs
}

# 执行最终方法
function run() {
    doParse "$@"
    echo "Generating $PROJECT_NAME ..."

    projectInit

    createDir
}

# 执行
run "$@"