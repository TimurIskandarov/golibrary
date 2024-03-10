# golibrary
База данных библиотеки как пример применения супер сервисов и фасадов с использованием принципов чистой архитектуры


# go-swagger

## Установка из бинарных дистрибутивов

Релизы [go-swagger](https://goswagger.io/install.html) распространяются в виде бинарных файлов, созданных на основе подписанных тегов. Он публикуется как [github релиз](https://github.com/go-swagger/go-swagger/tags), rpm, deb и docker образ.

## Docker образ
Сначала возьмите образ:
```sh
docker pull quay.io/goswagger/swagger
```
### Для пользователей Windows (msys2):
```sh
notepad /c/msys64/home/host/.bashrc
alias swagger='docker run --rm -it --env GOPATH=/go -v $USERPROFILE/Desktop:/go/c/Users/Host/Desktop -w /go$(pwd) quay.io/goswagger/swagger'
```
Пояснения к алиасу:
> - в контейнер встроен каталог `/go`
> - если проекты хранятся на рабочем столе, то монтируем `$USERPROFILE/Desktop`
> - в оболочке `msys2` это `/c/Users/Host/Desktop` (выполните `pwd` в терминале)

### Запуск
Успешность установки можно проверить командой 
```sh
swagger version
```
Запуск:
```sh
export MSYS2_ARG_CONV_EXCL="*"
swagger generate spec -o ./public/swagger.json --scan-models
```
Пояснения:
При вызове нативных исполняемых файлов из контекста `Cygwin` все аргументы, которые выглядят как пути `unix`, автоматически преобразуются в пути `Windows`. Так как все контейнеры на образе линукса, то необходимо отключить автоматическое преобразование путей. 
Подробнее: https://www.msys2.org/docs/filesystem-paths/#automatic-unix-windows-path-conversion

### Для пользователей Linux:
