-- Целевой запрос:
SELECT
    teachers.surname,
    teachers.name,
    streams.number,
    grades.performance
FROM grades
JOIN teachers ON grades.teacher_id = teachers.id
JOIN streams ON grades.stream_id = streams.id
WHERE streams.number >= 200;

-- Я предполагаю, запрос ускорит следующий индекс:
CREATE INDEX 'streams_number_idx' ON streams(number);

-- Я не уверен, но возможно, запрос ускорят также индексы:
CREATE INDEX 'grades_teacher_id_idx' ON grades(teacher_id);
CREATE INDEX 'grades_stream_id_idx' ON grades(stream_id);
