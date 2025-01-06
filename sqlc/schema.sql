CREATE TABLE Pages (
    PageURL TEXT PRIMARY KEY,
    PageId SERIAL UNIQUE,
    CommentsCount INTEGER NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Users (
    UserId SERIAL PRIMARY KEY,
    UserName TEXT NOT NULL,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    EmailId TEXT NOT NULL
);

CREATE TABLE Comments (
    CommentId SERIAL PRIMARY KEY,
    PageId INTEGER NOT NULL,
    UserId INTEGER NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    EditedBool BOOLEAN DEFAULT False,
    UpVotes INTEGER DEFAULT 0,
    DownVotes INTEGER DEFAULT 0,
    CommentData TEXT NOT NULL,
    ParentId INTEGER DEFAULT 0,
    FOREIGN KEY (PageId) REFERENCES Pages(PageId),
    FOREIGN KEY (UserId) REFERENCES Users(UserId)
);

CREATE TABLE UserRelations (
    ViewerUserId INTEGER NOT NULL,
    CommentorUserId INTEGER NOT NULL,
    Tag TEXT NOT NULL,
    Positivity BOOLEAN NOT NULL,
    PRIMARY KEY (ViewerUserId, CommentorUserId),
    FOREIGN KEY (ViewerUserId) REFERENCES Users(UserId),
    FOREIGN KEY (CommentorUserId) REFERENCES Users(UserId)
);1