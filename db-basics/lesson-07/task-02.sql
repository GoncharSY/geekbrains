-- Удаление данных о преподавателе:
BEGIN TRANSACTION;
    DELETE FROM grades WHERE teacher_id = 3;
    DELETE FROM teachers WHERE id = 3;
COMMIT;

-- Возврат данных о преподавателе:
BEGIN TRANSACTION;
    INSERT INTO teachers (id, name, surname, email)
    VALUES (3, 'Елена', 'Малышева', 'malisheva.e@google.com');
    INSERT INTO grades (teacher_id, stream_id, performance)
    VALUES (3, 1, 4.7), (3, 5, 0);
COMMIT;
