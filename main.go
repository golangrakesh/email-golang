package main

var config = Config{}

func init() {
	config.Read()
}

type Todo struct {
	Id int32
	Distributor string
	Total int32
	Punched int32
	Pending int32 
	Done  bool
}

type TodoPageData struct {
	Username string
	PageTitle string
	Todos     []Todo
}


func main() {
	subject := "Daily Report"
	receiver:= "rakesh.rai@pratham.net"
	r := NewRequest([]string{receiver}, subject)

	data := TodoPageData{
			Username:"Rakesh",
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Id:1, Distributor:"Rakesh ",Total:10,Punched:8,Pending:2,Done:false},
				{Id:2, Distributor:"Kunal",Total:11,Punched:11,Pending:0,Done:false},
				{Id:3, Distributor:"Dolly",Total:12,Punched:0,Pending:12,Done:false},
				{Id:4, Distributor:"Ram ",Total:15,Punched:15,Pending:0,Done:false},
				{Id:5, Distributor:"Ramesh ",Total:17,Punched:15,Pending:2,Done:false},
				{Id:6, Distributor:"Amit ",Total:30,Punched:28,Pending:2,Done:false},
				{Id:7, Distributor:"Aman ",Total:40,Punched:32,Pending:8,Done:false},
				
			},
		}

	r.Send("templates/template.html", data)
}