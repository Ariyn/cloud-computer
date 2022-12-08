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
> util/build
```

## How to run

This code will run sample [half-adder](https://en.wikipedia.org/wiki/Adder_(electronics)#:~:text=%5Bedit%5D-,Half%20adder%5Bedit%5D,-Half%20adder%20logic)
You need 3 different terminals for now.
"watch outputs", "set input", "run program" will run on different terminal

Terminal 1
```bash
/cloud-computer > half-adder/run.sh
```

Terminal 2
```bash
/cloud-computer > go run ./watcher -names xor_output_1 -names and_output_1
# xor_output_1 means Sum
# and_output_1 means Carry
```

Terminal 3
```bash
/cloud-computer > redis-cli
127.0.0.1:6379> publish input_1 1
# xor_output_1 = 1, and_output_1 = 0

127.0.0.1:6379> publish input_1 0
# xor_output_1 = 0, and_output_1 = 0

127.0.0.1:6379> publish input_1 1
127.0.0.1:6379> publish input_2 1
# xor_output_0 = 0, and_output_1 = 1
```

## Next things

[ ] group by name
[ ] control config with file
[ ] make more comfortable with creating large circuit
