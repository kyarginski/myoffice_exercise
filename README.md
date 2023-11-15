Тестовое задание Go (myoffice)
====

Необходимо реализовать CLI-утилиту, которая реализует асинхронную обработку входящих URL из файла, переданного в качестве аргумента данной утилите.
Формат входного файла: на каждой строке – один URL. URL может быть очень много! Но могут быть и невалидные URL.

Пример входного файла:

https://myoffice.ru

https://yandex.ru

По каждому URL получить контент и вывести в консоль его размер и время обработки. Предусмотреть обработку ошибок.

# Решение
-----

## Запуск

```shell
go run ./cmd/parser -source=source/input_urls.txt -parallel=0
```

Параметры запуска:
- source - путь к файлу с URLs
- parallel - количество параллельных горутин для обработки URLs, по умолчанию 0 = количество CPU (на MacBook Pro 2022 = 12)

## Результат

```shell
Start with 10 parallel working
URL: 12123 Size: 0 bytes Time: 1.625µs Error: invalid URL
URL: www.ya.ru Size: 0 bytes Time: 1.625µs Error: invalid URL
URL: https://yandex.ru Size: 12622 bytes Time: 218.675167ms Error:
URL: https://www.facebook.com Size: 68936 bytes Time: 333.537458ms Error:
URL: https://www.wikipedia.org Size: 89032 bytes Time: 375.882125ms Error:
URL: www.google.com Size: 0 bytes Time: 7.583µs Error: invalid URL
URL: https://www.youtube.com Size: 395733 bytes Time: 378.0265ms Error:
URL: htps://myoffice.ru/ Size: 0 bytes Time: 38.208µs Error: error fetching htps://myoffice.ru/: Get "htps://myoffice.ru/": unsupported protocol scheme "htps"
URL: https://www.medium.com Size: 122068 bytes Time: 277.645708ms Error:
URL: https://myoffice.comm/ Size: 0 bytes Time: 1.897125ms Error: error fetching https://myoffice.comm/: Get "https://myoffice.comm/": dial tcp: lookup myoffice.comm: no such host
URL: https://myoffice.ru Size: 200096 bytes Time: 553.1365ms Error:
URL: https://www.quora.com Size: 80 bytes Time: 243.727209ms Error:
URL: https://ya.ru/ Size: 12594 bytes Time: 204.76325ms Error:
URL: https://www.reddit.com Size: 489808 bytes Time: 668.831042ms Error:
URL: https://www.instagram.com Size: 351635 bytes Time: 798.590167ms Error:
URL: https://www.stackoverflow.com Size: 178069 bytes Time: 885.476375ms Error:
URL: https://www.linkedin.com Size: 129539 bytes Time: 890.860542ms Error:
URL: https://www.twitter.com Size: 171123 bytes Time: 1.343475375s Error:
URL: http://myoffice.com/ Size: 60076 bytes Time: 1.346923625s Error:

Total processing time: 1.725675208s
Total count URLs: 19
```

Отображается информация по каждому URL:
- URL - адрес
- Size - размер контента в байтах
- Time - время обработки
- Error - ошибка, если есть

В конце выводится общее время обработки и количество обработанных URL.

Также для ограничения времени выполнения запросов используется таймаут 3 секунды (прописан в коде, в случае необходимости может быть вынесен в качестве параметра).