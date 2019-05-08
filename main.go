package main

var config = Config{}

func init() {
	config.Read()
}

type Todo struct {
	Id int32
	Name string
	Total int32
	Complete int32
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
				{Id:1, Name:"Rakesh ",Total:10,Complete:8,Pending:2,Done:false},
				{Id:2, Name:"Kunal",Total:11,Complete:11,Pending:0,Done:false},
				{Id:3, Name:"Dolly",Total:12,Complete:0,Pending:12,Done:false},
				{Id:4, Name:"Ram ",Total:15,Complete:15,Pending:0,Done:false},
				{Id:5, Name:"Ramesh ",Total:17,Complete:15,Pending:2,Done:false},
				{Id:6, Name:"Amit ",Total:30,Complete:28,Pending:2,Done:false},
				{Id:7, Name:"Aman ",Total:40,Complete:32,Pending:8,Done:false},
				
			},
		}

	r.Send("templates/template.html", data)
}
