# OCR识别Demo

## 环境

    System: Ubuntu 20.04 LTS
    Go version: go 1.14/4
    tesseract: 4.1.1
    "github.com/otiai10/gosseract/v2"
    Mysql server 8.0

## 准备工作
准备了OCR基础库, 中文, 英文, 日语 (当然也可以自己添加别的啦~)
在Ubuntu 18.04, 20.04上均可以正常运行

    sudo apt install tesseract-ocr
    sudo apt install libtesseract-dev
    sudo apt install tesseract-ocr-chi-sim tesseract-ocr-chi-sim-vert tesseract-ocr-eng tesseract-ocr-jpn tesseract-ocr-jpn-vert
    
## 实现功能

1. 支持多语言, 比如: 有一张图片, 里面同时有中文, 英文, 日文, 可以一起检索, 一起输出
(使用rune, 也就是int32
, 也是mysql中的utf8mb4)
2. 用户在识别的时候自己输入一个标识符(userId), 日后可以检索自己的历史记录(有可能和别人的标识符重合, 从而检索出别人的结果, 后续可以做成登陆的形式来避免此类情况)
3. 用户在使用服务时会提示是否将自己的图片资源保留在服务器上, 图片使用ksuid命名, 系统匿名使用. 可以供系统对比图片识别结果和实际结果的差异, 提升服务质量, 当然, 这个是用户可选的
4. 用户可以指定tesseract支持的其他语言, 比如: 俄语, 法语etc
5. 使用fileType来识别文件类型, 不依赖拓展名
6. 剔除无用字符, 例如: \n, \r, [Space]

## 接口设计

1. 上传图片并获取识别结果

    Method: POST
    
    URL: /ocr
    
    Content-Type: multipart/form-data
    ```
    imgFile: [file] binary // 非空
    userId: userID string // 可选, 如果用户需要保存自己之前的记录, 就填写一下, 后续可以根据userId去查找自己之前的记录 
    isAgree: string // 可选
    wantRecognizeLans // 可选
    ```
   
2. 查询历史识别结果

    Method: GET
    
    URL: /query
    
    ```
    userId: string // 可选
    ```

## 线上实验环境地址

提示: 服务器位于AWS Lightsail服务器, 速度会较慢, 请耐心等待, 界面太丑了, 我会后续完善一下下😁

[点击进入](http://node.fenr.men:8001)

## 后续安排

- [ ] 前端美化
- [x] Docker化

    Docker 使用指南
    
    根据数据和应用互相隔离的原则, 使用两个容器, 一个mysql容器, 另外一个则是应用容器
    使用方法如下:
    首先, 确保本机安装docker环境
    在./mysql/dockerScript/ 目录下, 运行runmysqlOCR.sh 脚本, 执行之后使用`docker inspect ocrmysql`拿到容器的地址, 然后填写到`mysql/mysql_config.json`文件的ip地址中, 容器暴露到本地的端口默认是3310, 觉得不爽在`mysql/dockerScript/runmysqlOCR.sh`中改成自己想要的就好了
    在`./`下, 运行`build.sh`脚本, 即可在http://0.0.0.0:4001访问哟~
- [ ] 测试代码