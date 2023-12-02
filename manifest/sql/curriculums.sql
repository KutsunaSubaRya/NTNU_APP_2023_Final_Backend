CREATE TABLE curriculums
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INTEGER        NOT NULL,
    course_id  INTEGER UNIQUE NOT NULL,
    created_at TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (course_id) REFERENCES courses (id)
);

CREATE Table courses
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    number   TEXT NOT NULL UNIQUE, -- 開課序號
    code     TEXT NOT NULL,        -- 科目代碼
    name     TEXT NOT NULL,        -- 課程名稱
    teacher  TEXT NOT NULL,        -- 教師
    time     TEXT NOT NULL,        -- 時間
    location TEXT NOT NULL         -- 地點
);