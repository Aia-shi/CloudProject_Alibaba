-- Table: periods
CREATE TABLE IF NOT EXISTS periods (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(256) NOT NULL DEFAULT 'no name',
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: expenses
CREATE TABLE IF NOT EXISTS expenses (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    period_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    title VARCHAR(256) NOT NULL DEFAULT 'no title',
    amount FLOAT NOT NULL DEFAULT 0,
    description VARCHAR(512) NOT NULL,
    date VARCHAR(256) NOT NULL DEFAULT '01-01-2026',
    status VARCHAR(256) NOT NULL DEFAULT 'zaplanowany',
    category VARCHAR(256) NOT NULL DEFAULT 'rachunki',
    FOREIGN KEY (period_id) REFERENCES periods(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: incomes
CREATE TABLE IF NOT EXISTS incomes (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    period_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    title VARCHAR(256) NOT NULL DEFAULT 'no title',
    amount FLOAT NOT NULL DEFAULT 0,
    description VARCHAR(512) NOT NULL,
    date VARCHAR(256) NOT NULL DEFAULT '01-01-2026',
    category VARCHAR(256) NOT NULL DEFAULT 'praca',
    FOREIGN KEY (period_id) REFERENCES periods(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
