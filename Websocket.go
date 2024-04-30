package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

// User 定义一个数据模型(user表)
// 列名是字段名的蛇形小写(PassWd->pass_word)
type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

var db *gorm.DB

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	//处理日志保存
	// 创建日志文件
	logFile, err := os.OpenFile("example.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err) // 如果无法创建文件，使用 Fatal 日志并终止程序
	}
	defer logFile.Close() // 确保文件在函数结束时关闭

	// 设置日志输出到文件
	log.SetOutput(logFile)

	// 记录日志
	log.Println("这是一个日志记录")
	log.Println("这是另一个日志记录")

	//数据库操作
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/qiu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("创建数据库连接失败:%v", err)
	}
	defer db.Close()

	// 自动迁移数据结构(table schema)
	// 注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
	// db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
	db.AutoMigrate(&User{})

	// 添加唯一索引
	db.Model(&User{}).AddUniqueIndex("name_email", "id", "name", "email")

	// 插入记录
	db.Create(&User{Name: "bgbiao", Age: 18, Email: "bgbiao@bgbiao.top"})
	db.Create(&User{Name: "xxb", Age: 18, Email: "xxb@bgbiao.top"})

	var user User
	var users []User
	// 查看插入后的全部元素
	fmt.Printf("插入后元素:\n")
	log.Printf("插入后元素:\n")
	db.Find(&users)
	fmt.Println(users)
	log.Println(users)

	// 查询一条记录
	db.First(&user, "name = ?", "bgbiao")
	fmt.Println("查看查询记录:", user)
	log.Println("查看查询记录:", user)

	// 更新记录(基于查出来的数据进行更新)
	db.Model(&user).Update("name", "biaoge")
	fmt.Println("更新后的记录:", user)
	log.Println("更新后的记录:", user)

	// 删除记录
	db.Delete(&user)

	// 查看全部记录
	fmt.Println("查看全部记录:")
	log.Println("查看全部记录:")

	db.Find(&users)
	fmt.Println(users)
	log.Println(users)

	// 创建HTTP服务器
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 处理WebSocket连接
	for {
		// 读取消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Received message:", string(p))

		// 发送消息
		err = conn.WriteMessage(messageType, []byte("Hello, world!"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
