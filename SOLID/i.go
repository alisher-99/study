package SOLID

import "fmt"

// Принцип разделения интерфейса

type Writer interface {
	Write(document string)
}

type Reader interface {
	Read(document string)
}

type MultiFunctionDevice interface {
	Writer
	Reader
}

type SimpleWriter struct{}

func (p *SimpleWriter) Write(document string) {
	// логика печати
}

type FullWR struct{}

func (c *FullWR) Write(document string) {
	// ...
}

func (c *FullWR) Read(document string) {
	// ...
}

////////////////////////////////////////////////////////////////////////////////

type Runner interface {
	Run(document string)
}

type Walker interface {
	Walk(document string)
}

type MultiSportsman interface {
	Runner
	Walker
}

type SimpleRunner struct{}

func (p *SimpleRunner) Run(distance string) {
	// логика бега
}

type FullRW struct{}

func (c *FullRW) Run(distance string) {
	// ...
}

func (c *FullRW) Walk(document string) {
	// ...
}

////////////////////////////////////////////////////////////////////////////////

// Storable определяет интерфейс для устройств хранения данных
type Storable interface {
	Read() string
	Write(data string)
}

// Readable определяет интерфейс только для чтения данных из устройства
type Readable interface {
	Read() string
}

// Writable определяет интерфейс только для записи данных в устройство
type Writable interface {
	Write(data string)
}

// HDD представляет жесткий диск
type HDD struct{}

// Read читает данные с жесткого диска
func (h HDD) Read() string {
	return "Чтение данных с жесткого диска"
}

// Write записывает данные на жесткий диск
func (h HDD) Write(data string) {
	fmt.Printf("Запись данных '%s' на жесткий диск\n", data)
}

// SSD представляет твердотельный накопитель
type SSD struct{}

// Read читает данные с твердотельного накопителя
func (s SSD) Read() string {
	return "Чтение данных с твердотельного накопителя"
}

// Write записывает данные на твердотельный накопитель
func (s SSD) Write(data string) {
	fmt.Printf("Запись данных '%s' на твердотельный накопитель\n", data)
}

func main() {
	hdd := HDD{}
	ssd := SSD{}

	// Используем методы только для чтения
	readableDevices := []Readable{hdd, ssd}
	for _, device := range readableDevices {
		fmt.Println(device.Read())
	}

	// Используем методы только для записи
	writableDevices := []Writable{hdd, ssd}
	for _, device := range writableDevices {
		device.Write("Пример данных")
	}
}
