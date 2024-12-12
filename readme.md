Вот улучшенный текст, с учётом всех предложений:

---

# TradeCLI: Торговый симулятор в терминале

**TradeCLI** — это консольное приложение для безопасной симуляции торговли фьючерсами и обучения основам трейдинга. Проект создан для экспериментов с торговыми стратегиями и анализа данных в реальном времени.

---

## Возможности приложения

- **Получение котировок**: Данные о текущих ценах акций и фьючерсов.
- **Торговые операции**: Покупка и продажа активов с обновлением портфеля.
- **Управление портфелем**: Просмотр текущих позиций, расчет прибыли/убытков.
- **Аналитика**: ASCII-графики для анализа трендов и отображения статистики.
- **Безрисковая среда**: Полный контроль над торговыми симуляциями без реальных вложений.

---

## Установка

### 1. Убедитесь, что установлен Go
Загрузите и установите [Go](https://go.dev/dl/), если он ещё не установлен.

### 2. Склонируйте проект
```bash
git clone https://github.com/karim/TradeCLI.git
cd TradeCLI
```

### 3. Установите зависимости
```bash
go mod tidy
```

### 4. Забилдите приложение
```bash
go build main.go
```

---

## Как пользоваться

Для отображения всех доступных команд выполните:
```bash
TradeCLI -h
```

Основные команды:

- **`buy [symbol] [quantity]`** — Покупка указанного количества акций/фьючерсов.  
  Пример:  
  ```bash
  TradeCLI buy SBERF 5
  ```
  Результат:  
  ```
  Куплено: 5 фьючерсов на SBER.
  ```

- **`sell [symbol] [quantity]`** — Продажа указанного количества активов.  
  Пример:  
  ```bash
  TradeCLI sell SBERF 3
  ```
  Результат:  
  ```
  Продано: 3 фьючерса на SBER.
  ```

- **`portfolio`** — Отображение текущего портфеля и результатов торговли.  
  Пример:  
  ```bash
  TradeCLI portfolio
  ```
  Результат:  
  ```
  Портфель:
  - SBERF: 2 шт. (21510 руб./шт.)
  Итоговая стоимость: 43020 руб.
  ```

- **`stats`** — Вывод аналитики и трендов.  
  Пример:  
  ```bash
  TradeCLI stats
  ```

---

## Преимущества

- Поддержка бесплатного API для получения актуальных данных.
- Локальное хранение данных в SQLite.
- Удобная настройка интерфейса через библиотеки `Cobra` и `Viper`.

---

## Планы на будущее

- Реализация многопоточной обработки данных.
- Интеграция в Telegram-бота для работы с портфелем.
- Расширение функционала на торговлю криптовалютами.
- Добавление возможностей для моделирования реальной биржевой активности.

---

## 🖋️ От автора

Я всегда открыт для ваших предложений и обратной связи. Если у вас есть идеи, пишите в Telegram: [@Aleeeeph](https://t.me/Aleeeeph).