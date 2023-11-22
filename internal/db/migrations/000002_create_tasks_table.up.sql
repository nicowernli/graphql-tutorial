CREATE TABLE IF NOT EXISTS Tasks(
    ID VARCHAR(32) PRIMARY KEY,
    Title VARCHAR (255) NOT NULL,
    Description TEXT,
    DueDate INTEGER,
    CreatedAt INTEGER NOT NULL,
    UpdatedAt INTEGER NOT NULL,
    RemovedAt INTEGER,
    INDEX(RemovedAt)
);