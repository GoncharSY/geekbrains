CREATE TABLE IF NOT EXISTS 'teachers' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL,
    'surname' TEXT NOT NULL,
    'email' text
);

CREATE TABLE IF NOT EXISTS 'courses' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS 'streams' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'number' INTEGER NOT NULL,
    'course_id' INTEGER NOT NULL,
    'start_date' TEXT,
    'student_amount' INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY ('course_id') REFERENCES 'cources'('id')    
);

CREATE TABLE IF NOT EXISTS 'average_grades' (
    'teacher_id' INTEGER NOT NULL,
    'stream_id' INTEGER NOT NULL,
    'grade' REAL NOT NULL DEFAULT 0,
    PRIMARY KEY ('teacher_id', 'stream_id'),
    FOREIGN KEY ('teacher_id') REFERENCES 'teachers' ('id'),
    FOREIGN KEY ('stream_id') REFERENCES 'streams' ('id')   
);
