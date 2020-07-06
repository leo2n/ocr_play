# OCR识别Demo

## 环境

    System: Ubuntu 20.04 LTS
    Go version: go 1.14/4
    tesseract: 4.1.1
    "github.com/otiai10/gosseract/v2"
    Mysql server 8.0

## 接口设计

1. 上传图片并获取结果

    Method: POST
    
    URL: /ocr
    
    Content-Type: multipart/form-data
    ```text
       imgFile: [file] binary
       userId: userID string
    ```
   
    
    
    

2. 将获取的结果做成可查询的状态
