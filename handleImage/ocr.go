package handleImage

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"teletraan/mysql"
	"teletraan/public"
	"time"
)

var maxFileSize int64 = 10 << 20 // 10 * 1 * 1024 *1024, 1 << 10 = 1024, 单位byte

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/index.gtpl")
		err := t.Execute(w, nil)
		if err!=nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
}

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
	userId := r.PostForm.Get("userId")              // 没有的话默认是空字符串, 也是会写入到mysql中的哦  "" != null
	isAgreeUsStoreFile := r.PostForm.Get("isAgree") // "yes"是同意, "no"是抗议
	wantRecognizeLans := r.PostForm.Get("wantRecognizeLans")
	// language 默认为english
	if wantRecognizeLans == "" {
		wantRecognizeLans = defaultSupportLanguages
	}
	log.Printf("%+v 的lansList是:%s", r.RemoteAddr, wantRecognizeLans)
	if isInSupportLanguagesRange(wantRecognizeLans) == false {
		log.Printf("对不起, 传入的语言%s 不支持", wantRecognizeLans)
		http.Error(w, fmt.Sprintf("对不起, 传入的语言%s 不支持", wantRecognizeLans), 400)
		return
	}

	file, _, err := r.FormFile("imgFile")
	if err != nil {
		log.Println("获取文件失败!", err)
		http.Error(w, "获取文件失败!", 400)
		return
	}
	defer file.Close()
	// 根据文件扩展名来判定是否是图片文件?
	//fileName := handler.Filename
	//// 获取后缀, 如果不是jpg || png, 返回错误信息
	//suffix := strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1]
	//if suffix != "jpg" && suffix != "png" {
	//	http.Error(w, "support image type is: jpg || png", 403)
	//	return
	//}

	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v bytes\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	imgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("将文件从io流转为bytes时出现错误", err)
		http.Error(w, "将文件从io流转为bytes时出现错误", 500)
		return
	}
	// 根据用户的选择, 是否将文件保存在服务器上
	var fileStoreName string
	if isAgreeUsStoreFile == "yes" {
		fileStoreName = ksuid.New().String()
		err = public.SaveFileOnSpecificPath(imgBytes, fileStoreName)
		if err != nil {
			log.Printf("save file error")
		}
	}
	// 根据filetype判断是否是图片文件?
	kind, err := detectFileTypeByBytes(imgBytes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if kind != "image/jpg" && kind != "image/png" {
		http.Error(w, "file type not allowed", 500)
		return
	}

	// 将[]byte转为[]rune
	runes, err := GetImageContent(imgBytes, strings.Split(wantRecognizeLans, ";"))
	if err != nil {
		log.Printf("%+v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}
	recognize := public.ConvertRunesToStrings(runes)
	// 不管存在不存在userId, 都把它保存在Mysql中, 保存的形式暂定为: id int, user_id varchar(128), result text, create_time timestamp ()
	err = mysql.SaveResultToMysql(userId, string(runes), time.Now().Format("2006-01-02 15:04:05"), fileStoreName)
	if err != nil {
		log.Printf("Write data to mysql error: %s", err.Error())
	}
	// 把结果包装成json形式
	ocrResponse := OCRresponse{
		Code:    8001,
		Msg:     "解析成功!",
		Content: recognize,
		//Strings: string(runes),
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

// 根据userId来获取OCR识别的历史记录
func QueryOCR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request Method restrict POST", 405)
		return
	}

	userIdList, ok := r.URL.Query()["userId"]
	if !ok || len(userIdList) == 0 {
		log.Printf("Url Param 'userId' is missing")
		http.Error(w, "url param 'userId' is missing", 400)
		return
	}
	userId := userIdList[0]

	QueryResult, err := mysql.QueryData(userId)
	if err != nil {
		log.Printf("query %s error\n", userId)
		http.Error(w, err.Error(), 500)
		return
	}
	// 将query结果转为json形式发送
	queryResponse := Queryresponse{
		Code:    8001,
		Msg:     "解析成功!",
		Content: QueryResult,
		//Strings: string(runes),
	}
	content, err := commonResp(queryResponse)
	if err != nil {
		log.Printf("从结构体转bytes出错了", err)
		http.Error(w, "从结构体转bytes出错了", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
	return
}
