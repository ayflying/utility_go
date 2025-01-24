create table shiningu_game_act
(
    uid        bigint   not null comment '玩家编号',
    act_id     int      not null comment '活动编号',
    action     text     null comment '活动配置',
    updated_at datetime null comment '更新时间',
    primary key (uid, act_id)
)
    comment '玩家活动数据' charset = utf8mb4;

create index shiningu_game_act_uid_index
    on shiningu_game_act (uid);

