#### go_gin_api.authorized_api
Authorized caller table

| Serial number | name | description | type | key | empty | extra | default value |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int(11) unsigned | PRI | NO | auto_increment | |
| 2 | business_key | caller key | varchar(30) | | NO | | |
| 3 | method | request method | varchar(30) | | NO | | |
| 4 | api | request address | varchar(100) | | NO | | |
| 5 | is_deleted | whether to delete 1: yes -1: no | tinyint(1) | | NO | | -1 |
| 6 | created_at | created time | timestamp | | NO | | CURRENT_TIMESTAMP |
| 7 | created_user | created by | varchar(60) | | NO | | |
| 8 | updated_at | update time | timestamp | | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 9 | updated_user | updated by | varchar(60) | | NO | | |