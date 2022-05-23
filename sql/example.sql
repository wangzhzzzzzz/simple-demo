DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`        varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `avatar`      varchar(128)        NOT NULL DEFAULT '' COMMENT '头像',
    `level`       int(10)             NOT NULL DEFAULT 1 COMMENT '用户等级',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'Jerry', '', 1, '2022-04-01 10:00:00', '2022-04-01 10:00:00'),
       (2, 'Tom', '', 2, '2022-04-01 10:00:00', '2022-04-01 10:00:00');