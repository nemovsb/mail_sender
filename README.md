# mail_sender

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
curl localhost:8081/send -X POST -d '{"mails":["example@mail.ru"],"templateid": 2}'

curl localhost:8081/recipient/create -X POST -d '{"recipients": [{"mailaddr": "example1@gmail.com",	"name": "TestName1",	"surname": "TestSurname1",	"birthday": "13.04.1991"}, {"mailaddr": "example2@gmail.com","name": "TestName2", "surname": "TestSurname2","birthday": "20.12.1965"}]}'

curl localhost:8081/recipient -X GET

curl localhost:8081/template -X GET

curl localhost:8081/template/create -X POST -d '{"template": "some html template"}'