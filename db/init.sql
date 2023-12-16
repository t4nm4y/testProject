-- Creating a table for movies
CREATE TABLE movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    language VARCHAR(50) NOT NULL,
    genre VARCHAR(50) NOT NULL
);

-- Inserting 5 random movies
INSERT INTO movies (name, language, genre) VALUES
    ('Movie1', 'English', 'Action'),
    ('Movie2', 'Hindi', 'Comedy'),
    ('Movie3', 'German', 'Drama'),
    ('Movie4', 'Hindi', 'Thriller'),
    ('Movie5', 'English', 'Science Fiction');
