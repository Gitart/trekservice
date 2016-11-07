REM Сервис
REM Savchenko Arthur
REM 07-09.2015 

@echo off
SETLOCAL
:: start

rem Path to current Programm API Service
SET GOPATH=%CD%
SET GOBIN=%CD%\BIN

SET IPPORT=http://10.10.10.10:7777/

rem Частота проверки сервиса
SET INTERVAL=10

rem Сервисы для контроля
SET SREVICES=10.10.10.10:5555;195.168.10.10:5551;195.168.0.1:5557
SET NOTIFY=Service is down 

rem путь к компилятору
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
cls

title Run "GO" 
color 0f

rem Cтарт компиляция и  запуск сервиса
rem go build -o trek.exe 
trek.exe >> log.txt

@pause
