# 導入方法
https://github.com/gaku3601/auto-generate-sql-for-e2e/releases  
から各環境に適した実行ファイルをダウンロードする  
## mac
agsをダウンロードして.zshrcなどで定義しているpathの通っているところへ配置する(/usr/local/bin等)  
以下コマンドで実行権限を与える  
```
chmod 777 agse2e
```

agse2e -hでヘルプが表示されれば利用可能です  
セキュリティ周りでpopupが出る場合はmacのセキュリティープライバシーの設定からagsを許可してください　　

## win
agse2e.exeをダウンロードしてpathの通っているところへ配置する  

# 使い方
```
agse2e generate -p [対象エクセルファイル]
```
を実行することで自動的にSQLファイルが作成される  
SQLは対象エクセルファイルと同一のフォルダに配置される

## debug

```
go run main.go generate -p ./example/ex1.xlsx
```
でxlsxファイルと同じ場所にsqlファイルが生成される

## build
script/build.shを回せば、distフォルダにバイナリファイルが格納される
