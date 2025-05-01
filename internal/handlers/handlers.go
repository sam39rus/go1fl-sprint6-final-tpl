package handlers

import (
	"bufio"
	"fmt"
	servicePkg "go1fl-sprint6-final-tpl/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ServeIndexPage отправляет HTML-файл index.html на корень сайта (/).
func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

// UploadHandler обслуживает запросы на /upload.
// Получает файл из формы, обрабатывает его, сохраняет результат в новый файл и возвращает результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST-методы поддерживаются.", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму без ограничений по размеру файла
	if err := r.ParseMultipartForm(0); err != nil {
		http.Error(w, "Ошибка при разборе формы.", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Чтение файла
	buf := bufio.NewReader(file)
	data, err := io.ReadAll(buf)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла.", http.StatusInternalServerError)
		return
	}

	// Преобразуем данные с помощью пакета service
	convertedData, err := servicePkg.Convert(string(data))
	if err != nil {
		http.Error(w, "Ошибка при преобразовании данных.", http.StatusInternalServerError)
		return
	}

	// Генерируем уникальное имя файла на основе времени UTC
	now := time.Now().UTC()
	filename := now.Format("20060102_150405") + filepath.Ext(header.Filename)
	outputPath := "./uploads/" + filename

	// Создаем директорию ./uploads/, если её нет
	err = os.MkdirAll("./uploads/", 0755)
	if err != nil {
		http.Error(w, "Ошибка при создании директории.", http.StatusInternalServerError)
		return
	}

	// Открываем и записываем файл
	fout, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Ошибка при создании выходного файла.", http.StatusInternalServerError)
		return
	}
	defer fout.Close()

	// Записываем результат в файл
	if _, err := fout.WriteString(convertedData); err != nil {
		http.Error(w, "Ошибка при записи результата.", http.StatusInternalServerError)
		return
	}

	// Формируем ответ пользователю
	response := fmt.Sprintf("Данные успешно обработаны.\nИсходный файл: %s\nРезультат: %s\nКонвертированные данные сохранены в: %s", header.Filename, convertedData, outputPath)

	fmt.Fprintln(w, response)
}
