package precheck

import "regexp"

func ValidateAccount(phone string) bool {
	// Регулярное выражение для проверки номера телефона
	// Пример: "+1234567890", "123-456-7890", "(123) 456-7890", "1234567890"
	phoneRegex := `^\+?[1-9]\d{1,14}$`

	// Компилируем регулярное выражение
	re := regexp.MustCompile(phoneRegex)

	// Проверяем соответствие строки шаблону
	return re.MatchString(phone)
}
