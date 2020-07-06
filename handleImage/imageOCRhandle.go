package handleImage

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var maxFileSize int64 = 10 << 20 // 10 * 1 * 1024 *1024, 1 << 10 = 1024, 单位byte

// 处理用户上传的图片并返回OCR结果
func OCR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "请求方法请使用POST", 405)
		return
	}
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, "File Size can not bigger than 10 MB", 403)
		return
	}
	//userId := r.PostForm.Get("userId")
	file, handler, err := r.FormFile("imgFile")
	if err != nil {
		log.Println("获取文件失败!", err)
		http.Error(w, "获取文件失败!", 400)
		return
	}
	defer file.Close()
	fileName := handler.Filename
	// 获取后缀, 如果不是jpg || png, 返回错误信息
	suffix := strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1]
	if suffix != "jpg" && suffix != "png" {
		http.Error(w, "support image type is: jpg || png", 403)
		return
	}
	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v bytes\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	imgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("将文件从io流转为bytes时出现错误", err)
		http.Error(w, "将文件从io流转为bytes时出现错误", 500)
		return
	}
	runes, err := GetImageContent(imgBytes)
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	recognize := convertRunesToStrings(runes)
	// 不管存在不存在userId, 都把它保存在Mysql中

	// 把结果包装成json形式
	ocrResponse := OCRresponse{
		Code:    8001,
		Msg:     "解析成功!",
		Content: recognize,
		Strings: string(runes),
	}
	content, err := commonResp(ocrResponse)
	if err != nil {
		log.Printf("从结构体转bytes出错了", err)
		http.Error(w, "从结构体转bytes出错了", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
	return
}
