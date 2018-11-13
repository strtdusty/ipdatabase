### IPDatabase

Allows you to search for IP addresses across a list of your know IP/CIDR ranges.  If you see a log message with an IP in it but you aren't sure what resource that represents, you can run

```
$ ipdatabase search 192.168.0.1  
192.168.0.0/16 home-network
192.168.0.0/18 media-subnet
```



```
Usage: ipdatabase <flags> <subcommand> <subcommand args>

Subcommands:
        add              Add a new entry
        commands         list all command names
        delete           Delete an entry
        flags            describe all known top-level flags
        help             describe subcommands and their syntax
        list             List all entries
        search           Search all entries

```