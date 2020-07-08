# OCR识别Demo

## 环境

    System: Ubuntu 20.04 LTS
    Go version: go 1.14/4
    tesseract: 4.1.1
    "github.com/otiai10/gosseract/v2"
    Mysql server 8.0

## 实现功能

1. 支持多语言, 比如: 有一张图片, 里面同时有中文, 英文, 日文, 可以一起检索, 一起输出
(使用rune, 也就是int32
, 也是mysql中的utf8mb4)
2. 用户在识别的时候自己输入一个标识符(userId), 日后可以检索自己的历史记录(有可能和别人的标识符重合, 从而检索出别人的结果, 后续可以做成登陆的形式来避免此类情况)
3. 用户在使用服务时会提示是否将自己的图片资源保留在服务器上, 图片使用ksuid命名, 系统匿名使用. 可以供系统对比图片识别结果和实际结果的差异, 提升服务质量, 当然, 这个是用户可选的
4. 用户可以指定tesseract支持的其他语言, 比如: 俄语, 法语etc

## 接口设计

1. 上传图片并获取识别结果

    Method: POST
    
    URL: /ocr
    
    Content-Type: multipart/form-data
    ```
    imgFile: [file] binary // 非空
    userId: userID string // 可选, 用户需要保存自己之前的记录, 
    isAgree: string // 可选
    wantRecognizeLans // 可选
    ```
   
2. 查询历史识别结果

    Method: GET
    
    URL: /query
    
    ```
    userId: string // 非空
    ```

## 线上实验环境地址

提示: 服务器位于阿里云新加坡轻量应用服务器, 速度会较慢, 请耐心等待, 界面太丑了, 我会后续完善一下下😁

