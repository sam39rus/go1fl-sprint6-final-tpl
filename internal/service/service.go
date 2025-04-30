package service

import (
	"errors"
	"strings"

	morsePkg "go1fl-sprint6-final-tpl/pkg/morse"
)

// Convert автоматически определяет тип строки (текст или код Морзе) и выполняет требуемое преобразование.
// Возвращает ошибку, если данные нельзя корректно преобразовать.
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input) // удаляем начальные и финальные пробельные символы

	if isMorseCode(input) {
		return convertFromMorse(input)
	}
	return convertToMorse(input)
}

// isMorseCode проверяет, является ли строка кодом Морзе.
// Строка признаётся кодом Морзе, если содержит только символы ".", "-", и пробел.
func isMorseCode(s string) bool {
	return strings.IndexAny(s, "^.- ") == -1
}

// convertFromMorse преобразует строку из кода Морзе в текст.
// Возвращает ошибку, если полученный результат пуст.
func convertFromMorse(morse string) (string, error) {
	result := morsePkg.ToText(morse)
	if result == "" {
		return "", errors.New("не удалось расшифровать код Морзе")
	}
	return result, nil
}

// convertToMorse преобразует строку из текста в код Морзе.
// Возвращает ошибку, если полученный результат пуст.
func convertToMorse(text string) (string, error) {
	result := morsePkg.ToMorse(text)
	if result == "" {
		return "", errors.New("не удалось преобразовать текст в код Морзе")
	}
	return result, nil
}
