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
$ mv zendesk_search_darwin-amd64 zendesk_search
$ chmod 777 zendesk_search
```
Windows
```
ren zendesk_search_windows-amd64.exe zendesk_search.exe
```
## 3. Make sure ```zendesk_search``` is in the same folder along with json files

## 4. List searchable keys on a json file
Linux/Mac
```
zendesk_search checkfile -f users.json
```
Windows
```
zendesk_search.exe checkfile -f tickets.json
```
## 5. Search data on a given key of the json file
Linux/Mac
```
zendesk_search -f organizations.json -k domain_names -d "ecratic.com"
```
Windows
```
zendesk_search.exe -f users.json -k name -d "Jessica Raymond"
```
