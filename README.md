# Jute

[![godoc](https://godoc.org/github.com/go-zookeeper/jute?status.svg)](https://godoc.org/github.com/go-zookeeper/jute)
[![build](https://img.shields.io/github/workflow/status/go-zookeeper/jute/unittest/master)](https://github.com/go-zookeeper/jute/actions?query=workflow%3Aunittest)
[![coverage](https://img.shields.io/codecov/c/github/go-zookeeper/jute)](https://codecov.io/gh/go-zookeeper/jute)
[![license](https://img.shields.io/github/license/go-zookeeper/jute)](https://raw.githubusercontent.com/go-zookeeper/jute/master/LICENSE)


Jute is a implementation of [Apache Zookeepers](https://zookeeper.apache.org/)'s jute IDL and generation code written in pure Go.  This module includes a parser to real the `jute` DDL file into an AST and then a generator to generate serialization/deserialization code.

Jute started out as a serialization format for [Hadoop](https://hadoop.apache.org/) and is no longer being used (Hadoop moved to [Apache Avro](https://avro.apache.org/) as such only Zookeeper seems to be using it.  **Note:** There may be work in future versions of Zookeeper to use a different format as well (see: [ZOOKEEPER-102](https://issues.apache.org/jira/browse/ZOOKEEPER-102)).

Each jute module is mapped to a Go package with each class getting it's own go file. 

# Usage

```
$ go get github.com/go-zookeeper/jute/cmd/jutec
$ jutec input.jute
```

You can control the output folder structure with the `go.moduleMap` and `go.prefix` options.  

 `go.prefix` can control the import prefix for the generated packages.  Packages generated will append this prefix to the package names, however the prefix is not used in the generated directories.

 `go.moduleMap` uses regular expressions to rewrite module names to go packages names.  The format for this option is `<regexp match>:<regexp replacement>`.  If the replacement string is `-` then the module is skipped and no code is generated.  All moduleMaps are processed in order. Unlike `go.prefix` any mapped modules will be represented in the generated directory structure for the packages.  

## Example
This command will ignore the `org.apache.zookeeper.server.*` and the `org.apache.zookeeper.txn` modules.  All other modules matching `org.apache.zookeeper` will drop the `org.apache.zookeeper` prefix.  All generated modules will have the prefix `github.com/go-zookeeper/zk/internal` appended to it (for imports) but it won't be incuded in the directories.

```
$ jutec \
    -go.moduleMap=org.apache.zookeeper.server:- \
    -go.moduleMap=org.apache.zookeeper.txn:- \
    -go.moduleMap=org.apache.zookeeper: \
    -go.prefix=github.com/go-zookeeper/zk/internal \
    testdata/zookeeper.jute
```


# Caveats
This is pretty experimental code.  There are a lot of error checking missing, imports are not implemented and there are probably some pretty bad bugs.  You have been warned!

# TODO
## Parser
- [ ] Check for duplicate field names 
- [ ] Validate external module class fields
- [ ] Block reserved keywords

## Generator
- [ ] Support imported jute files