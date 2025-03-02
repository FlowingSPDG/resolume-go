# Resolume Arena & Avenue REST API Client for Go

このライブラリは、Resolume ArenaおよびAvenueのREST APIをGo言語から利用するためのクライアントライブラリです。

## 機能

- 製品情報の取得
- エフェクトの一覧取得と操作
- ソースの一覧取得
- パラメータの取得と設定
- コンポジションの操作
- レイヤーの操作
- レイヤーグループの操作
- デッキの操作
- クリップの操作
- サムネイルの取得と設定

## インストール

```bash
go get github.com/FlowingSPDG/resolume-go
```

## 使用方法

### クライアントの初期化

```go
client, err := resolume.NewClient("localhost", "8080")
if err != nil {
    log.Fatal(err)
}
```

### 製品情報の取得

```go
product, err := client.GetProduct()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Product: %s\n", product.Name)
fmt.Printf("Version: %d.%d.%d\n", product.Major, product.Minor, product.Micro)
```

### エフェクトの一覧取得

```go
effects, err := client.GetEffects()
if err != nil {
    log.Fatal(err)
}
for _, effect := range effects.Video {
    fmt.Printf("Effect: %s (ID: %s)\n", effect.Name, effect.IDString)
}
```

### レイヤーの操作

```go
// レイヤーの取得
layer, err := client.GetLayer(1)
if err != nil {
    log.Fatal(err)
}

// レイヤーの更新
layer.Name = &resolume.StringParameter{Value: "New Layer Name"}
err = client.ReplaceLayer(1, layer)
if err != nil {
    log.Fatal(err)
}
```

### クリップの操作

```go
// クリップの取得
clip, err := client.GetClipByPosition(1, 1)
if err != nil {
    log.Fatal(err)
}

// クリップの接続
err = client.ConnectClipByID(clip.ID, nil) // nilはクリックを模擬
if err != nil {
    log.Fatal(err)
}
```

### サムネイルの操作

```go
// サムネイルの取得
thumbnail, err := client.GetClipThumbnail(1, 1)
if err != nil {
    log.Fatal(err)
}
defer thumbnail.Close()

// サムネイルの保存
file, err := os.Create("thumbnail.png")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
io.Copy(file, thumbnail)
```

## サンプルコード

より詳細な使用例は `example` ディレクトリを参照してください：

- `example/product/main.go` - 製品情報の取得
- `example/test/main.go` - エフェクトとソースの一覧取得
- `example/thumbnail/main.go` - サムネイルの操作

## ライセンス

MIT License

## 注意事項

このライブラリを使用するには、ResolumeのWebサーバー機能を有効にする必要があります。
設定は環境設定の「Webserver」セクションで行えます。