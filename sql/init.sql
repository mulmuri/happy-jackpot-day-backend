INSERT INTO user (id, password, status)
VALUES (usertest, abcdef, Member);



INSERT INTO user (id, password, status)
VALUES ("admin", "admin000", "Admin");

INSERT INTO personal_info (user_id, name, email, phone_number, account_bank_name, account_number)
VALUES (1, "오규민", "darblue31415@gmail.com", "01041290572", "kakaobank", "333317307189");

INSERT INTO relation (user_id, recommender_id)
VALUES (1, 1);

INSERT INTO mileage (user_id)
VALUES (1);

INSERT INTO weekly_mileage (user_id)
VALUES (1);

INSERT INTO mileage_earned (user_id)
VALUES (1);




INSERT INTO user (id, password, status)
VALUES ("member", "member00", "Member");

INSERT INTO personal_info (user_id, name, email, phone_number, account_bank_name, account_number)
VALUES (2, "오규민", "darb1415@gmail.com", "01041214232", "kakaobank", "333317307189");

INSERT INTO relation (user_id, recommender_id)
VALUES (2, 1);

INSERT INTO mileage (user_id)
VALUES (2);

INSERT INTO weekly_mileage (user_id)
VALUES (2);

INSERT INTO mileage_earned (user_id)
VALUES (2);



INSERT INTO user (id, password, status)
VALUES ("awater", "awater00", "Awater");

INSERT INTO personal_info (user_id, name, email, phone_number, account_bank_name, account_number)
VALUES (3, "오규민", "da15@gmail.com", "01041169032", "kakaobank", "333317307189");

INSERT INTO relation (user_id, recommender_id)
VALUES (3, 1);

INSERT INTO mileage (user_id)
VALUES (3);

INSERT INTO weekly_mileage (user_id)
VALUES (3);

INSERT INTO mileage_earned (user_id)
VALUES (3);
