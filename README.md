# Промежуточная аттестация модуля «Go-разработчик»
## Cities information service

### Сборка
go build cmd/main.go

### Пакеты
github.com/go-chi/chi/v5 - маршрутизатор для http

### Cервер
server: localhost

port: 8080

### Путь к сsv файлу с информацией о городах
path: "golang-cities-information-service/cities.csv"


### Примеры запросов к серверу


#### Получение информации о городе по его id

    GET localhost:8080/321

#### Добавление новой записи в список городов

    POST http://localhost:8080/create

```JSON
{
  "name": "test",
  "region": "test",
  "district": "test",
  "population": 1,
  "foundation": 1
}
```

#### Удаление информации о городе по указанному id

    DELETE localhost:8080/1
    

#### Обновление информации о численности населения города по указанному id

    PUT localhost:8080/population/1

```JSON
{
  "population": 1111111
}
```

#### Получение списка городов по указанному региону
    GET localhost:8080/region/Кемеровская%20область

#### Получение списка городов по указанному округу
    GET localhost:8080/district/Сибирский

#### Получения списка городов по указанному диапазону численности - населения
    GET localhost:8080/population/range

```JSON
{
  "start": 20000,
  "end": 900000
}
```

#### Получения списка городов по указанному диапазону года основания

    GET localhost:8080/foundation/range

```JSON
{
  "start" : 1200,
  "end" : 1600
}
```


### Пример сsv файла

829,Екатеринбург,Свердловская область,Уральский,1377738,1723
693,Пермь,Пермский край,Приволжский,1000679,1723
643,Омск,Омская область,Сибирский,1154000,1716
321,Новокузнецк,Кемеровская область,Сибирский,547885,1618
606,Нижний Новгород,Нижегородская область,Приволжский,1250615,1221
704,Владивосток,Приморский край,Дальневосточный,592069,1860
222,Иркутск,Иркутская область,Сибирский,587225,1661
121,Волгоград,Волгоградская область,Южный,1021244,1589
1009,Хабаровск,Хабаровский край,Дальневосточный,577668,1858
177,Махачкала,Дагестан,Северо-Кавказский,577990,1844
769,Самара,Самарская область,Приволжский,1164900,1586
771,Тольятти,Самарская область,Приволжский,719484,1737
1002,Ульяновск,Ульяновская область,Приволжский,613793,1648
490,Москва,Москва,Центральный,11514330,1147
380,Краснодар,Краснодарский край,Южный,744933,1793
653,Оренбург,Оренбургская область,Приволжский,570329,1743
744,Ростов-на-Дону,Ростовская область,Южный,1091544,1749
410,Красноярск,Красноярский край,Сибирский,1000000,1628
5,Барнаул,Алтайский край,Сибирский,612091,1730
634,Новосибирск,Новосибирская область,Сибирский,1498921,1893
1109,Ярославль,Ярославская область,Центральный,591486,1010
62,Уфа,Башкортостан,Приволжский,1062300,1574
781,Санкт-Петербург,Санкт-Петербург,Северо-Западный,4848742,1703
989,Тюмень,Тюменская область,Уральский,581758,1586
922,Казань,Татарстан,Приволжский,1143546,1005
159,Воронеж,Воронежская область,Центральный,889989,1586
797,Саратов,Саратовская область,Приволжский,836900,1590
315,Кемерово,Кемеровская область,Сибирский,532884,1918
1058,Челябинск,Челябинская область,Уральский,1130273,1736
993,Ижевск,Удмуртия,Приволжский,628117,1760


## Описание промежуточной аттестации

### Цель работы
Проверить и закрепить знания, полученные на курсе «Go-разработчик»: 

* основы синтаксиса языка;
* условные операторы и циклы;
* работа с файловой системой;
* структуры данных;
* сериализация;
* многопоточность;
* обмен данными по сети.


### Что нужно было сделать
Разработать сервис, предоставляющий информацию о городах. Данные хранятся в файле. В момент старта сервиса данные из файла кешируются в память, в момент завершения работы сервиса данные перезаписываются обратно в файл.


### В итоге создан сервис, имеющий следующий функционал:

* получение информации о городе по его id;
* добавление новой записи в список городов;
* удаление информации о городе по указанному id;
* обновление информации о численности населения города по указанному id;
* получение списка городов по указанному региону;
* получение списка городов по указанному округу;
* получения списка городов по указанному диапазону численности населения;
* получения списка городов по указанному диапазону года основания.


### В проекте реализовано:
* Безопасная работа с данными в памяти.
* Для завершения работы используется паттерн Graceful Shutdown.
* Заголовки HTTP-запросов соответствуют выполняемым операциям.
* Ответы содержат правильные коды HTTP-состояний.
* Для кодирования данных используется формат JSON.
