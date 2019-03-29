# Sync Core

>  Open, Freedom, Shared & Connected

TODO List

* [ ]  -> sync manager
* [ ]  -> sync scheduler
* [ ]  -> sync worker
* [ ]  sync build-in rproxy
* [ ]  sync build-in rsync
* [ ]  sync build-in python

Experimential

* [ ]  sync build-in script
* [ ]  sync plug-in github
* [ ]  sync plug-in commercial

## Authors

Shiwei Wang

## TODO

* [x]  vscode
* [x]  docker
* [x]  golang
* [x]  github.com/gin-gonic/gin
* [x]  github.com/BurntSushi/toml
* [x]  gopkg.in/op/go-logging.v1
* [x]  gopkg.in/urfave/cli.v1

sync manager

* [ ]  sync manager
* [ ]  sync REST interface

sync scheduler

* [ ]  sync scheduler

sync worker

* [ ]  sync worker
* [ ]  sync manager interface
* [ ]  sync plug-in interface

## Design

https://github.com/tuna/tunasync

```
# Architecture

- Manager: Central instance for status and job management
- Worker: Runs mirror jobs

+------------+ +---+                  +---+
| Client API | |   |    Job Status    |   |    +----------+     +----------+ 
+------------+ |   +----------------->|   |--->|  mirror  +---->|  mirror  | 
+------------+ |   |                  | w |    |  config  |     | provider | 
| Worker API | | H |                  | o |    +----------+     +----+-----+ 
+------------+ | T |   Job Control    | r |                          |       
+------------+ | T +----------------->| k |       +------------+     |       
| Job/Status | | P |   Start/Stop/... | e |       | mirror job |<----+       
| Management | | S |                  | r |       +------^-----+             
+------------+ |   |   Update Status  |   |    +---------+---------+         
+------------+ |   <------------------+   |    |     scheduler     |
|   BoltDB   | |   |                  |   |    +-------------------+
+------------+ +---+                  +---+


# Job Run Process


PreSyncing           Syncing                               Success
+-----------+     +-----------+    +-------------+     +--------------+
|  pre-job  +--+->|  job run  +--->|  post-exec  +-+-->| post-success |
+-----------+  ^  +-----------+    +-------------+ |   +--------------+
               |                                   |
               |      +-----------------+          | Failed
               +------+    post-fail    |<---------+
                      +-----------------+
```

## Reference

https://github.com/tuna/tunasync

https://github.com/ideal/mirror

https://github.com/etix/mirrorbits

https://github.com/ehouse/mirrors

## License

[Sync Core](https://github.com/AirWSW/sync-core) is released under the 
[GNU General Public License v3.0](https://github.com/AirWSW/sync-core/blob/master/LICENSE).
