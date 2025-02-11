以下は、Go のジェネリクス（総称型）を使ったプログラミングの入門チュートリアルです。これから説明する内容は、Go 1.18 以降の環境が前提となります。素直かつ前向きに取り組んでください。最初は戸惑うかもしれませんが、チャレンジするあなたなら必ず身につけられます！

---

## 1. 前提条件

- **Go のインストール**: Go 1.18 以降が必要です。まだインストールしていなければ、[Installing Go](https://go.dev/doc/install) の手順に従ってください。
- **コードエディタ**: お好みのエディタを使用してください。
- **コマンドライン**: ターミナル（または PowerShell, cmd）を用意しましょう。

---

## 2. 作業用フォルダの作成とモジュールの初期化

まず、作業用のフォルダを作り、Go モジュールを初期化します。

1. ターミナルを開き、ホームディレクトリに移動します。

   - Linux/Mac:
     ```sh
     $ cd
     ```
   - Windows:
     ```sh
     C:\> cd %HOMEPATH%
     ```

2. 新しいフォルダ `generics` を作成し、そのフォルダに移動します。

   ```sh
   $ mkdir generics
   $ cd generics
   ```

3. モジュールを初期化します。ここでは、モジュールパスを `example/generics` としていますが、必要に応じて変更してください。
   ```sh
   $ go mod init example/generics
   ```
   ※ この時点で `go.mod` ファイルが生成されます。

---

## 3. 非ジェネリック関数の作成

まず、ジェネリクスを使わずに、２種類のマップ（`int64` と `float64` の値を持つマップ）の合計を求める関数を作成します。

1. 作業用フォルダ内に `main.go` を作成します。
2. 以下のコードを入力してください。

   ```go
   package main

   import "fmt"

   // SumInts は、マップ m の int64 値を合計して返します。
   func SumInts(m map[string]int64) int64 {
       var s int64
       for _, v := range m {
           s += v
       }
       return s
   }

   // SumFloats は、マップ m の float64 値を合計して返します。
   func SumFloats(m map[string]float64) float64 {
       var s float64
       for _, v := range m {
           s += v
       }
       return s
   }

   func main() {
       // int64 の値を持つマップを初期化
       ints := map[string]int64{
           "first":  34,
           "second": 12,
       }

       // float64 の値を持つマップを初期化
       floats := map[string]float64{
           "first":  35.98,
           "second": 26.99,
       }

       fmt.Printf("Non-Generic Sums: %v and %v\n",
           SumInts(ints),
           SumFloats(floats))
   }
   ```

3. 保存して、以下のコマンドで実行してみましょう。
   ```sh
   $ go run .
   ```
   実行結果は以下のようになります。
   ```
   Non-Generic Sums: 46 and 62.97
   ```

> **励ましの一言**  
> この調子です！基本的な関数が動けば、次はジェネリクスにチャレンジしましょう。

---

## 4. ジェネリック関数の追加

次に、ジェネリクスを使って、整数と浮動小数点数のどちらにも対応可能な１つの関数を作成します。  
ジェネリクスでは、関数に型パラメータを指定して、引数の型に対して柔軟に対応できるようになります。

1. `main.go` に、以下のジェネリック関数 `SumIntsOrFloats` を追加してください。

   ```go
   // SumIntsOrFloats は、マップ m の値を合計します。
   // この関数は、マップの値が int64 または float64 の場合に利用できます。
   func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
       var s V
       for _, v := range m {
           s += v
       }
       return s
   }
   ```

2. 追加後、`main` 関数内に下記のコードを追加し、ジェネリック関数を呼び出します。今回は型引数を明示的に指定します。

   ```go
       fmt.Printf("Generic Sums: %v and %v\n",
           SumIntsOrFloats[string, int64](ints),
           SumIntsOrFloats[string, float64](floats))
   ```

3. 再度、実行して結果を確認します。
   ```sh
   $ go run .
   ```
   結果は以下の通り：
   ```
   Non-Generic Sums: 46 and 62.97
   Generic Sums: 46 and 62.97
   ```

> **ストレートな言い方**  
> 型パラメータを指定して呼び出すのは少し冗長ですが、今は仕方ありません。次で型引数を省略できるようにします。

---

## 5. 型引数の省略（型推論）の利用

Go コンパイラは、関数引数から型引数を推論できる場合があります。これにより、呼び出し側で型引数を省略できます。

1. `main` 関数に以下のコードを追加してください。

   ```go
       fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
           SumIntsOrFloats(ints),
           SumIntsOrFloats(floats))
   ```

2. 実行すると、次のような出力が得られます。

   ```sh
   $ go run .
   ```

   出力:

   ```
   Non-Generic Sums: 46 and 62.97
   Generic Sums: 46 and 62.97
   Generic Sums, type parameters inferred: 46 and 62.97
   ```

> **前向きに進めよう**  
> 型推論のおかげでコードがすっきりしましたね！この機能は開発の負担を大幅に減らしてくれます。

---

## 6. 型制約（Type Constraint）の宣言

同じ制約（ここでは、`int64` または `float64` であること）を再利用できるように、型制約を新たなインターフェースとして宣言します。

1. `main` 関数の上、インポート直後に以下を追加してください。

   ```go
   // Number は、int64 または float64 を許容する型制約です。
   type Number interface {
       int64 | float64
   }
   ```

2. 次に、この型制約を利用したジェネリック関数 `SumNumbers` を作成します。

   ```go
   // SumNumbers は、マップ m の値を合計します。整数と浮動小数点数の両方に対応します。
   func SumNumbers[K comparable, V Number](m map[K]V) V {
       var s V
       for _, v := range m {
           s += v
       }
       return s
   }
   ```

3. `main` 関数に、以下の呼び出しコードを追加します。

   ```go
       fmt.Printf("Generic Sums with Constraint: %v and %v\n",
           SumNumbers(ints),
           SumNumbers(floats))
   ```

4. 再度、コードを実行し、結果が期待通りであることを確認してください。
   ```sh
   $ go run .
   ```
   出力:
   ```
   Non-Generic Sums: 46 and 62.97
   Generic Sums: 46 and 62.97
   Generic Sums, type parameters inferred: 46 and 62.97
   Generic Sums with Constraint: 46 and 62.97
   ```

> **素直なアドバイス**  
> 型制約をインターフェースに切り出すことで、複数の関数で同じ制約を再利用でき、コードがとても読みやすくなります。あなたならこれもマスターできるはずです！

---

## 7. チュートリアルのまとめと完成コード

これで、ジェネリクスの基本を理解するためのチュートリアルは完了です。非ジェネリックな関数と、ジェネリックな関数（型引数を明示的に指定する場合と、推論に任せる場合）、そして型制約の使い方を学びました。

以下に、全体の完成コードを示します。まずはコードを動かして、どんどん理解を深めてください！

```go
package main

import "fmt"

// Number は、int64 または float64 を許容する型制約です。
type Number interface {
	int64 | float64
}

func main() {
	// int64 の値を持つマップを初期化
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// float64 の値を持つマップを初期化
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts は、マップ m の int64 値を合計して返します。
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats は、マップ m の float64 値を合計して返します。
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats は、マップ m の値を合計します。
// この関数は、マップの値が int64 または float64 の場合に利用できます。
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers は、マップ m の値を合計します。整数と浮動小数点数の両方に対応します。
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
```

---

## 最後に

最初は戸惑うかもしれませんが、ジェネリクスはコードの再利用性や拡張性を大きく向上させる強力な機能です。これを理解し活用できれば、あなたのプログラミングスキルは一段と上がります。どんどん実践して、疑問があればその都度調べたり、質問したりしてください。あなたなら必ず上手くいきます！

Happy coding!
