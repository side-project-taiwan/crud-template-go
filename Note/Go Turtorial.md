# Go-Project-Template-Temp

1. fill .env
2. go run cmd/main.go

# *identifier
*是指標的標識符。用於指示變量是指標類型。
Test1都是指標類型不直接存儲資料，是存儲了指向其他資料的記憶體地址。
Test1 *controller.Test1ontroller 表示一個指向 controller.Test1ontroller 型別的指標。
controller.Test1ontroller 是一個物件
Test1 是一個指標指向記憶體地址助於傳遞物件的引用，而不是直接複製物件的值，提高性能並允許對物件進行原地修改。