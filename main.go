package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// 商品構造体
// 実際のデータベース構造とは異なるが、フロントから情報を受け取る際はこのような形となっているため
// ポスト時
type Product struct {
	ProductName string   // 商品名
	BrandName   string   //ブランド名
	ImagePathes []string // 画像パス
}

// フロントに商品情報を表示するとき専用の構造体
type Display struct {
	ProductID   int    // 商品ID
	ProductName string // 商品名
	BrandID     int    // ブランドID
	BrandName   string //ブランド名
	ImageID     int    //画像ID
	ImagePath   string // 画像パス
}

// ブランド構造体
type Brand struct {
	BrandID   int    // ブランドID
	BrandName string // ブランド名
}

// 画像構造体
type Image struct {
	ImageID   int    // 画像ID
	ProductID int    // 商品ID
	ImagePath string // 画像パス
}

// データベースへの接続
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// データベースの初期化
func initializeDB() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブルの作成
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		product_id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT NOT NULL,
		brand_id INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS brands (
		brand_id INTEGER PRIMARY KEY AUTOINCREMENT,
		brand_name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS images (
		image_id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER NOT NULL,
		image_path TEXT NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// 商品情報を全て取得するAPIハンドラー
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT p.product_id, p.product_name, b.brand_id, b.brand_name, i.image_id, i.image_path 
	FROM products p 
	JOIN brands b ON p.brand_id = b.brand_id 
	JOIN images i ON p.product_id = i.product_id
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var displays []Display
	for rows.Next() {
		var d Display
		err := rows.Scan(&d.ProductID, &d.ProductName, &d.BrandID, &d.BrandName, &d.ImageID, &d.ImagePath)
		if err != nil {
			log.Fatal(err)
		}
		displays = append(displays, d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(displays)
}

// 商品を追加するAPIハンドラー
func addProduct(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var p Product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// var brandname string
	var brandID int64
	var count_brand int //ブランドが重複するかどうかのフラグ
	// ブランドがすでに存在するか確認
	err = db.QueryRow("SELECT COUNT(*) FROM brands WHERE brand_name = ?", p.BrandName).Scan(&count_brand)
	if err != nil {
		log.Fatal(err)
	}
	if count_brand == 0 {
		// ブランドが存在しない場合は新規追加
		result, err := db.Exec("INSERT INTO brands (brand_name) VALUES (?)", p.BrandName)
		if err != nil {
			log.Fatal(err)
		}
		brandID, _ = result.LastInsertId()
		fmt.Println("ブランドを新規追加しました")
	} else {
		_ = db.QueryRow("SELECT brand_id from brands WHERE brand_name = ?", p.BrandName).Scan(&brandID)
		fmt.Println("このブランドはすでに存在します")
	}

	// 商品を追加、重複ある場合は
	var productID int64
	var count_product int
	err = db.QueryRow("SELECT COUNT(*) FROM products WHERE (product_name, brand_id) = (?, ?)", p.ProductName, brandID).Scan(&count_product)
	if err != nil {
		log.Fatal(err)
	}
	if count_product == 0 {
		result2, err := db.Exec("INSERT INTO products (product_name, brand_id) VALUES (?, ?)", p.ProductName, brandID)
		if err != nil {
			log.Fatal(err)
		}
		productID, _ = result2.LastInsertId()
		fmt.Println("商品を新規追加しました")
	} else {
		_ = db.QueryRow("SELECT product_id from products WHERE product_name = ?", p.ProductName).Scan(&productID)
		fmt.Println("この商品はすでに存在します")
	}

	// 画像を追加
	for _, imagePath := range p.ImagePathes {
		var count_image_path int
		err = db.QueryRow("SELECT COUNT(*) FROM images WHERE (product_id, image_path) = (?, ?)", productID, imagePath).Scan(&count_image_path)
		if err != nil {
			log.Fatal(err)
		}
		if count_image_path == 0 {
			_, err := db.Exec("INSERT INTO images (product_id, image_path) VALUES (?, ?)", productID, imagePath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("画像パスを新規追加しました")
		} else {
			fmt.Println(("この画像パスはすでに存在します"))
		}
	}

	w.WriteHeader(http.StatusCreated)
}

// 特定のブランド名の商品を取得するAPIハンドラー
func getProductsByBrand(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ブランド名をURLパラメータから取得
	params := mux.Vars(r)
	brandName := params["brand"]

	rows, err := db.Query(`
	SELECT p.product_id, p.product_name, b.brand_id, b.brand_name, i.image_id, i.image_path 
	FROM products p 
	JOIN brands b ON p.brand_id = b.brand_id 
	JOIN images i ON p.product_id = i.product_id
	WHERE b.brand_name = ?
	`, brandName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var displays []Display
	for rows.Next() {
		var d Display
		err := rows.Scan(&d.ProductID, &d.ProductName, &d.BrandID, &d.BrandName, &d.ImageID, &d.ImagePath)
		if err != nil {
			log.Fatal(err)
		}
		displays = append(displays, d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(displays)
}

func main() {
	// データベースの初期化
	initializeDB()

	// ルーターの作成
	r := mux.NewRouter()

	// ルーティングの設定
	r.HandleFunc("/products", getAllProducts).Methods("GET")
	r.HandleFunc("/products", addProduct).Methods("POST")
	r.HandleFunc("/products/{brand}", getProductsByBrand).Methods("GET") // 特定のブランドの商品を取得する

	// サーバーの起動
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
