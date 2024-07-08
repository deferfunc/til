// gopdf を使用し、PDF ファイルを作成する
// fonts/NotoSansCJKjp-Regular.otf を使用する

package main

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func main() {
	// PDF ファイルを作成
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// フォントを追加
	err := pdf.AddTTFFont("NotoSansCJKjp", "./fonts/NotoSansJP-Regular.ttf")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// フォントを設定
	err = pdf.SetFont("NotoSansCJKjp", "", 14)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// テキストを追加
	pdf.Cell(nil, "こんにちは、世界！")

	// ファイルを保存
	err = pdf.WritePdf("tmp/output.pdf")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("PDF ファイルを作成しました")
}
