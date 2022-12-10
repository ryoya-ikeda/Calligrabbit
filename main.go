package main 

import (
    "github.com/gin-gonic/gin"
    "encoding/csv"
    "log"
    "os"
    // 乱数生成用
    "time"
    "math/rand"
)

func main(){
    r := gin.Default()
    r.GET("/words", func(ctx *gin.Context){
        word0, word1  := GetWords()
        ctx.JSON(200, gin.H{"0":word0,"1":word1,})
    })
    //ポート指定
    r.Run(":8080")
}

// 2つのcsvファイルから1つずつランダムに文字を選ぶ
func GetWords()(string, string) {
    
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
    rand.Seed(time.Now().UnixNano())//現在時刻をシード値に
    return rows0[rand.Intn(LengthOfFile0)][0], rows1[rand.Intn(LengthOfFile1)][0]

}

