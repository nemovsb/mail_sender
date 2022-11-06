# mail_sender

Задание:
Написать на Go небольшой сервис отправки имейл рассылок.
Возможности сервиса:
 1. Отправка рассылок с использованием html макета и списка подписчиков.
    > Реализовоно через метод /send
 2. Отправка отложенных рассылок.
    > Реализовано через метод /send/defer
 3. Использование переменных в макете рассылки. (Пример: имя, фамилия, день рождения из списка подписчиков)
    > Доступные для использования в шаблоне переменные:

        {{ .Name}}, 
        {{ .Surname}}, 
        {{ .Birthday}}
 4. Отслеживание открытий писем.
    > Реализовано через метод /track с использованием пикселя со ссылкой на данный метод.





## Config example
    {
        "http_server": {
            "port": "8081"
        },
        "sender": {
            "from": "nemov.test.sb1@gmail.com",
            "password": "***",
            "smtphost": "smtp.gmail.com",
            "smtpport": "587"
        }
    }

## Request examples

### Send mails:
    curl localhost:8081/send -X POST -d '{"mailingsendid": "1234321", "mails":["example1@gmail.com"],"templateid": 1}'

### Send Email with Sending Time:
    curl localhost:8081/send/defer -X POST -d '{"exectime" : "2022-10-07T13:25:05Z", "mailingsendid": "1234321", "mails":["example1@gmail.com"],"templateid": 1}'

### Send Email with Sending Time:
    curl localhost:8081/send/defer -X POST -d '{"exectime" : "2022-11-08T16:00:05Z", "mailingsendid":"1234321", "mails":["example2@gmail.com"],"templateid": 0}'

### Create new recipients
    curl localhost:8081/recipient/create -X POST -d '{"recipients": [{"mailaddr": "example1@gmail.com",	"name": "TestName1",	"surname": "TestSurname1",	"birthday": "13.04.1991"}, {"mailaddr": "example2@gmail.com","name": "TestName2", "surname": "TestSurname2","birthday": "20.12.1965"}]}'

### Get all saved recipients
    curl localhost:8081/recipient -X GET

### Get all stored templates
    curl localhost:8081/template -X GET

### Add a new template to storage
    curl localhost:8081/template/create -X POST -d '{"template": "some html template"}'