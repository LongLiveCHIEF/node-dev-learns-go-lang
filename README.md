# node-dev-learns-go-lang

I created this project to help myself learn go. I figure the best way to get
up to speed, is to take on some challenges that would be simple tasks for
me to complete in node.js.


## Challenge 1 - Create an HTTP Server 

- [ ] Static file server
- [ ] Returns a directory index on `host.com/path/` fragments
- [ ] Responds with 400 error code if directory and file path not found 

## Challenge 2 - Create a command line health check client

```
$ isgreen <ip_address> [--ports]

ip_address status:
port1: green
port2: green
port3: red
```
## Challenge 3 - Create a command line REPL fuzzy file finder

- [ ] typing `finder` drops into REPL mode with `find>` prompt
- [ ] As you type characters, an autocomplete style list appears in terminal below the prompt

## Challenge 4 - Create a File System Cleanup Utility

To get an idea of what this program should do, try running: (on *nix OS)

```
$ find $HOME -d type -maxdepth 1 -exec du -sh {} + |sort -nr
$ find $HOME -f type -exec du -sh {} + |sort -nr
$ find $HOME -size +1000k -exec du -sh {} + |sort -nr
```

- [ ] given a path, returns largest directories|files
- [ ] can return unmodified files older than `x` days 
- [ ] can report back on unused disk space and lv's
- [ ] can report back on unmounted LV's or disks with no mounted filesystems
