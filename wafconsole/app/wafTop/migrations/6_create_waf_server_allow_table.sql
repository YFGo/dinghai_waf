create table waf_server_allow(    id         bigint unsigned auto_increment        primary key,    created_at datetime(3) null,    updated_at datetime(3) null,    deleted_at datetime(3) null,    allow_id   bigint      not null comment '''白名单id''',    server_id  bigint      not null comment '''server_id''');create index idx_waf_server_allow_deleted_at    on waf_server_allow (deleted_at);