## Excuting an order
Execute the script in the root directory: `./scripts/gormgen.sh addr user pass name tables`;
-addr: database address, for example: 127.0.0.1:3306
-user: account, for example: root
-pass: password, for example: root
-name: database name, for example: go_gin_api
-tables: table name, the default is *, multiple table names can be separated by ",", for example: user_demo

E.g:
```
./scripts/gormgen.sh 127.0.0.1:3306 root root go_gin_api user_demo
```

## Reference
- https://github.com/MohamedBassem/gormgen