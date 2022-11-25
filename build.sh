set -e
set -x

GOOS=linux go build
imageName="user-srv:latest"
docker rmi -f "$imageName"
docker rmi -f registry.cn-hangzhou.aliyuncs.com/zj_image/"$imageName"

docker build . -t "$imageName"
docker tag "$imageName" registry.cn-hangzhou.aliyuncs.com/zj_image/"$imageName"
docker push registry.cn-hangzhou.aliyuncs.com/zj_image/"$imageName"