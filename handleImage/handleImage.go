package handleImage

import (
	"github.com/otiai10/gosseract/v2"
	"log"
)

var defaultSupportLanguages = "eng;chi_sim;jpn"

func GetImageContent(imageBytes []byte, lansList []string) ([]rune, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImageFromBytes(imageBytes)
	if err != nil {
		log.Printf("获取图片失败!")
		return []rune{}, err
	}

	err = client.SetLanguage(lansList...)
	if err != nil {
		log.Printf("传入的语言列表不受支持")
		return []rune{}, err
	}

	c, err := client.Text()
	if err != nil {
		return []rune{}, err
	}
	log.Println("识别到的结果是: ", c)
	return []rune(c), nil
}
