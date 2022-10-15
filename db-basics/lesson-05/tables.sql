-- ================================================================================================
-- Таблица преподавателей =========================================================================
-- ================================================================================================

CREATE TABLE IF NOT EXISTS 'teachers' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL,
    'surname' TEXT NOT NULL,
    'email' text
);

-- ================================================================================================
-- Таблица курсов =================================================================================
-- ================================================================================================

CREATE TABLE IF NOT EXISTS 'courses' (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'name' TEXT NOT NULL
);

-- ================================================================================================
-- Таблица потоков ================================================================================
-- ================================================================================================

CREATE TABLE IF NOT EXISTS "streams_tmp" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "number" INTEGER NOT NULL,
    "course_id" INTEGER NOT NULL,
    "student_amount" INTEGER NOT NULL DEFAULT 0,
    "started_at" TEXT,
    "finished_at" TEXT,
    FOREIGN KEY ("course_id") REFERENCES "cources"("id")    
);

INSERT INTO "streams_tmp" (
    "id",
    "number",
    "course_id",
    "student_amount",
    "started_at",
    "finished_at"
) SELECT
    "id",
    "number",
    "course_id",
    "student_amount",
    "started_at",
    "finished_at"
FROM "streams";

-- ================================================================================================
-- Таблица оценок =================================================================================
-- ================================================================================================

CREATE TABLE IF NOT EXISTS 'grades_tmp' (
    'teacher_id' INTEGER NOT NULL,
    'stream_id' INTEGER NOT NULL,
    'performance' REAL NOT NULL DEFAULT 0,
    PRIMARY KEY ('teacher_id', 'stream_id'),
    FOREIGN KEY ('teacher_id') REFERENCES 'teachers' ('id'),
    FOREIGN KEY ('stream_id') REFERENCES 'streams' ('id')   
);

INSERT INTO "grades_tmp" (
    "teacher_id",
    "stream_id",
    "performance"
)  SELECT
    "teacher_id",
    "stream_id",
    "performance"
FROM "grades";
