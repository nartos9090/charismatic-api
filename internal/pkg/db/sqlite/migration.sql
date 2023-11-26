-- SQLite
CREATE TABLE IF NOT EXISTS admins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    fullname TEXT NOT NULL,
    passwd TEXT NOT NULL,
    passwdSalt TEXT NOT NULL
);

INSERT INTO admins (id, email, fullname, passwd, passwdSalt)
VALUES (1, 'admin@mail.com', 'admin-san', '1qHPozUB22DOgjo7x6yNbSqtoSqDlNgxQN6eCj6Y+ZnhNZHpKfa9W6stCGuWzWCrKdCZ7t5wTDXKxqacggU91A==', 'Xr2TrTXhdY5CSCAf');

CREATE TABLE IF NOT EXISTS employees (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fullname TEXT NOT NULL,
    leaveQuota INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS employee_leave_submissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    employeeId INTEGER NOT NULL,
    leaveDate TEXT NOT NULL
);
