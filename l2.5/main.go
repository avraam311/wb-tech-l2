package main

// объявляем структура с полем msg string
type customError struct {
	msg string
}

// добавляем этой структуре метод Error, который возвращает string
func (e *customError) Error() string {
	return e.msg
}

// функия test, которая возвращает указатель на структуру customError
// но по факту возвращает nil типа *costomError
func test() *customError {
	// ... do something
	return nil
}

func main() {
	// объявляем переменную err типа error
	// error - это интерфейс, значит его статический тип - error
	var err error
	// присваиваем туда результат работы функции test
	// это будет его динамическим типом
	err = test()
	// если err не равен nil, печатаем "error" и завершаем основную горутину
	if err != nil {
		println("error")
		return
	// если err равен nil, выводим "ok"
	}
	println("ok")

	// в выводе мы получим только "error", так как интерфейс равен nil,
	// только если его и статический, и динамический типы равны nil
	// а у нас только динамический тип равен nil
}
