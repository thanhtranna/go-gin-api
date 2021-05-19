#### go_gin_api.menu
Left menu bar table

| Serial number | name | description | type | key | empty | extra | default value |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int(11) unsigned | PRI | NO | auto_increment | |
| 2 | pid | parent class id | int(11) unsigned | | NO | | 0 |
| 3 | name | menu name | varchar(32) | | NO | | |
| 4 | link | link address | varchar(100) | | NO | | |
| 5 | icon | icon | varchar(60) | | NO | | |
| 6 | level | menu type 1: first level menu 2: second level menu | tinyint(1) unsigned | | NO | | 1 |
| 7 | is_used | whether to enable 1: yes -1: no | tinyint(1) | | NO | | 1 |
| 8 | is_deleted | whether to delete 1: yes -1: no | tinyint(1) | | NO | | -1 |
| 9 | created_at | created time | timestamp | | NO | | CURRENT_TIMESTAMP |
| 10 | created_user | created by | varchar(60) | | NO | | |
| 11 | updated_at | update time | timestamp | | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 12 | updated_user | updated by | varchar(60) | | NO | | |
