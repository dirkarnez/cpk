@echo off

REM for private repo soon
set USER_NAME=dirkarnez
set PASSWORD=????

set TARGET_OWNER=dirkarnez
set TARGET_REPOSITORY=sfml-prebuilt
set TARGET_LIBRARY=sfml
set TARGET_TAG=v2.5.1
set TARGET_COMPILER=mingw

cpk.exe
pause