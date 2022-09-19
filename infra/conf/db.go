package conf

import (
	"encoding/json"
	"io/ioutil"
)

// databaseの基本情報を入れた構造体
type ConfDB struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	DbName  string `json:"db-name"`
	Charset string `json:"charset"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
}

func ReadConfDB() (*ConfDB, error) {
	const conffile = "infra/conf/db.json"
	//newは指定した型のポインタ型を生成する関数
	conf := new(ConfDB)
	//[]byteにファイルを読み込む
	readFile, err := ioutil.ReadFile(conffile)
	if err != nil {
		return conf, err
		//fmt.Errorf("failed to read json conf file: %w", err)
	}
	//confにデータを収納
	err = json.Unmarshal(readFile, conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
