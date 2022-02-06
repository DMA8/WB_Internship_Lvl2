package main


import "../pkg"
/*
Цепочка обязанностей — это поведенческий паттерн,
позволяющий передавать запрос по цепочке потенциальных обработчиков,
 пока один из них не обработает запрос.
*/


// Пример - чат техподдержки. Сначала отвечает бот, если он не решил вопрос, он передается оператору,
//если оператор не решил вопрос - он передает вашу проблему другому профильному спецу
/*
Как и многие другие поведенческие паттерны, Цепочка обязанностей базируется на том, чтобы превратить отдельные поведения в объекты. 
В нашем случае каждая проверка переедет в отдельный класс с единственным методом выполнения. 
Данные запроса, над которым происходит проверка,будут передаваться в метод как аргументы.
*/
func main(){
	client := &pkg.Client{}
	operator := &pkg.Operator{}
	bot := pkg.ChatBot{}
	specialist := &pkg.Specialist{}
	bot.Next = operator
	operator.Next = specialist
	bot.Execute(client)
}