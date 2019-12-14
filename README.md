# Jute


[![build](https://img.shields.io/github/workflow/status/go-zookeeper/jute/unittest/master)](https://github.com/go-zookeeper/jute/actions?query=workflow%3Aunittest)
[![coverage](https://img.shields.io/codecov/c/github/go-zookeeper/jute)](https://codecov.io/gh/go-zookeeper/jute)
[![license](https://img.shields.io/github/license/go-zookeeper/jute)](https://raw.githubusercontent.com/go-zookeeper/jute/master/LICENSE)


Jute is a implementation of Hadoop's record serialization format written in pure Go.  This module includes a parser to real the `jute` DDL file into an AST and then a generator to generate serialization/deserialization code.