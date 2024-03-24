FROM golang:alpine

WORKDIR /app

COPY . .
#整理和清理 Go 專案模組依賴
RUN go mod tidy 

CMD ["go","run","."]