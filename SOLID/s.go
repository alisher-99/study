package SOLID

import (
	"fmt"
	"os"
)

// Logger - простой логгер
type Logger struct{}

// LogWriter - интерфейс для записи логов
type LogWriter interface {
	WriteLog(message string)
}

// ConsoleWriter - запись логов в консоль
type ConsoleWriter struct{}

// WriteLog записывает лог в консоль
func (cw ConsoleWriter) WriteLog(message string) {
	fmt.Println("Console:", message)
}

// FileWriter - запись логов в файл
type FileWriter struct {
	FileName string
}

// WriteLog записывает лог в файл
func (fw FileWriter) WriteLog(message string) {
	// Здесь реализация записи лога в файл
	fmt.Printf("File '%s': %s\n", fw.FileName, message)
}

func main() {
	logger := Logger{}

	consoleWriter := ConsoleWriter{}
	fileWriter := FileWriter{FileName: "app.log"}

	// Пример использования: логгирование в консоль
	logger.WriteLog(consoleWriter, "Пример сообщения для консоли")

	// Пример использования: логгирование в файл
	logger.WriteLog(fileWriter, "Пример сообщения для файла")
}

// WriteLog записывает лог с использованием указанного writer
func (l Logger) WriteLog(writer LogWriter, message string) {
	writer.WriteLog(message)
}

/////////////////////////////////////////////////////////////////////////////////

// FileManager отвечает за работу с файлами
type FileManager struct{}

// ReadFile считывает данные из файла
func (fm FileManager) ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// WriteFile записывает данные в файл
func (fm FileManager) WriteFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fileManager := FileManager{}

	// Пример использования: чтение данных из файла
	data, err := fileManager.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println("Данные из файла:", string(data))

	// Пример использования: запись данных в файл
	err = fileManager.WriteFile("output.txt", []byte("Пример записываемых данных"))
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
		return
	}
	fmt.Println("Данные успешно записаны в файл.")
}

/////////////////////////////////////////////////////////////////////////////////

// User отвечает за управление данными пользователя
type User struct{}

// Validate валидирует данные пользователя
func (u User) Validate(name, email string) error {
	if name == "" {
		return fmt.Errorf("имя пользователя не может быть пустым")
	}
	if email == "" {
		return fmt.Errorf("email пользователя не может быть пустым")
	}
	return nil
}

// Save сохраняет данные пользователя
func (u User) Save(name, email string) error {
	// Здесь реализация сохранения данных пользователя
	fmt.Printf("Данные пользователя '%s' (%s) сохранены\n", name, email)
	return nil
}

func main() {
	user := User{}

	// Пример использования: валидация данных пользователя
	err := user.Validate("John", "john@example.com")
	if err != nil {
		fmt.Println("Ошибка валидации данных пользователя:", err)
		return
	}

	// Пример использования: сохранение данных пользователя
	err = user.Save("John", "john@example.com")
	if err != nil {
		fmt.Println("Ошибка сохранения данных пользователя:", err)
		return
	}
}
