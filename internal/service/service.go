package service

import (
	"errors"
	"strings"

	morsePkg "go1fl-sprint6-final-tpl/pkg/morse"
)

// Convert автоматически определяет тип строки (текст или код Морзе) и выполняет требуемое преобразование.
// В случае невозможности интерпретации данных возвращает ошибку.
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
	return strings.ContainsOnly(s, ".- ") // проверка, состоят ли символы строки только из разрешенных
}

// convertFromMorse преобразует строку из кода Морзе в текст.
// Если не удалось декодировать строку, возвращается ошибка.
func convertFromMorse(morse string) (string, error) {
	result, err := morsePkg.ToText(morse)
	if err != nil {
		return "", errors.New("не удалось расшифровать код Морзе: " + err.Error())
	}
	return result, nil
}

// convertToMorse преобразует строку из текста в код Морзе.
// Пока пакет morse не возвращает ошибок, но мы предусмотрели случай возможного изменения поведения.
func convertToMorse(text string) (string, error) {
	result := morsePkg.ToMorse(text)
	if result == "" {
		return "", errors.New("не удалось преобразовать текст в код Морзе")
	}
	return result, nil
}
