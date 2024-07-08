package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// Chromedpのコンテキストを作成
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// タイムアウトを設定
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// タスクを実行
	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://yahoo.co.jp`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// PDFを生成
			var err error
			buf, _, err = page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			return err
		}),
	); err != nil {
		log.Fatal(err)
	}

	// PDFファイルを保存
	if err := saveToFile("tmp/yahoo.pdf", buf); err != nil {
		log.Fatal(err)
	}

	log.Println("PDF ファイルを作成しました")
}

// saveToFile はバイトスライスをファイルに保存します。
func saveToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
