# Go-Project-Template-Temp

1. fill .env
2. go run cmd/main.go

# *identifier
*是指標的標識符。用於指示變量是指標類型。
Test1 *controller.Test1ontroller [Test1]表示一個指向 controller.Test1ontroller 型別的指標。
Test1指標類型不直接存儲資料，是存儲指向資料的記憶體地址。
Test1 是一個指標指向記憶體地址助於傳遞物件的引用，而不是直接複製物件的值，提高性能並允許對物件進行原地修改。


# go - Public, Private - Upper Case, Lower Case
Go 中，函數名首字母小寫表示函數是私有的，,即便同一個package也只能在同一go file內部存取。
[example]=>service.initStocksService(db) 改為 service.InitStocksService(db)。
