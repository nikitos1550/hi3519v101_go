```
https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/venc.go
это главный файл

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/venc-v3.go 
это чип запвисимая часть, там сейчас есть иницилизщация тестовых потоков

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/loop.go 
это общая часть по евент лупу выгребания данных

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/loop-query-v3.go 
это чип зависимая часть евент лупа, там функция получения данных

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/loop-query-callback.go 
это коллбек в голанг из си евент лупа

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/frame.go 
это абстракция обьекта кадра

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/frames.go 
а вот это набор кадров, который работает как циклический буффер

https://github.com/nikitos1550/hi3519v101_go/blob/testing/application/pkg/mpp/venc/encoder.go 
это не готово еще
```
