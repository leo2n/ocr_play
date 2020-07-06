package handleImage

import (
	"encoding/json"
	"log"
)

type OCRresponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Content []string `json:"content"`
	Strings string   `json:"strings"`
}

func commonResp(c OCRresponse) ([]byte, error) {
	v, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return []byte("a"), err
	}
	return v, nil
}
