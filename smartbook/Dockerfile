# 基础镜像:fedora
FROM fedora

# 拷贝smartbook可执行文件（linux环境下的）到/app文件夹内，其实app文件夹不写死，也可以放在/user/local文件夹下的
COPY smartbook /app/smartbook
# 工作目录是/app
WORKDIR /app

# CMD是执行命令
# ENTRYPOINT是最佳实践
ENTRYPOINT ["/app/smartbook"]
