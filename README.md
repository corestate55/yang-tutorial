# YANG-Tutorial Turing Machine

Project to study YANG/Golang.

Target: Turing Machine Implementation used in [pyang Yang Tutorial](https://github.com/mbj4668/pyang/wiki/Tutorial)

## Reference

* [InstanceValidation · mbj4668/pyang Wiki](https://github.com/mbj4668/pyang/wiki/InstanceValidation)
* [IETF pyang Tutorial](https://www.ietf.org/slides/slides-edu-pyang-tutorial-01.pdf) (pdf)
* [pyang/doc/tutorial at master · mbj4668/pyang · GitHub](https://github.com/mbj4668/pyang/tree/master/doc/tutorial)
* [DSDLMappingTutorial < Main < TWiki](http://www.yang-central.org/twiki/bin/view/Main/DSDLMappingTutorial)

## Build/Run
Install dependency tools at first.
* [GitHub \- favadi/protoc\-go\-inject\-tag: Inject custom tags to protobuf golang struct](https://github.com/favadi/protoc-go-inject-tag)
* [GitHub \- openconfig/goyang: YANG parser and compiler to produce Go language objects](https://github.com/openconfig/goyang)

```
$ go get github.com/openconfig/goyang
$ go get github.com/favadi/protoc-go-inject-tag
```

and make server and client
```
$ make
```

Run server
```
$ ./tmserver
```
Run client (in another terminal),  and type command like below.
```
$ ./tmclient
command> help
command> get
command> init data/turing-machine-rpc.xml
command> config data/turing-machine-config.xml
command> get
command> run
command> get
command> exit
```

Then client send gRPC message to server, server works as turing-machine.
```
2018/04/29 17:18:48 Initialize: TapeContent: 110111

input        | output
state symbol | state symbol move
   S2      1 |  [S3]     0  <=  write separator
   S3      1 |  [S0]        <=  go home
   S3        |  [S4]         -: final step
   S0      1 |  [S0]         -: left summand
   S0      0 |  [S1]     1   -: separator
   S1        |  [S2]        <=  right end
   S1      1 |  [S0]         -: right summand

Step State | Tape                   | Next Write Move
   1  [S0] | <1>| 1 | 0 | 1 | 1 | 1 | [S0]         -: left summand
   2  [S0] |  1 |<1>| 0 | 1 | 1 | 1 | [S0]         -: left summand
   3  [S0] |  1 | 1 |<0>| 1 | 1 | 1 | [S1]     1   -: separator
   4  [S1] |  1 | 1 | 1 |<1>| 1 | 1 | [S0]         -: right summand
   5  [S1] |  1 | 1 | 1 | 1 |<1>| 1 | [S0]         -: right summand
   6  [S1] |  1 | 1 | 1 | 1 | 1 |<1>| [S0]         -: right summand
tape out-of-range
   7  [S1] |  1 | 1 | 1 | 1 | 1 | 1 | [S2]        <=  right end
   8  [S2] |  1 | 1 | 1 | 1 | 1 |<1>| [S3]     0  <=  write separator
   9  [S3] |  1 | 1 | 1 | 1 |<1>| 0 | [S0]        <=  go home
  10  [S3] |  1 | 1 | 1 |<1>| 1 | 0 | [S0]        <=  go home
  11  [S3] |  1 | 1 |<1>| 1 | 1 | 0 | [S0]        <=  go home
  12  [S3] |  1 |<1>| 1 | 1 | 1 | 0 | [S0]        <=  go home
  13  [S3] | <1>| 1 | 1 | 1 | 1 | 0 | [S0]        <=  go home
tape out-of-range
  14  [S3] |  1 | 1 | 1 | 1 | 1 | 0 | [S4]         -: final step
  15  [S4] | <1>| 1 | 1 | 1 | 1 | 0 | END
```
