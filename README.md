ssh-pub-key-mgr
========

A Microservice that polls a file, and calls github for SSH Keys.

```
NAME:
   ssh-pub-key-mgr - set source flag, and send in the path to write authorized_keys

USAGE:
   ssh-pub-key-mgr [global options] command [command options] [arguments...]

VERSION:
   v0.0.0-dev

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --keystore value, -k value            keystore backend (default: "github")
   --refresh-interval value, -r value    interval to check for updates in seconds (default: 600)
   --source file://PATH, -s file://PATH  location to file://PATH or http(s)://URL containing allowed users and fingerprints
   --help, -h                            show help
   --version, -v                         print the version
```

## Building

`make`


## Running

`./bin/ssh-pub-key-mgr`

## License
Copyright (c) 2014-2016 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
