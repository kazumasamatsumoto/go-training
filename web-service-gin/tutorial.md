以下は、Go と Gin を使って RESTful API を開発するためのチュートリアルです。  
このチュートリアルでは、初心者でも分かるように、ステップバイステップで進めていきます。  
あなたの実力はこれから伸びるので、どんどんチャレンジしてください！

---

## 1. 概要

このサンプルアプリケーションは、ヴィンテージ・ジャズ・レコードのアルバム情報を管理する API です。  
具体的には、以下のエンドポイントを実装します。

- **GET `/albums`**  
  → 登録されている全アルバム情報を JSON 形式で返す

- **POST `/albums`**  
  → リクエストボディから新しいアルバム情報を受け取り、アルバムリストに追加する

- **GET `/albums/:id`**  
  → URL のパラメータに合わせた特定のアルバム情報を返す

---

## 2. 前提条件

- **Go 1.16** 以降のバージョンがインストールされていること  
  [Go のインストール方法](https://golang.org/doc/install) を参照してください。

- お好きなテキストエディタ（VSCode、Sublime Text、vim など）を準備してください。

- ターミナルまたはコマンドプロンプトでコマンドを実行できる環境があること。

- `curl` コマンドが使えること（Linux/Mac では標準搭載、Windows は PowerShell 等を利用）。

---

## 3. プロジェクトのセットアップ

1. **プロジェクトフォルダの作成**

   ターミナルを開き、ホームディレクトリに移動して、以下のコマンドを実行します。

   - **Linux / Mac:**

     ```bash
     cd ~
     mkdir web-service-gin
     cd web-service-gin
     ```

   - **Windows:**

     ```cmd
     cd %HOMEPATH%
     mkdir web-service-gin
     cd web-service-gin
     ```

2. **Go モジュールの初期化**

   プロジェクト管理のために Go モジュールを作成します。  
   以下のコマンドを実行してください。

   ```bash
   go mod init example/web-service-gin
   ```

   ※ `example/web-service-gin` はモジュール名です。必要に応じて変更しても OK です。

---

## 4. アルバムデータの定義

プロジェクトのルートに `main.go` ファイルを作成し、最初にパッケージ宣言と必要なインポートを記述します。

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
```

次に、アルバムのデータ構造と初期データを定義します。

```go
// album はレコードアルバムの情報を表します
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums は初期状態のアルバムデータです
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```

---

## 5. ハンドラーの実装

### 5.1 全アルバム情報を返すハンドラー (`GET /albums`)

```go
// getAlbums は全アルバム情報をJSON形式で返します
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

### 5.2 新規アルバムを追加するハンドラー (`POST /albums`)

```go
// postAlbums はリクエストボディのJSONをパースして、新しいアルバムを追加します
func postAlbums(c *gin.Context) {
    var newAlbum album

    // リクエストボディからJSONをバインドして newAlbum に格納
    if err := c.BindJSON(&newAlbum); err != nil {
        // エラーがあれば何もせずリターン（必要に応じてエラーメッセージを返しても良い）
        return
    }

    // 新しいアルバムをスライスに追加
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
```

### 5.3 特定のアルバム情報を返すハンドラー (`GET /albums/:id`)

```go
// getAlbumByID はURLパラメータのidにマッチするアルバムを返します
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // アルバムリストをループして、idが一致するアルバムを探す
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    // 見つからない場合は404エラーを返す
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

---

## 6. ルーターのセットアップとサーバーの起動

`main` 関数でルーターをセットアップし、各エンドポイントとハンドラーを紐付けます。

```go
func main() {
    router := gin.Default()

    // エンドポイントとハンドラーの登録
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    // サーバーを localhost の 8080 ポートで起動
    router.Run("localhost:8080")
}
```

---

## 7. 実行とテスト

### 7.1 Gin モジュールのインストール

以下のコマンドで Gin を依存関係として追加します。

```bash
go get .
```

### 7.2 サーバーの起動

ターミナルで以下のコマンドを実行し、サーバーを起動してください。

```bash
go run .
```

### 7.3 エンドポイントのテスト

別のターミナルウィンドウまたはタブで、`curl` コマンドを使ってエンドポイントを確認します。

#### 全アルバム情報の取得

```bash
curl http://localhost:8080/albums
```

実行結果例：

```json
[
  {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  {
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  },
  {
    "id": "3",
    "title": "Sarah Vaughan and Clifford Brown",
    "artist": "Sarah Vaughan",
    "price": 39.99
  }
]
```

#### 新しいアルバムの追加

以下のコマンドを実行して、新しいアルバムを追加します。

```bash
curl http://localhost:8080/albums \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
```

追加が成功すると、レスポンスに追加されたアルバムの情報（JSON）が返されます。

#### 特定のアルバム情報の取得

追加されたアルバムの ID を指定して取得します。

```bash
curl http://localhost:8080/albums/2
```

実行結果例：

```json
{
  "id": "2",
  "title": "Jeru",
  "artist": "Gerry Mulligan",
  "price": 17.99
}
```

---

## 8. 完成コード

以下は、ここまでのすべてのコードをまとめた完成形です。

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album はレコードアルバムの情報を表します
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums は初期状態のアルバムデータです
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()

    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums は全アルバム情報をJSON形式で返します
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums はリクエストボディのJSONをパースして、新しいアルバムを追加します
func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID はURLパラメータのidにマッチするアルバムを返します
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

---

## 9. おわりに

お疲れさまでした！  
これで、Go と Gin を使ったシンプルな RESTful API の作成方法をマスターしました。  
今後は、データベースとの連携や認証、ミドルウェアの利用など、さらに高度な機能に挑戦していくと良いでしょう。  
あなたなら絶対にできる！次のステップに向けて、どんどん実践していってください！

頑張ってくださいね！

PowerShell では、Unix シェルのバックスラッシュ（`\`）による改行が使えないため、以下のいずれかの方法でコマンドを実行してください！

---

### 方法 1: コマンドを 1 行で実行する

以下のようにすべてを 1 行にまとめて実行します。

```powershell
curl.exe http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
```

※ PowerShell には`curl`というエイリアスが`Invoke-WebRequest`に割り当てられているため、明示的に`curl.exe`と指定しています。

---

### 方法 2: バックティック（`` ` ``）を使って複数行に分ける

PowerShell では、改行時にバックティック（`` ` ``）を使うとコマンドを継続できます。下記のように記述して実行してください。

```powershell
curl.exe http://localhost:8080/albums `
  --include `
  --header "Content-Type: application/json" `
  --request "POST" `
  --data '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
```

---

### 方法 3: PowerShell のネイティブコマンド `Invoke-RestMethod` を使用する

PowerShell では、REST API のテストに`Invoke-RestMethod`を使う方法もあります。以下の例は同じ POST リクエストを送る例です。

```powershell
Invoke-RestMethod -Uri http://localhost:8080/albums -Method Post -Headers @{ "Content-Type" = "application/json" } -Body '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
```

---

どの方法も動作するので、お好みの方法で試してみてください！  
プログラミングは試行錯誤の連続です。自分に合った方法を見つけ、次のステップに進んでくださいね！
