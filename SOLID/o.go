package SOLID

import (
	"fmt"
	"math"
)

// Shape определяет общий интерфейс для всех геометрических фигур
type Shape interface {
	Area() float64
}

// Circle представляет круг
type Circle struct {
	Radius float64
}

// Area вычисляет площадь круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle представляет прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

// Area вычисляет площадь прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TotalArea вычисляет общую площадь для всех фигур
func TotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 3, Height: 4}

	shapes := []Shape{circle, rectangle}

	totalArea := TotalArea(shapes)
	fmt.Println("Общая площадь всех фигур:", totalArea)
}

//////////////////////////////////////////////////////////////////////

// PaymentMethod определяет интерфейс для обработки платежей
type PaymentMethod interface {
	ProcessPayment(amount float64) error
}

// Sberbank представляет платеж через Sberbank
type Sberbank struct{}

// ProcessPayment обрабатывает платеж через Sberbank
func (cc Sberbank) ProcessPayment(amount float64) error {
	fmt.Printf("Платеж на %.2f руб. обработан через Sberbank\n", amount)
	return nil
}

// PayPal представляет платеж через PayPal
type PayPal struct{}

// ProcessPayment обрабатывает платеж через PayPal
func (pp PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Платеж на %.2f руб. обработан через PayPal\n", amount)
	return nil
}

// PaymentProcessor обрабатывает платежи с использованием указанного метода оплаты
type PaymentProcessor struct {
	paymentMethod PaymentMethod
}

// NewPaymentProcessor создает новый экземпляр PaymentProcessor с указанным методом оплаты
func NewPaymentProcessor(paymentMethod PaymentMethod) *PaymentProcessor {
	return &PaymentProcessor{paymentMethod: paymentMethod}
}

// ProcessPayment обрабатывает платеж с указанной суммой
func (pp *PaymentProcessor) ProcessPayment(amount float64) error {
	return pp.paymentMethod.ProcessPayment(amount)
}

func main() {
	creditCardPaymentProcessor := NewPaymentProcessor(Sberbank{})
	creditCardPaymentProcessor.ProcessPayment(100.50)

	payPalPaymentProcessor := NewPaymentProcessor(PayPal{})
	payPalPaymentProcessor.ProcessPayment(200.75)
}

//////////////////////////////////////////////////////////////////////

// Notifier определяет интерфейс для всех способов отправки уведомлений
type Notifier interface {
	SendNotification(message string) error
}

// EmailNotifier представляет способ отправки уведомлений по электронной почте
type EmailNotifier struct{}

// SendNotification отправляет уведомление по электронной почте
func (en EmailNotifier) SendNotification(message string) error {
	fmt.Println("Отправлено уведомление по электронной почте:", message)
	return nil
}

// SMSNotifier представляет способ отправки уведомлений по SMS
type SMSNotifier struct{}

// SendNotification отправляет уведомление по SMS
func (sn SMSNotifier) SendNotification(message string) error {
	fmt.Println("Отправлено уведомление по SMS:", message)
	return nil
}

// NotificationService обрабатывает отправку уведомлений различными способами
type NotificationService struct {
	notifiers []Notifier
}

// NewNotificationService создает новый экземпляр NotificationService с указанными способами отправки уведомлений
func NewNotificationService(notifiers ...Notifier) *NotificationService {
	return &NotificationService{notifiers: notifiers}
}

// SendNotification отправляет уведомление всем способами
func (ns *NotificationService) SendNotification(message string) {
	for _, notifier := range ns.notifiers {
		notifier.SendNotification(message)
	}
}

func main() {
	// Создаем экземпляр NotificationService с двумя способами отправки уведомлений
	notificationService := NewNotificationService(EmailNotifier{}, SMSNotifier{})

	// Отправляем уведомление всем способами
	notificationService.SendNotification("Пример уведомления")
}
