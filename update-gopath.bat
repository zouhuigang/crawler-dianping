@echo off
wmic ENVIRONMENT where "name='gopath'" delete
wmic ENVIRONMENT create name="gopath",username="<system>",VariableValue="%~dp0"
echo %gopath%
echo %~dp0
pause