package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

func NewConnection(path string) (*sql.DB, error) {
	// Параметры строки подключения (DSN) оптимизированные под SSD:
	// 1. journal_mode=WAL — параллельное чтение не блокирует запись и наоборот.
	// 2. synchronous=FULL — КРИТИЧНО ДЛЯ SSD: гарантирует 100% сохранность данных при сбое питания.
	// 3. busy_timeout=5000 — если база занята записью, горутина ждет до 5 секунд перед ошибкой.
	// 4. foreign_keys=ON — принудительное включение поддержки внешних ключей.
	dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)&_pragma=synchronous(FULL)&_pragma=busy_timeout(500)&_pragma=foreign_keys=ON", path)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия базы данных: %w", err)
	}

	// Настройка пула соединений встроенного менеджера sql.DB под SSD
	// На SSD запись происходит мгновенно, поэтому пул можно делать более "агрессивным"
	db.SetMaxOpenConns(100)                 // Максимальное количество одновременно открытых соединений
	db.SetMaxIdleConns(20)                  // Количество неактивных соединений, удерживаемых в памяти
	db.SetConnMaxLifetime(15 * time.Minute) // Время жизни одного соединения в пуле
	db.SetConnMaxIdleTime(5 * time.Minute)  // Время, через которое неиспользуемое соединение закроется

	// Проверяем соединение с диском
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("база данных недоступна: %w", err)
	}
	return db, nil
}
