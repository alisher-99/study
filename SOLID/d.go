package SOLID

import "fmt"

// В этой реализации WebService зависит только от абстракции Notifier, а не от конкретной реализации.
// Теперь мы можем легко изменить способ отправки уведомлений, просто передавая другую реализацию Notifier
// при создании WebService, не изменяя кода WebService самого по себе.
// Это соблюдает принцип инверсии зависимостей (DIP).

// Интерфейс для отправки уведомлений
type Notifier interface {
	SendNotification(message string)
}

// Класс, отвечающий за отправку уведомлений через email
type EmailNotifier struct{}

// Реализация метода для отправки уведомления через email
func (en EmailNotifier) SendNotification(message string) {
	fmt.Println("Отправка уведомления по email:", message)
}

// Класс, представляющий веб-сервис
type WebService struct {
	notifier Notifier // Интерфейсный тип, а не конкретная реализация
}

// Метод для обработки запросов
func (ws WebService) HandleRequest() {
	// Используется абстракция Notifier, вместо конкретной реализации
	ws.notifier.SendNotification("Привет, мир!")
}

func main() {
	// Создание веб-сервиса с передачей конкретной реализации Notifier
	webService := WebService{
		notifier: EmailNotifier{},
	}
	webService.HandleRequest()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

// В этом примере UserController зависит только от интерфейса UserService, а не от конкретной реализации.
// Это позволяет нам легко заменять реализацию сервиса пользователей без изменения кода контроллера.

// Интерфейс для сервиса пользователей
type UserService interface {
	GetUserByID(id int) string
}

// Класс для реализации сервиса пользователей
type UserServiceImpl struct{}

// Реализация метода GetUserByID для UserServiceImpl
func (us UserServiceImpl) GetUserByID(id int) string {
	// Предположим, что здесь есть логика получения пользователя из базы данных
	return fmt.Sprintf("Пользователь с ID %d", id)
}

// Контроллер, обрабатывающий запросы пользователя
type UserController struct {
	UserService UserService // Интерфейсный тип, а не конкретная реализация
}

// Метод для получения пользователя по ID
func (uc UserController) GetUser(id int) string {
	return uc.UserService.GetUserByID(id)
}

func main() {
	// Создание контроллера с передачей конкретной реализации UserService
	userController := UserController{
		UserService: UserServiceImpl{},
	}

	// Получение пользователя по ID
	user := userController.GetUser(123)
	fmt.Println(user)
}

////////////////////////////////////////////////////////////////////////////////

// В этом примере Application зависит только от интерфейса MessagePrinter, а не от конкретных
// реализаций (например, ConsolePrinter или FilePrinter). Таким образом, мы можем легко заменить
// реализацию MessagePrinter на любую другую, не изменяя кода Application.
// Это демонстрирует принцип инверсии зависимостей (DIP).

// Интерфейс для сервиса печати сообщений
type MessagePrinter interface {
	Print(message string)
}

// Конкретная реализация для печати в консоль
type ConsolePrinter struct{}

// Реализация метода Print для печати в консоль
func (cp ConsolePrinter) Print(message string) {
	fmt.Println("Консоль:", message)
}

// Конкретная реализация для печати в файл
type FilePrinter struct{}

// Реализация метода Print для печати в файл
func (fp FilePrinter) Print(message string) {
	fmt.Println("Файл:", message)
}

// Класс приложения, зависящий от абстракции MessagePrinter
type Application struct {
	printer MessagePrinter // Интерфейсный тип, а не конкретная реализация
}

// Метод для печати сообщения
func (app Application) SendMessage(message string) {
	app.printer.Print(message)
}

func main() {
	// Создание экземпляра приложения с передачей конкретной реализации MessagePrinter
	application := Application{
		printer: ConsolePrinter{}, // Можно заменить на FilePrinter или любую другую реализацию MessagePrinter
	}

	// Печать сообщения
	application.SendMessage("Привет, мир!")
}
