#!/bin/bash

# Docker 설치 여부 확인
if ! command -v docker &> /dev/null
then
    echo "Downloading Docker..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh

    # Docker 설치 성공 여부 확인
    if ! command -v docker &> /dev/null
    then 
        echo "Download Docker Successful"
        exit 1
    else 
        echo "Failed to Download Docker"
    fi
else 
    echo "Docker Already Exists"
fi

# MySQL 컨테이너 실행
echo "Starting MySQL Container..."
docker run --name first-mysql -e MYSQL_ROOT_PASSWORD=master -p 3306:3306 -d mysql

# 실행 결과 확인
if [ $? -eq 0 ]; then
    echo "MySQL Container Started Successfully"
else
    echo "Failed to Start MySQL Container"
fi

# 컨테이너가 실행될 동안 잠시 대기
sleep(10)

go run ../main.go