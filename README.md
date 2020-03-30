# zendesk_search
A Simple CLI to search data over a given json file

## 1. Download the executable for your OS
https://github.com/xzfvictor/zendesk_search/releases

## 2. Rename the executable
Linux
```
$ mv zendesk_search_linux-amd64 zendesk_search
$ chmod 777 zendesk_search
```
MAC
```
$ zendesk_search_linux-amd64 zendesk_search
$ chmod 777 zendesk_search
```
Windows
```
zendesk_search_windows-amd64.exe -> zendesk_search.exe
```
## 3. Make sure the ```zendesk_search``` is in the same folder with json files

## 4. List searchable keys
```
zendesk_search checkfile -f FILENAME
```
## 5. Search data on a given key
```
zendesk_search -f FILENAME -k KEY -d "data"
```
