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
type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

// BeforeCreate GORM hook
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	// 实现Birthday自动设置
	if u.Birthday == nil {
		currentTime := time.Now()
		if err := scope.SetColumn("Birthday", &currentTime); err != nil {
			log.Println("birthday设置出错：", err)
		}
	}
	return nil
}

var db *gorm.DB
var logFile *os.File
var err error

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	//数据库操作
	//db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/qiu?charset=utf8&parseTime=True&loc=Local")
	// 初始化数据库连接
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "qiu"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("创建数据库连接失败: %v", err)
	}

	// 自动迁移数据结构
	db.AutoMigrate(&User{})

	// 添加唯一索引
	//db.Model(&User{}).AddUniqueIndex("name_email", "name", "age")

	// 初始化日志文件
	logFile, err = os.OpenFile("example.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}

	// 设置日志输出到文件
	log.SetOutput(logFile)

	// 记录日志文件初始化成功的消息
	log.Println("日志文件初始化成功")
}

func main() {
	defer mainExit()

	// 创建HTTP服务器
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("服务器开启在端口:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, errGrader := upgrader.Upgrade(w, r, nil)
	if errGrader != nil {
		log.Println("连接升级失败：", errGrader)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("关闭连接失败: %v", err)
		}
	}()

	for {
		messageType, p, errRead := conn.ReadMessage()
		if errRead != nil {
			log.Println("读取消息失败/连接关闭", errRead)
			return
		}
		log.Println("收到消息:", string(p))

		// 示例：根据接收到的消息插入新用户
		//newUser := User{Name: "bgbiao", Age: 99999, Email: "bgbia"}
		//result := db.Create(&newUser)
		//if errCreate := result.Error; errCreate != nil {
		//	log.Println("创建数据失败：", errCreate)
		//} else {
		//	log.Printf("创建数据成功, %d rows affected", result.RowsAffected)
		//}

		// start
		// 插入记录
		db.Create(&User{Name: "bgbiao", Age: 99999, Email: "bgbiao@bgbiao.top"})
		db.Create(&User{Name: "xxb", Age: 18, Email: "xxb@bgbiao.top"})
		db.Create(&User{Name: "xxb", Age: 18, Email: "xxb@bgbiao.to"})

		var user User
		var user2 User
		var users []User
		var users1 []User
		var users2 []User
		var users3 []User
		// 查看插入后的全部元素
		log.Printf("插入后元素:\n")
		db.Find(&users)
		log.Println(users)

		// 查询一条记录
		db.First(&user, "name = ?", "bgbiao")
		log.Println("查看查询记录:", user)

		db.First(&user2, "age = ?", 18)
		log.Println("查看记录2：", user2)

		// 更新记录(基于查出来的数据进行更新)
		db.Model(&user2).Update("name", "biaoge")
		log.Println("更新后的记录:", user2)

		//复杂查询
		db.Where("name = ?", "xxb").Where("age = ?", 18).Find(&users1)
		db.Where("name = ?", "biaoge").Or("age = ?", 18).Find(&users2)
		db.Where("name = ?", "John").Or(db.Where("name = ?", "biaoge").Where("age = ?", 99999)).Find(&users3)
		log.Println("users1：", users1)
		log.Println("users2：", users2)
		log.Println("users3：", users3)

		// 删除记录
		db.Delete(&user)

		// 查看全部记录
		log.Println("查看全部记录:")

		db.Find(&users)
		log.Println(users)
		// end

		errSend := conn.WriteMessage(messageType, []byte("Hello, world!"))
		if errSend != nil {
			log.Println("发送数据失败：", errSend)
			return
		}
	}
}

func mainExit() {
	// 程序结束前关闭数据库连接
	if err := db.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}
	// 确保日志文件在程序结束时关闭
	// 注意：因为 init() 函数没有直接的方法来在程序结束时执行代码，
	// 通常需要在main()中或程序的其他适当位置显式调用关闭文件的操作。
	if err := logFile.Close(); err != nil {
		log.Printf("关闭logFile失败: %v", err)
	}
}
