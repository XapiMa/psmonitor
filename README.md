# psChecker
This is a tool that provides process health checks.
Check if there are processes with the specified information.

## Installation
```
$ go get github.com/XapiMa/psChecker/cmd/psChecker
```

or

```
$ git clone https://github.com/XapiMa/psChecker.git
$ go build ./psChecker/cmd/psChecker
```

If you need a different Architecture executable file:

```
 $ git clone https://github.com/XapiMa/psChecker.git
 $ GOOS=linux GOARCH=amd64 go build main.go -o psChecker
```
Please refer to the official document for details of available environment.
https://golang.org/doc/install/source#environment


## Usage
### show current processes
```
$ psChecker show

- exec: /sbin/auditd
  cmd: /sbin/auditd 
  user: root
  pid: 4875
- exec: /usr/sbin/NetworkManager
  cmd: /usr/sbin/NetworkManager    --no-daemon
  user: root
  pid: 4988
```


### Write whitelist and blacklist
Write whitelist and blacklist in follow format:

- whitelist.yml
```
- user: root
  pid: 4875
  exec: /sbin/auditd
- exec: /usr/sbin/NetworkManager
- cmd: /System/Library/CoreServices/appleeventsd --server
  user: _appleevents
  regexp: .*server$
```

- blacklist.yaml
```
- exec: /usr/sbin/badScript
- cmd : ./badObject
- regexp: .*crack.*
```

Warn when there is no process with the value set to alive and when there is a process with the value set to dead.

Possible values are user, pid, exec, cmd, open and regexp.
- user: Execution user name
- pid: Process ID
- exec: executable file path
- cmd: runtime command line string
- open: files of opend by pocess
- regexp: regular expression to search for all user, exec, cmd

But open is not supported on mac.

## Execution

1. get current process list
```
# sudo psChecker show -o path/to/output
```

2. write whitelist.yml and blacklist.yml

3. monitoring processes
```
$ sudo psChecker monitor -w path/to/whitelist.yml -b path/to/blacklist.yml
```

If you want to see details such as errors for both `psChecker show` and `psChecker monitor`, give -v option.
