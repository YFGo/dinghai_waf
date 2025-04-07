create TABLE if not exists principal_info (
    id bigint not null auto_increment,
    principal_type varchar(25) not null comment '通知网站负责人类型',
    principal_name varchar(35) not null comment '通知网站负责人姓名',
    principal_position varchar(35) not null comment '通知网站负责人职位',
    notice_info varchar(35) not null comment '通知网站负责人信息',
    create_time datetime not null default current_timestamp comment '创建时间',
    update_time datetime not null default current_timestamp on update current_timestamp comment '更新时间',
    primary key (id)
);