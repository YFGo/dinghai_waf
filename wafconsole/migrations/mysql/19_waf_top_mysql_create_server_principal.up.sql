create TABLE if not exists server_principal (
    id bigint not null auto_increment,
    server_id bigint not null,
    principal_id bigint not null,
    create_time datetime not null default current_timestamp comment '创建时间',
    update_time datetime not null default current_timestamp on update current_timestamp comment '更新时间',
    primary key (id)
);
