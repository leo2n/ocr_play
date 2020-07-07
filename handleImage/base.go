package handleImage

import (
	"encoding/json"
	"github.com/h2non/filetype"
	"log"
	"teletraan/mysql"
)

type OCRresponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Content []string `json:"content"`
	//Strings string   `json:"strings"`
}

type Queryresponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Content []mysql.QueryOCRresult1 `json:"content"`
	//Strings string   `json:"strings"`
}

// 将struct转为[]byte
func commonResp(c interface{}) ([]byte, error) {
	v, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return []byte("a"), err
	}
	return v, nil
}

// 根据文件的二进制中的特征字符来判断是否是某一个类型的文件
func detectFileTypeByBytes(b []byte) (string, error) {
	kind, err := filetype.Image(b)
	//log.Printf("%+v\n", kind)
	if err != nil {
		log.Printf("file detect error")
		return "", err
	}
	return kind.MIME.Value, nil
}