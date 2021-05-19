## repository

#### Data access layer.

-`./db_repo` to access DB data
-`./cache_repo` to access Cache data

#### SQL recommendations:
-It is recommended that each table contain fields: primary key (id), mark deletion (is_deteled), created time (created_at), update time (updated_at)

```mysql
`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT'primary key',
`is_deleted` tinyint(1) NOT NULL DEFAULT'-1' COMMENT'Whether to delete 1: Yes -1: No',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT'created time',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT'update time',
```

#### Naming conventions:

-The package name should end with `_repo`;
-The package name in the `./db_repo` directory is named with `data table name` + `_repo`;

#### Script to generate MySQL CURD

-Use script to automatically generate CURD code based on table structure, see document: `./cmd/gormgen/README.md`