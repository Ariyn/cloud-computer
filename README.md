# cloud-computer

This project simulate logic-gates based on cloud computing concept.

Each logic gates will run as single process and they connected with redis [pub/sub](https://redis.io/docs/manual/pubsub/).
So you can make distributed cpu.

## Requirements

- redis on localhost
  
  Redis should support pub/sub
  
- go 1.19.4

  Other version can handle this project, but I didn't test them.

## How to install

```bash
> git clone github.com/ariyn/cloud-computer
> cd cloud-computer
> util/build.sh
```

## How to build

pass .ll script path to parser.
parser will print bash scripts to stdout.

```bash
/cloud-computer > go run ./parser -file half-adder.ll > bin/half-adder
```

You can see half adder .ll script at [here](https://github.com/Ariyn/cloud-computer/blob/main/logic-layouts/half-adder.ll). Spec of .ll script is [here](https://github.com/Ariyn/cloud-computer/blob/main/logic-layouts/spec.ll).

## How to run

This code will run sample [half-adder](https://en.wikipedia.org/wiki/Adder_(electronics)#:~:text=%5Bedit%5D-,Half%20adder%5Bedit%5D,-Half%20adder%20logic).

You need 3 different terminals for now.
"watch outputs", "input", "run program" runs on different terminals.

### Terminal 1
```bash
/cloud-computer > bin/half-adder -name ha
```

Half adder's name will be ha. Every children of half adder will be ha.XXX

Basically name of single gate is combination of parent, grand parent, ...

Terminal 2
```bash
/cloud-computer > go run ./watcher -names ha.sum -names ha.carry
# carry and sum is defined in half-adder.ll and compiled into bash script
```

Watcher will subscribe names and print when they changed.
You can watch any thing you want to see.

WARN: You can't watch input, because there is no "input". Every gates publish their "output" to redis and subscribe other's "output" as input.

Terminal 3
```bash
/cloud-computer > redis-cli
127.0.0.1:6379> publish ha.inputs.1 1
# ha.sum = 1, ha.carry = 0

127.0.0.1:6379> publish ha.inputs.1 0
# ha.sum = 0, ha.carry = 0

127.0.0.1:6379> publish ha.inputs.2 1
# ha.sum = 1, ha.carry = 0
127.0.0.1:6379> publish ha.inputs.1 1
# ha.sum = 0, ha.carry = 1
```

If you publish 1 or 0 to ha.inputs.1, subscribing gates will receive, do stuff and publish their output.

WARN: There is no "half adder" gate. ".ll" script will be translated into bunch of "not", "or", "and", "xor", "flip flop" gates. Even you published to ha.inputs.1, some children of half adder will receive.

## Next things

- [x] group by name
- [ ] control config with file
- [ ] make more comfortable tools to create large circuit
  - [ ] gate dump
  - [ ] more comfortable watcher
  - [ ] more comfortable input controller
