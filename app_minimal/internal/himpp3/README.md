Концепция API

Note: Проект подразумевает поддежрку широкого набора железа, с совершенно 
разным функционалом. Это относится и к разным моделям в рамках одного 
производителя, что уж говорить о разнице между чипами разных вендоров.
При реализации на сколько это возможно была сформирована обощенная
минимальная модель, работающая на всех платформах (в основном это касается
только функционала видеотракта).

1) Все описанные методы доступны для всех сборок проекта, не зависимо от
аппаратной платформы.

2) Метод может быть имплементирован, не имплементирован или не поддерживаться 
(подразумевается отсутсвие принципиальной возможности имплементации для некой 
аппаратной платформы).

3) Каждый отдельный метод подразумевает только одну модель поведения, принимает
на вход и отдает на выход строго описанные структуры данных. Никакой разницы 
в поведении метода в зависимости от аппаратной платформы выполнения нет 
(это требуется для качественного автоматического тестирования).

4) Некоторые платформы дают намного больше интересных возможностей, чем может 
быть учтено в минимальной обощенной модели. Поэтому существуют методы API 
имплементированные для подмножеств поддерживаемых аппаратных плафторм. Вполне 
законно существования отдельных методов реализующих сходный функционал, 
имплементированных для разных платформ. Такие методы мы будем называть 
расширенными.

5) Расширенные методы API могут изменять настройки примененные через обобщенные
методы. Это нужно учитывать при эксплуатации и настройке.

10) API, а именно его верхний HTTP уровень создается согласно принципам RESTful.

################################################################################3

/experimental/himpp3/venc/                  GET     -> LIST

/experimental/himpp3/venc/mjpeg             GET     -> LIST
/experimental/himpp3/venc/mjpeg/cbr         GET     -> LIST
/experimental/himpp3/venc/mjpeg/cbr         POST    -> CREATE
/experimental/himpp3/venc/mjpeg/cbr/{id}    GET     -> INFO
/experimental/himpp3/venc/mjpeg/cbr/{id}    PUT     -> UPDATE
/experimental/himpp3/venc/mjpeg/cbr/{id}    DELETE  -> DELETE
/experimental/himpp3/venc/mjpeg/vbr         GET     -> LIST
/experimental/himpp3/venc/mjpeg/vbr         POST    -> CREATE
/experimental/himpp3/venc/mjpeg/vbr/{id}    GET     -> INFO
/experimental/himpp3/venc/mjpeg/vbr/{id}    PUT     -> UPDATE
/experimental/himpp3/venc/mjpeg/vbr/{id}    DELETE  -> DELETE
/experimental/himpp3/venc/mjpeg/fixqp       GET     -> LIST
/experimental/himpp3/venc/mjpeg/fixqp       POST    -> CREATE
/experimental/himpp3/venc/mjpeg/fixqp/{id}  GET     -> INFO
/experimental/himpp3/venc/mjpeg/fixqp/{id}  PUT     -> UPDATE
/experimental/himpp3/venc/mjpeg/fixqp/{id}  DELETE  -> DELETE

/experimental/himpp3/venc/h264              GET     -> LIST
/experimental/himpp3/venc/h264/cbr          GET     -> LIST
/experimental/himpp3/venc/h264/cbr          POST    -> CREATE
/experimental/himpp3/venc/h264/cbr/{id}     GET     -> INFO
/experimental/himpp3/venc/h264/cbr/{id}     PUT     -> UPDATE
/experimental/himpp3/venc/h264/cbr/{id}     DELETE  -> DELETE
/experimental/himpp3/venc/h264/vbr          GET     -> LIST
/experimental/himpp3/venc/h264/vbr          POST    -> CREATE
/experimental/himpp3/venc/h264/vbr/{id}     GET     -> INFO
/experimental/himpp3/venc/h264/vbr/{id}     PUT     -> UPDATE
/experimental/himpp3/venc/h264/vbr/{id}     DELETE  -> DELETE
/experimental/himpp3/venc/h264/fixqp        GET     -> LIST
/experimental/himpp3/venc/h264/fixqp        POST    -> CREATE
/experimental/himpp3/venc/h264/fixqp/{id}   GET     -> INFO
/experimental/himpp3/venc/h264/fixqp/{id}   PUT     -> UPDATE
/experimental/himpp3/venc/h264/fixqp/{id}   DELETE  -> DELETE

/experimental/himpp3/venc/h265              GET     -> LIST
/experimental/himpp3/venc/h265/cbr          GET     -> LIST
/experimental/himpp3/venc/h265/cbr          POST    -> CREATE
/experimental/himpp3/venc/h265/cbr/{id}     GET     -> INFO
/experimental/himpp3/venc/h265/cbr/{id}     PUT     -> UPDATE
/experimental/himpp3/venc/h265/cbr/{id}     DELETE  -> DELETE
/experimental/himpp3/venc/h265/vbr          GET     -> LIST
/experimental/himpp3/venc/h265/vbr          POST    -> CREATE
/experimental/himpp3/venc/h265/vbr/{id}     GET     -> INFO
/experimental/himpp3/venc/h265/vbr/{id}     PUT     -> UPDATE
/experimental/himpp3/venc/h265/vbr/{id}     DELETE  -> DELETE
/experimental/himpp3/venc/h265/fixqp        GET     -> LIST
/experimental/himpp3/venc/h265/fixqp        POST    -> CREATE
/experimental/himpp3/venc/h265/fixqp/{id}   GET     -> INFO
/experimental/himpp3/venc/h265/fixqp/{id}   PUT     -> UPDATE
/experimental/himpp3/venc/h265/fixqp/{id}   DELETE  -> DELETE

Пример расширенных методов

/experimental/himpp3/venc/h265/vbr/{id}/advanced/hi3516av200    GET -> INFO
/experimental/himpp3/venc/h265/vbr/{id}/advanced/hi3516av200    PUT -> UPDATE
