# Ewallet-infotecs-intern-task
Ссылка на гитхаб - https://github.com/lty2107/Ewallet-infotecs-intern-task
собрать приложение - docker build -t ewallet-infotecs .
запустить - docker run -p 8080:8080 ewallet-infotecs 

API endpoints:

/api/send - для перевода денег
{
    "from":"",
    "to":"",
    "amount":""
}

/api/transactions?count=N - получение N последних переводов

/api/wallet/{address}/balance - узнать баланс кошелька с адресом {address}

При первом запуске приложения создается 10 кошельков 

Файл для запуска приложения без докера - app.go