CREATE TABLE actors (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(50),
    LastName VARCHAR(50),
    Gender VARCHAR(10),
    DateOfBirth DATE,
);

-- Создаем таблицу для фильмов
CREATE TABLE films (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(150),
    Description VARCHAR(1000),
    PublishDay DATE,
    Rating INT
);

-- Создаем таблицу для связи между актерами и фильмами
CREATE TABLE actor_film (
    actorID INT,
    FilmID INT,
    FOREIGN KEY (actorID) REFERENCES actors(ID),
    FOREIGN KEY (FilmID) REFERENCES films(ID),
    PRIMARY KEY (actorID, FilmID)
);