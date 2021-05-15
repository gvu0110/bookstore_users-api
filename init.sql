CREATE DATABASE IF NOT EXISTS users_db;
USE users_db;
CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    first_name NVARCHAR(45) DEFAULT '',
    last_name NVARCHAR(45) DEFAULT '',
    email VARCHAR(100) UNIQUE NOT NULL,
    date_created DATETIME NOT NULL,
    status VARCHAR(45) NOT NULL,
    password VARCHAR(100) NOT NULL
);
INSERT INTO users (
    first_name,
    last_name,
    email,
    date_created,
    status,
    password
) VALUES (
    "Adam",
    "Vu",
    "adam.vu@gmail.com",
    "2006-01-02 15:04:05",
    "active",
    "e807f1fcf82d132f9bb018ca6738a19f" /* password: 1234567890 */
);
INSERT INTO users (
    first_name,
    last_name,
    email,
    date_created,
    status,
    password
) VALUES (
    "Katie",
    "Do",
    "katie.do@gmail.com",
    "2007-01-02 15:04:05",
    "active",
    "202cb962ac59075b964b07152d234b70" /* password: 123 */
);
INSERT INTO users (
    first_name,
    last_name,
    email,
    date_created,
    status,
    password
) VALUES (
    "Mikhaila",
    "Santos",
    "mikhaila.santos@gmail.com",
    "2008-01-02 15:04:05",
    "active",
    "900150983cd24fb0d6963f7d28e17f72" /* password: abc */
);