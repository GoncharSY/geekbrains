# Урок 5. Concurrency часть 2: основы типов данных из пакета `sync`

## Домашнее задание

1. Напишите программу, которая запускает **𝑛** потоков и дожидается завершения их всех.
2. Реализуйте функцию для разблокировки мьютекса с помощью `defer`.
3. Протестируйте производительность операций чтения и записи на множестве действительных чисел, безопасность которого обеспечивается `sync.Mutex` и `sync.RWMutex` для разных вариантов использования:

    - 10% запись, 90% чтение;
    - 50% запись, 50% чтение;
    - 90% запись, 10% чтение.