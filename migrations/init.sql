CREATE TABLE IF NOT EXISTS article (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    title TEXT DEFAULT NULL,
    url TEXT NOT NULL,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    readAt DATETIME DEFAULT NULL,
    FOREIGN KEY (userId) REFERENCES user (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    tgId INTEGER NOT NULL UNIQUE,
    tgUsername TEXT DEFAULT NULL UNIQUE,
    desc TEXT DEFAULT NULL,
    public BOOLEAN NOT NULL DEFAULT FALSE,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
