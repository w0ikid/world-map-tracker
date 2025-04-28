# 🌍 World Map Tracker

**World Map Tracker** — это интерактивная платформа для путешественников, которая позволяет визуализировать посещенные страны и города, создавать персонализированные карты путешествий и делиться своими приключениями с ИИ.

---

## 📌 Технологии

- **Frontend**: Angular
- **Backend**: Golang + Gin
- **База данных**: PostgreSQL
- **Аутентификация**: Cookie-сессии
- **AI-помощник**: Groq AI API (GET API)[https://console.groq.com/home] Очень быстрый и бесплатный 

---

## 🚀 Как запустить проект локально

### 1. Подготовить базу данных (PostgreSQL)

Убедитесь, что у вас установлен и запущен PostgreSQL.  
Создайте базу данных:

```bash
createdb world_map_tracker
```

Или через psql:

```sql
CREATE DATABASE world_map_tracker;
```

(Не забудьте настроить переменные окружения для подключения к базе.)

---

### 2. Запуск Backend

Перейдите в папку backend:

```bash
cd backend
```

Скачайте зависимости:

```bash
go mod tidy
```

Настройте `.env` файл с переменными окружения:

**Вы можете не париться у меня .env все равно открытый. Можете сразу подключиться к удаленному БД :>**

```env
GROQ_API_KEY=
HOSTdb=
PORTdb=
USERdb=
PASSWORDdb=
DBNamedb=
SSL_MODEdb=
PORTHTTP=
```

Запустите миграции базы данных (если есть миграционный скрипт) или настройте схемы через код.

Запустите сервер:

```bash
go run main.go
```

---

### 3. Запуск Frontend

Перейдите в папку frontend:

```bash
cd frontend
```

Установите зависимости:

```bash
npm install
```

Настройте окружение, если нужно (например, файл `environment.ts`):

```typescript
export const environment = {
  production: false,
  apiUrl: 'http://localhost:1488' // адрес бекенда <---- тут исправьте {порт}
};
```

Запустите фронтенд:

```bash
ng serve
```

**Помните что в CORS у меня указан несколько адресов**
**Если вы хотите перейти на другой адрес то **

```bash
AllowOrigins:     []string{
			"http://localhost:4200",
			"https://world-map-tracker-nine.vercel.app",
			"https://world-map-tracker-nine.vercel.app/",
      },
```

backend/internal/app/start/http.go <- тут 

Фронтенд будет доступен на:  
👉 `http://localhost:4200`

---

## 📈 Планы на развитие

- Добавление социальных функций (друзья, лента путешествий)
- Интеграция с сервисами бронирования (авиабилеты, отели)
- Более детальная визуализация маршрутов и статистика по странам

---

## 🤝 Контакты

Если есть вопросы: **w0ikid** telegram

