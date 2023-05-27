package main

import (
	"github.com/yosiopp/biletojy/data"
	"time"
	"fmt"
)

func main() {
	dao := data.NewDao()
	defer dao.Close()
	ticket := data.Ticket{Title: "こんにちは世界", Content: `
	# biletojy最初の課題です
	markdown記法が使えます。
	コードブロックにmermaidを指定すればmermaidも使用できます
	`, Tags: "test first_issue", CreatedBy: "yosiopp", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	data.AddTicket(dao, &ticket)
	fmt.Println(ticket)
}
