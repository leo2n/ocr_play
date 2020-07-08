package public

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//convert []rune to []string, remove空格, \r, \n
func ConvertRunesToStrings(r []rune) []string {
	m := make([]string, 0, len(r))
	// remove \r \n space
	for _, v := range r {
		if v == 32 || v == 10 || v == 13 {
			continue
		}
		m = append(m, string(v))
	}
	return m
}

// save file on ./imageStore directory
func SaveFileOnSpecificPath(fileBytes []byte, fileName string) error {
	cwd, err := os.Getwd()
	log.Printf("%v", cwd)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(cwd+"/imageStore/"+fileName, fileBytes, 400)
	if err != nil {
		return err
	}
	return nil
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/teletraan.ico")
}
