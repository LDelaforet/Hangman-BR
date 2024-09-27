@echo off
set displayargs=false
set InstanceNumber=0
if "%1" == "exec" goto InstanceCalculator

:a
if "%1" == "display" set displayargs=true 
if "%4" == "display" set displayargs=true 
if "%1" == "exec" title HangmanBR Client[%2] - Instance numero %InstanceNumber% sur %3
if %displayargs% == true (
    echo ARGUMENTS: %*
    pause
)
if "%1" == "exec" goto exec

:build 
go build -o client.exe
set InstanceNumber=3
if %displayargs% == true (
    start test.bat exec %InstanceNumber% %InstanceNumber% display
) else (
    start test.bat exec %InstanceNumber% %InstanceNumber%
)
goto end

:exec
if "%2" GTR "1" goto more
goto finally
   
:more
set /a "result = %2 - 1"
if %displayargs% == true (
    start test.bat exec %result% %3 display
) else (
    start test.bat exec %result% %3
)

:finally
client.exe
pause
exit

:InstanceCalculator
set /a "InstanceNumber = %3 - %2 + 1 "
goto a

:end