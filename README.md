# request-logging-tool
Application to send http requests and log the md5 responses with a parallel request worker limit flag 


##How to run
* test project
    ```makefile
    make tests
    ``` 
  
* build the project 
    ```makefile
    make build
    ```
NOTE : this will create an executable file `myhtp`
```bash
  ./myhttp -parallel 3 google.com yahoo.com http://www.adjust.com
```
```bash
  ./myhttp http://www.adjust.com http://google.com
```

##Errors
* If the domain is incorrect and not possible to do an request it will return an error specifying the respective domain and the error
* but all other working domains will respond successfully
