CREATE DATABASE management;
USE management;



CREATE TABLE user (
    user_id   bigint(64)  NOT NULL AUTO_INCREMENT,
    id        varchar(32) NOT NULL,
    password  varchar(72) NOT NULL,
    status    varchar(8)  NOT NULL,

    PRIMARY KEY (user_id),
    UNIQUE KEY (id),
    KEY (status)
);



CREATE TABLE personal_info (
    user_id           bigint(64)   NOT NULL,
    name              varchar(16)  NOT NULL,
    email             varchar(128) NOT NULL,
    phone_number      varchar(16)  NOT NULL,
    account_bank_name varchar(32)  NOT NULL,
    account_number    varchar(16)  NOT NULL,

    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id),
    UNIQUE KEY (phone_number),
    UNIQUE KEY (email)
);



CREATE TABLE relation (
    user_id        bigint(64) NOT NULL,
    recommender_id bigint(64) NOT NULL,

    PRIMARY KEY (user_id),
    FOREIGN KEY (recommender_id) REFERENCES user (user_id)
);



CREATE TABLE mileage (
    user_id bigint(64) NOT NULL,
    amount  bigint(64) NOT NULL DEFAULT 0,

    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);



CREATE TABLE weekly_mileage (
    user_id  bigint(64) NOT NULL,
    monday   bigint(64) DEFAULT 0,
    tuesday  bigint(64) DEFAULT 0,
    wensday  bigint(64) DEFAULT 0,
    thursday bigint(64) DEFAULT 0,
    friday   bigint(64) DEFAULT 0,

    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);



CREATE TABLE mileage_earned (
    user_id  bigint(64) NOT NULL,
    monday   bigint(64) DEFAULT 0,
    tuesday  bigint(64) DEFAULT 0,
    wensday  bigint(64) DEFAULT 0,
    thursday bigint(64) DEFAULT 0,
    friday   bigint(64) DEFAULT 0,

    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);



CREATE TABLE mileage_degree (
    user_id bigint(64) NOT NULL,
    degree  tinyint(8) NOT NULL,
    amount  bigint(64) NOT NULL DEFAULT 0,

    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);



CREATE TABLE mileage_request (
    req_id     bigint(64) NOT NULL AUTO_INCREMENT,
    user_id    bigint(64) NOT NULL,
    amount     bigint(64) NOT NULL DEFAULT 0,
    created_at datetime   NOT NULL,
    state      varchar(8) NOT NULL,

    PRIMARY KEY (req_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);


