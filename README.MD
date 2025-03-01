# Resolume Go Client

GoによるResolume Arena & Avenue REST APIクライアントライブラリです。

## 機能

- 製品情報の取得
- エフェクト情報の取得
- ソース情報の取得
- サムネイル関連機能
  - サムネイルの取得
  - カスタムサムネイルの設定
  - サムネイルのリセット

## インストール

```bash
go get github.com/FlowingSPDG/resolume-go
```

## 使用方法

### 基本的な使用方法

```go
package main

import (
    "fmt"
    "log"

    "github.com/FlowingSPDG/resolume-go"
)

func main() {
    // クライアントの作成
    client, err := resolume.NewClient("localhost", "8080")
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // 製品情報の取得
    product, err := client.GetProduct()
    if err != nil {
        log.Fatalf("Failed to get product info: %v", err)
    }

    fmt.Printf("Product: %s\n", product.Name)
    fmt.Printf("Version: %d.%d.%d (revision %d)\n",
        product.Major,
        product.Minor,
        product.Micro,
        product.Revision)
}
```

### エフェクト情報の取得

```go
effects, err := client.GetEffects()
if err != nil {
    log.Fatalf("Failed to get effects: %v", err)
}

for _, effect := range effects.Video {
    fmt.Printf("Effect: %s (ID: %s)\n", effect.Name, effect.IDString)
}
```

### サムネイル関連機能

```go
// サムネイルの取得
thumbnail, err := client.GetClipThumbnail(1, 1) // Layer 1, Clip 1
if err != nil {
    log.Fatalf("Failed to get thumbnail: %v", err)
}
defer thumbnail.Close()

// カスタムサムネイルの設定
file, err := os.Open("custom_thumbnail.png")
if err != nil {
    log.Fatalf("Failed to open thumbnail: %v", err)
}
defer file.Close()

err = client.SetClipThumbnail(1, 1, file)
if err != nil {
    log.Fatalf("Failed to set thumbnail: %v", err)
}

// サムネイルのリセット
err = client.ResetClipThumbnail(1, 1)
if err != nil {
    log.Fatalf("Failed to reset thumbnail: %v", err)
}
```

## サンプル

`example`ディレクトリに以下のサンプルコードが用意されています：

- `product/main.go` - 製品情報取得のサンプル
- `test/main.go` - 基本的なAPI機能のテスト
- `thumbnail/main.go` - サムネイル関連機能のサンプル

## ライセンス

MIT License

## 参考

- [Resolume Arena & Avenue REST API Documentation](https://resolume.com/docs/restapi)