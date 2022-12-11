package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"

	// 乱数生成用
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("./static/syodou/*.tmpl")
	// スタート画面
	r.GET("/start", func(ctx *gin.Context) {
		ctx.HTML(200, "index1.html", gin.H{})
	})

	// 抽選処理
	r.GET("/words", func(ctx *gin.Context) {
		word0, word1 := GetWords()
		ctx.JSON(200, gin.H{"0": word0, "1": word1})
	})
	// 文字追加画面を表示
	r.GET("/add", func(ctx *gin.Context) {
		ctx.HTML(200, "index5.tmpl", gin.H{})
	})

	// 文字追加処理
	r.POST("/complete", Writecsv)

	//抽選画面から遷移
	r.GET("/result", func(ctx *gin.Context) {
		word0, word1 := GetWords()
		time.Sleep(3 * time.Second)
		//time.Sleep(3 * time.Second)
		ctx.HTML(200, "index4.tmpl", gin.H{"First": word0, "Second": word1})
	})

	//ポート指定
	r.Run(":8080")
}

// 2つのcsvファイルから1つずつランダムに文字を選ぶ
func GetWords() (string, string) {

	FileName0 := "words/hoge.csv"
	file0, err := os.Open(FileName0)
	FileName1 := "words/fuga.csv"
	file1, err := os.Open(FileName1)

	if err != nil {
		log.Fatal(err)
	}

	// 関数終了後ファイルを閉じるように設定
	defer file0.Close()
	defer file1.Close()

	// csvファイルを読み込む
	r0 := csv.NewReader(file0)
	rows0, err := r0.ReadAll()
	r1 := csv.NewReader(file1)
	rows1, err := r1.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// 読み込んだファイルの要素数を取得(乱数の上限値を決めるため)
	LengthOfFile0 := len(rows0)
	LengthOfFile1 := len(rows1)

	// 現在時刻をシード値にして乱数を生成
	rand.Seed(time.Now().UnixNano()) //現在時刻をシード値に
	return rows0[rand.Intn(LengthOfFile0)][0], rows1[rand.Intn(LengthOfFile1)][0]

}

// CSVファイルに任意の文字列を格納する関数
func Writecsv(c *gin.Context) {
	text := c.PostForm("moji")
	reg := "\r\n|\n"

	//ブラックリストのCSVと書き込み先のCSVを読み込む
	file2, err := ioutil.ReadFile("words/blacklist.csv")
	file3, err := os.OpenFile("words/fuga.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) // 書き込む先のファイル

	if err != nil {
		panic(err)
	}

	//ブラックリストのCSVを読み込み、検査できるように分割する
	arr1 := regexp.MustCompile(reg).Split(string(file2), -1)

	csvwrite := csv.NewWriter(file3)

	var is_black bool = false
	if arrayContains(arr1, text) {
		is_black = true
		c.HTML(200, "index1.tmpl", gin.H{})
	}
	if !is_black {
		csvwrite.Write([]string{text})
		c.HTML(200, "index6.tmpl", gin.H{})
	}
	csvwrite.Flush()
}

// ブラックリスト内にある文字が存在するのかを確認する関数
func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
