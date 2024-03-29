# Go-Project-Template-Temp

1. fill .env
2. go run cmd/main.go

# *identifier
Go語言中，*是指標的標識符，不直接存儲資料，是存儲指向資料的記憶體地址。
*apple 表示指向 [apple] 的指標(pointer) 儲存其他變數的記憶體位址。
&apple 表示指向 [apple] 的記憶體位址(LIKE 0x0000000C)。
``` go
func main() {
    apple := 5
    fmt.Println("水果籃子中的蘋果數量:", apple)
    // 定義一個指向 apple 變數的指標變數
    var pointer *int
    pointer = &apple // 將 apple 變數的地址賦值給指標變數 pointer

    // 通過指標間接修改水果籃子中的蘋果數量
    *pointer = 10 // 將指標指向的變數的值修改為 10

    fmt.Println("通過指標修改後的水果籃子中的蘋果數量:", apple)
}
```

# Go 語言中沒有像 C# 中的（Attributes）或 Java（Annotation）那樣的機制。
Go可以使用結構體標籤（Structtags）來實現類似的功能
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
json:"name" and json:"age" are structure tags used to specify the name of the field during JSON serialization.

# Go - Public, Private - Upper Case, Lower Case
Go 中，函數名首字母小寫表示函數是私有的，即便同一個package也只能在同一go file內部存取。
[example]=>service.initStocksService(db) 改為 service.InitStocksService(db)。

# Go - try...catch 
Go語言中沒有try...catch 結構，Go使用函數傳回值來明確地檢查和處理這些錯誤
     // 嘗試開啟一個文件
     file, err := os.Open("test.txt")
     if err != nil {
         // 如果開啟檔案時發生錯誤，處理錯誤並退出程式
         fmt.Println("Error:", err)
         return
     }
     defer file.Close()

     // 文件成功打開，讀取文件內容
     // (這裡省略了讀取檔案內容的部分)
     fmt.Println("File opened successfully")

