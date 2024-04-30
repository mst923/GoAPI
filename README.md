# GoAPI仕様書

## 動作確認手順

### 1. 開発環境のセットアップ(mac)
homebrewを使い、以下のコマンドでセットアップします。
```
$ brew install go
```

### 2. モジュールのインストール
git clone完了したらモジュールをインストールします。

```
$ go get github.com/gorilla/mux
$ go get github.com/mattn/go-sqlite3
```

### 3. サーバーの立ち上げ
以下のコマンドを入力して、サーバーを立ち上げます。

```
$ go run main.go
```

立ち上げ成功すると以下が表示されます。
```
Server listening on port 8080...
```

#### データベースの構成

以下のようにブランド名とブランドIDを格納するbrandsテーブル、商品ID、商品名、ブランドIDを格納するproductsテーブル、画像ID、商品ID、画像パスを格納するimagesテーブルの3つのテーブルを作成しました。リレーションが第３正規形となるようにテーブル構成を考えました。
###### brandsテーブル

```
| brand_id | brand_name      |
|----------|-----------------|
| 1        | KENZO           |
| 2        | KENZO           |
| 3        | MIU MIU         |
| 4        | Alexander McQUEEN |
```

###### productsテーブル
```
| product_id | product_name                                      | brand_id |
|------------|---------------------------------------------------|----------|
| 1          | KENZO 'TIGER CREST' POLO SHIRT                   | 1        |
| 2          | A.P.C. 'AURELIA' DENIM DRESS                     | 2        |
| 3          | MIU MIU VINTAGE LEATHER ANKLE BOOTS              | 3        |
| 4          | Alexander McQUEEN 'SLIM TREAD' ANKLE BOOTS       | 4        |
| 5          | STIVALETTO                                        | 4        |
| 6          | Alexander McQUEEN Black leather zipped card holder with logo | 4 |
| 7          | Alexander McQUEEN White and clay Oversize Sneaker | 4        |
| 8          | Alexander McQUEEN Black camera bag with leather details | 4 |
| 9          | A.P.C. 'GRACE SMALL' CROSSBODY BAG               | 2        |
| 10         | A.P.C. JAMIE' CROSSBODY BAG                      | 2        |
```
###### iamgesテーブル
```
データベースの画像テーブル:

| image_id | product_id | image_path                                                                                                         |
|----------|------------|--------------------------------------------------------------------------------------------------------------------|
| 1        | 1          | https://stok.store/cdn/shop/files/20220304040105603_E52---kenzo---FA65PO0014PU01B_4_M1.jpg                       |
| 2        | 1          | https://stok.store/cdn/shop/files/20220304040124567_E52---kenzo---FA65PO0014PU01B_5_M1.jpg                       |
| 3        | 1          | https://stok.store/cdn/shop/files/20220304040128900_E52---kenzo---FA65PO0014PU01B_7_M1.jpg                       |
| 4        | 2          | https://stok.store/cdn/shop/files/20211218140947410_E52---apc---COETKF05822IAL_1_M1.jpg                           |
| 5        | 2          | https://stok.store/cdn/shop/files/20211218140947600_E52---apc---COETKF05822IAL_2_M1.jpg                           |
| 6        | 3          | https://stok.store/cdn/shop/files/5T953DF0503F33F0002_01_M_2024-02-22T08-12-47.316Z.jpg                           |
| 7        | 3          | https://stok.store/cdn/shop/files/5T953DF0503F33F0002_04_M_2024-02-22T08-12-47.566Z.jpg                           |
| 8        | 4          | https://stok.store/cdn/shop/files/20220125140133136_E52---alexander_20mcqueen---690812W4SQ11053_1_M1.jpg           |
| 9        | 5          | https://stok.store/cdn/shop/files/757487WIDU11000_2023-07-07T07-27-52.221Z.jpg                                    |
| 10       | 5          | https://stok.store/cdn/shop/files/757487WIDU11000_5_P_2023-07-07T07-27-52.533Z.jpg                                  |
| 11       | 5          | https://stok.store/cdn/shop/files/757487WIDU11000_3_P_2023-07-07T07-27-52.392Z.jpg                                  |
| 12       | 6          | https://stok.store/cdn/shop/files/6831171AAMJ_O_ALEXQ-1070.a.jpg                                                    |
| 13       | 7          | https://stok.store/cdn/shop/files/718139WIEE5_O_ALEXQ-8742.a.jpg                                                     |
| 14       | 8          | https://stok.store/cdn/shop/files/7262921AAQ0_O_ALEXQ-1000.a.jpg                                                     |
| 15       | 9          | https://stok.store/cdn/shop/files/20230505000427217_A55---apc---COGFAF61413LZZ_1_M1.jpg                           |
| 16       | 9          | https://stok.store/cdn/shop/files/20230505000427378_A55---apc---COGFAF61413LZZ_2_M1.jpg                           |
| 17       | 9          | https://stok.store/cdn/shop/files/20230505000427514_A55---apc---COGFAF61413LZZ_3_M1.jpg                           |
| 18       | 9          | https://stok.store/cdn/shop/files/20230505000432125_A55---apc---COGFAF61413LZZ_4_M1.jpg                           |
| 19       | 10         | https://stok.store/cdn/shop/files/20220221181331577_E52---apc---PXBMWF63412LZZBLACK_3_M1.jpg                      |
| 20       | 10         | https://stok.store/cdn/shop/files/20220221181331639_E52---apc---PXBMWF63412LZZBLACK_4_M1.jpg                      |
```



### 4. データベースに格納されている全ての情報を一括で取得させる
以下のコマンドを叩きます。
```
$ curl http://localhost:8080/products
```

出力は次のようになります。

```
[{"ProductID":1,"ProductName":"KENZO 'TIGER CREST' POLO SHIRT","BrandID":1,"BrandName":"KENZO","ImageID":1,"ImagePath":"https://stok.store/cdn/shop/files/20220304040105603_E52---kenzo---FA65PO0014PU01B_4_M1.jpg"},{"ProductID":1,"ProductName":"KENZO 'TIGER CREST' POLO SHIRT","BrandID":1,"BrandName":"KENZO","ImageID":2,"ImagePath":"https://stok.store/cdn/shop/files/20220304040124567_E52---kenzo---FA65PO0014PU01B_5_M1.jpg"},{"ProductID":1,"ProductName":"KENZO 'TIGER CREST' POLO SHIRT","BrandID":1,"BrandName":"KENZO","ImageID":3,"ImagePath":"https://stok.store/cdn/shop/files/20220304040128900_E52---kenzo---FA65PO0014PU01B_7_M1.jpg"},{"ProductID":2,"ProductName":"A.P.C. 'AURELIA' DENIM DRESS","BrandID":2,"BrandName":"KENZO","ImageID":4,"ImagePath":"https://stok.store/cdn/shop/files/20211218140947410_E52---apc---COETKF05822IAL_1_M1.jpg"},{"ProductID":2,"ProductName":"A.P.C. 'AURELIA' DENIM DRESS","BrandID":2,"BrandName":"KENZO","ImageID":5,"ImagePath":"https://stok.store/cdn/shop/files/20211218140947600_E52---apc---COETKF05822IAL_2_M1.jpg"},{"ProductID":3,"ProductName":"MIU MIU VINTAGE LEATHER ANKLE BOOTS","BrandID":3,"BrandName":"MIU MIU","ImageID":6,"ImagePath":"https://stok.store/cdn/shop/files/5T953DF0503F33F0002_01_M_2024-02-22T08-12-47.316Z.jpg"},{"ProductID":3,"ProductName":"MIU MIU VINTAGE LEATHER ANKLE BOOTS","BrandID":3,"BrandName":"MIU MIU","ImageID":7,"ImagePath":"https://stok.store/cdn/shop/files/5T953DF0503F33F0002_04_M_2024-02-22T08-12-47.566Z.jpg"},{"ProductID":4,"ProductName":"Alexander McQUEEN 'SLIM TREAD' ANKLE BOOTS","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":8,"ImagePath":"https://stok.store/cdn/shop/files/20220125140133136_E52---alexander_20mcqueen---690812W4SQ11053_1_M1.jpg"},{"ProductID":5,"ProductName":"STIVALETTO","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":9,"ImagePath":"https://stok.store/cdn/shop/files/757487WIDU11000_2023-07-07T07-27-52.221Z.jpg"},{"ProductID":5,"ProductName":"STIVALETTO","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":10,"ImagePath":"https://stok.store/cdn/shop/files/757487WIDU11000_5_P_2023-07-07T07-27-52.533Z.jpg"},{"ProductID":5,"ProductName":"STIVALETTO","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":11,"ImagePath":"https://stok.store/cdn/shop/files/757487WIDU11000_3_P_2023-07-07T07-27-52.392Z.jpg"},{"ProductID":6,"ProductName":"Alexander McQUEEN Black leather zipped card holder with logo","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":12,"ImagePath":"https://stok.store/cdn/shop/files/6831171AAMJ_O_ALEXQ-1070.a.jpg"},{"ProductID":7,"ProductName":"Alexander McQUEEN White and clay Oversize Sneaker","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":13,"ImagePath":"https://stok.store/cdn/shop/files/718139WIEE5_O_ALEXQ-8742.a.jpg"},{"ProductID":8,"ProductName":"Alexander McQUEEN Black camera bag with leather details","BrandID":4,"BrandName":"Alexander McQUEEN","ImageID":14,"ImagePath":"https://stok.store/cdn/shop/files/7262921AAQ0_O_ALEXQ-1000.a.jpg"},{"ProductID":9,"ProductName":"A.P.C. 'GRACE SMALL' CROSSBODY BAG","BrandID":2,"BrandName":"KENZO","ImageID":15,"ImagePath":"https://stok.store/cdn/shop/files/20230505000427217_A55---apc---COGFAF61413LZZ_1_M1.jpg"},{"ProductID":9,"ProductName":"A.P.C. 'GRACE SMALL' CROSSBODY BAG","BrandID":2,"BrandName":"KENZO","ImageID":16,"ImagePath":"https://stok.store/cdn/shop/files/20230505000427378_A55---apc---COGFAF61413LZZ_2_M1.jpg"},{"ProductID":9,"ProductName":"A.P.C. 'GRACE SMALL' CROSSBODY BAG","BrandID":2,"BrandName":"KENZO","ImageID":17,"ImagePath":"https://stok.store/cdn/shop/files/20230505000427514_A55---apc---COGFAF61413LZZ_3_M1.jpg"},{"ProductID":9,"ProductName":"A.P.C. 'GRACE SMALL' CROSSBODY BAG","BrandID":2,"BrandName":"KENZO","ImageID":18,"ImagePath":"https://stok.store/cdn/shop/files/20230505000432125_A55---apc---COGFAF61413LZZ_4_M1.jpg"},{"ProductID":10,"ProductName":"A.P.C. JAMIE' CROSSBODY BAG","BrandID":2,"BrandName":"KENZO","ImageID":19,"ImagePath":"https://stok.store/cdn/shop/files/20220221181331577_E52---apc---PXBMWF63412LZZBLACK_3_M1.jpg"},{"ProductID":10,"ProductName":"A.P.C. JAMIE' CROSSBODY BAG","BrandID":2,"BrandName":"KENZO","ImageID":20,"ImagePath":"https://stok.store/cdn/shop/files/20220221181331639_E52---apc---PXBMWF63412LZZBLACK_4_M1.jpg"}]
```


### 5.データベースに商品情報を追加する
以下のようなコマンドを叩きます。ProductName, BrandName, ImagePathesの部分は適宜変更してください。

```
$ curl -X POST -H "Content-Type: application/json" -d '{
  "ProductName": "スタンスミス",
  "BrandName" : "Adidas",
  "ImagePathes": [
    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fkakaku.com%2Fitem%2FS0000964343%2F&psig=AOvVaw14rfO8s2FkBXgaKFDwMJ40&ust=1714588246093000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCLidu4zJ6oUDFQAAAAAdAAAAABAE",
    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.mensnonno.jp%2Ffashion%2Fbasic%2F62470%2F&psig=AOvVaw14rfO8s2FkBXgaKFDwMJ40&ust=1714588246093000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCLidu4zJ6oUDFQAAAAAdAAAAABAJ"
  ]
}' http://localhost:8080/products

```

成功すると、以下のようになります。
```
Server listening on port 8080...
ブランドを新規追加しました
商品を新規追加しました
画像パスを新規追加しました
画像パスを新規追加しました
```

productsテーブルを見ると、ちゃんと追加されていることが確認できます。

```
| product_id | product_name                                                     | brand_id |
|------------|------------------------------------------------------------------|----------|
| 1          | KENZO 'TIGER CREST' POLO SHIRT                                   | 1        |
| 2          | A.P.C. 'AURELIA' DENIM DRESS                                     | 2        |
| 3          | MIU MIU VINTAGE LEATHER ANKLE BOOTS                              | 3        |
| 4          | Alexander McQUEEN 'SLIM TREAD' ANKLE BOOTS                       | 4        |
| 5          | STIVALETTO                                                       | 4        |
| 6          | Alexander McQUEEN Black leather zipped card holder with logo     | 4        |
| 7          | Alexander McQUEEN White and clay Oversize Sneaker                 | 4        |
| 8          | Alexander McQUEEN Black camera bag with leather details           | 4        |
| 9          | A.P.C. 'GRACE SMALL' CROSSBODY BAG                               | 2        |
| 10         | A.P.C. JAMIE' CROSSBODY BAG                                      | 2        |
| 11         | スタンスミス                                                      | 5        |
```

もし、重複した要素を入力すると、重複している要素だけ追加されないようになります。次のようなコマンドを打つと、
```
$ curl -X POST -H "Content-Type: application/json" -d '{
  "ProductName": "ガゼル",
  "BrandName" : "Adidas",
  "ImagePathes": [
    "https://shop.adidas.jp/photo/HQ/HQ8717/z-HQ8717-standard-side_lateral_center_view-qeEAZ6C0Pr.jpg",
    "https://z-shopping.c.yimg.jp/279/79740279/79740279b_34_d_500.jpg"
  ]
}' http://localhost:8080/products
```
重複しているブランドは追加されず、商品と画像パスは重複していないので、追加されます。
```
このブランドはすでに存在します
商品を新規追加しました
画像パスを新規追加しました
画像パスを新規追加しました
```

### 6.指定したブランド名の商品のみを表示する
例えばAdidasの商品のみ表示したい場合は次のようなコマンドを叩きます。指定したURLのproducts以下にブランド名を指定してあげると良いです。
```
$ curl http://localhost:8080/products/Adidas
```

出力は次のようになります。
```
[{"ProductID":11,"ProductName":"スタンスミス","BrandID":5,"BrandName":"Adidas","ImageID":21,"ImagePath":"https://www.google.com/url?sa=i\u0026url=https%3A%2F%2Fkakaku.com%2Fitem%2FS0000964343%2F\u0026psig=AOvVaw14rfO8s2FkBXgaKFDwMJ40\u0026ust=1714588246093000\u0026source=images\u0026cd=vfe\u0026opi=89978449\u0026ved=0CBIQjRxqFwoTCLidu4zJ6oUDFQAAAAAdAAAAABAE"},{"ProductID":11,"ProductName":"スタンスミス","BrandID":5,"BrandName":"Adidas","ImageID":22,"ImagePath":"https://www.google.com/url?sa=i\u0026url=https%3A%2F%2Fwww.mensnonno.jp%2Ffashion%2Fbasic%2F62470%2F\u0026psig=AOvVaw14rfO8s2FkBXgaKFDwMJ40\u0026ust=1714588246093000\u0026source=images\u0026cd=vfe\u0026opi=89978449\u0026ved=0CBIQjRxqFwoTCLidu4zJ6oUDFQAAAAAdAAAAABAJ"},{"ProductID":12,"ProductName":"ガゼル","BrandID":5,"BrandName":"Adidas","ImageID":23,"ImagePath":"https://shop.adidas.jp/photo/HQ/HQ8717/z-HQ8717-standard-side_lateral_center_view-qeEAZ6C0Pr.jpg"},{"ProductID":12,"ProductName":"ガゼル","BrandID":5,"BrandName":"Adidas","ImageID":24,"ImagePath":"https://z-shopping.c.yimg.jp/279/79740279/79740279b_34_d_500.jpg"}]
```

## 作業時間記録表
1. 開発環境セットアップ 5分
2. データベース設計 30分
3. Go言語リサーチ 30分
4. APIの実装 2時間
5. デバッグ 2時間
6. リファクタリング 30分
7. 仕様書記述 30分

## 備考
削除メソッドは実装していない。
フロントエンドも実装していない。ターミナルでcurlコマンドでのアクセスがメインになる。