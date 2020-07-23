package mysql

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"teletraan/public"
)

func ReadConfig(fileName string) string {
	type MysqlConfig struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Ip string `json:"ip"`
		Port string `json:"port"`
	}
	f, err := os.Open(fileName)
	if err!=nil {
		log.Fatalln(err)
	}
	var m MysqlConfig
	err = json.NewDecoder(f).Decode(&m)
	if err!=nil {
		log.Fatalln(err)
	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Ip + ":" + m.Port + ")/db_play"
}

var dataSource = ReadConfig("mysql/mysql_config.json")

func InitMysqlConn() *sql.DB {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("conn establish error\n")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("conn establish error\n")
	}
	return db
}

// conn 设置为全局变量, 方便复用, 初始化失败时, 服务不能启动
var conn = InitMysqlConn()

func SaveResultToMysql(userId string, content string, timeString string, path string) error {
	_, err := conn.Exec("INSERT INTO ocr_history (userId, result, time, path) VALUES (?, ?, ?, ?)", userId, content, timeString, path)
	if err != nil {
		log.Printf("%+v insert error: %+v", conn, err)
		return err
	}
	return nil
}

type QueryOCRresult struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

type QueryOCRresult1 struct {
	Id      int    `json:"id"`
	Content []string `json:"content"`
	Time    string `json:"time"`
}

func QueryData(userId string) ([]QueryOCRresult1, error) {
	rows, err := conn.Query("SELECT id, result, time FROM ocr_history WHERE userId = ?", userId)
	if err != nil {
		return []QueryOCRresult1{}, err
	}
	defer rows.Close()
	result := make([]QueryOCRresult1, 0)
	for rows.Next() {
		s := QueryOCRresult{}
		err := rows.Scan(&s.Id, &s.Content, &s.Time)
		s1 := QueryOCRresult1{s.Id, public.ConvertRunesToStrings([]rune(s.Content)), s.Time}
		if err != nil {
			return []QueryOCRresult1{}, err
		}
		result = append(result, s1)
	}
	return result, nil
}
