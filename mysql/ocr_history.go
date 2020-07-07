package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"teletraan/public"
)

func InitMysqlConn() *sql.DB {
	db, err := sql.Open("mysql", "leo:123456@tcp(127.0.0.1:3306)/db_play")
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

func SaveResultToMysql(userId string, content string, timeString string) error {
	_, err := conn.Exec("INSERT INTO ocr_history (userId, result, time) VALUES (?, ?, ?)", userId, content, timeString)
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
