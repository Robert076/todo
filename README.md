# RESTful API for a to-do list (standard Go library)

This API follows the principles of *REST Architecture* while using only the standard Go library. Nothing external. In order to use this:

*Clone the repo*
>git clone https://github.com/Robert076/todo

*Run the main.go file*
>cd main
>go run main.go

*Play around with the requests*
>curl -X GET localhost:8080/todos
>curl -X GET localhost:8080/todos/*{id}*

>curl -X POST localhost:8080/todos/post?title=*yourTitle*&description=*yourDescription*

>curl -X PUT localhost:8080/todos/put?id=*existingId*title=*newTitle*&description=*newDescription*

>curl -X DELETE localhost:8080/todos/delete?id=*idToDelete*

There is no hosted database implemented yet. It is all local.
