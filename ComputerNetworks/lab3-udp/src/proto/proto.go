package proto

import "encoding/json"

// Request -- запрос клиента к серверу.
type Request struct {
	// Поле Command может принимать три значения:
	// * "quit" - прощание с сервером (после этого сервер рвёт соединение);
	// * "enq" - добавление числа в очередь на сервере;
	// * "deq" - удаление числа с вершины очереди;
	// * "peek" - вывод числа с вершины очереди;
	// * "len" - длина очереди;
	Command string `json:"command"`

	// Если Command == "enq", в поле Data должно лежать целое число
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
	Ident string `json:"ident"`

}

// Response -- ответ сервера клиенту.
type Response struct {
	// Поле Status может принимать три значения:
	// * "ok" - успешное выполнение команды "quit" или "enq";
	// * "failed" - в процессе выполнения команды произошла ошибка;
	// * "result" - операция выполнена.
	Status string `json:"status"`

	// Если Status == "failed", то в поле Data находится сообщение об ошибке.
	// Если Status == "result", в поле Data должно лежать число ответа.
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
	Ident string `json:"ident"`
}
