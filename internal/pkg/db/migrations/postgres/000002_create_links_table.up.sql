BEGIN;

CREATE TABLE IF NOT EXISTS Links(
    ID serial,
    Title VARCHAR (255),
    Address VARCHAR (255),
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    PRIMARY KEY (ID)
);

COMMIT;