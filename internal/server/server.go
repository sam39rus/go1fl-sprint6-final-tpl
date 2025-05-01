package server

import (
	"log"
	"net/http"
	"time"

	handlerPkg "go1fl-sprint6-final-tpl/internal/handlers"
)

// Server структура для хранения конфигурации сервера
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// CreateServer создает новый сервер с настройками роутинга и конфигурирует http.Server.
// Принятый аргумент — логгер (*log.Logger), который передается в конфигурацию сервера.
func CreateServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	// Регистрация хендлеров в роутере
	router.Handle("/", http.HandlerFunc(handlerPkg.ServeIndexPage))
	router.Handle("/upload", http.HandlerFunc(handlerPkg.UploadHandler))

	serverConfig := &http.Server{
		Addr:         ":8080",          // Адрес прослушивания порта
		Handler:      router,           // Роутер обработчиков
		ErrorLog:     logger,           // Журнал ошибок
		ReadTimeout:  5 * time.Second,  // Таймаут чтения
		WriteTimeout: 10 * time.Second, // Таймаут записи
		IdleTimeout:  15 * time.Second, // Таймаут неактивности соединения
	}

	return &Server{
		Logger: logger,
		Server: serverConfig,
	}
}

// Start запускает сервер и слушает указанный порт.
// Принимает объект Server и возвращает ошибку в случае неудачи.
func (srv *Server) Start() error {
	srv.Logger.Println("Запуск сервера...")
	return srv.Server.ListenAndServe()
}
