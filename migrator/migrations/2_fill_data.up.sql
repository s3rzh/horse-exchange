-- breeds
INSERT INTO breeds(name) VALUES('Чистокровная верховая');
INSERT INTO breeds(name) VALUES('Английский тяжеловоз');
INSERT INTO breeds(name) VALUES('Арабская лошадь');
INSERT INTO breeds(name) VALUES('Орловский рысак');
INSERT INTO breeds(name) VALUES('Старая фламандская');

-- horses
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Апполон', 'жеребец', 10, 810, 1, 'Ирландия. Дублин.', 'спокойная', 3000, 60000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Брюс', 'жеребец', 17, 800, 1, 'Ирландия. Уэксфорд.', 'добросовестная', 3050, 61000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Версаль', 'мерин', 4, 760, 1, 'Ирландия. Корк.', 'жизнерадостная', 3100, 65000);

INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Норд', 'жеребец', 9, 920, 2, 'Англия. Бирмингем.', 'добросовестная', 4500, 91000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Астория', 'кобыла', 12, 900, 2, 'Англия. Шеффилд.', 'нервная', 4300, 88000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Бэлла', 'кобыла', 5, 850, 2, 'Англия. Брайтон.', 'жизнерадостная', 4000, 83000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Цент', 'жеребец', 26, 800, 2, 'Англия. Портсмут.', 'добросовестная', 4100, 85000);

INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Джули', 'кобыла', 11, 830, 3, 'Катар. Доха.', 'спокойная', 5000, 100000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Холли', 'кобыла', 14, 800, 3, 'Катар. Эль-Вакра.', 'жизнерадостная', 4050, 95000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Лорд', 'мерин', 25, 790, 3, 'Катар. Эр-Райян.', 'спокойная', 4020, 91000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Мартин', 'жеребец', 9, 770, 3, 'Катар. Эд-Дайиан.', 'нервная', 4000, 90000);

INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Тюльпан', 'жеребец', 15, 850, 4, 'Россия. Воронеж.', 'добросовестная', 4200, 82000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Пегас', 'жеребец', 21, 730, 4, 'Россия. Самара.', 'спокойная', 3900, 79000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Маркиза', 'кобыла', 13, 800, 4, 'Россия. Пермь.', 'спокойная', 3000, 70000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Фараон', 'мерин', 8, 700, 4, 'Россия. Кострома.', 'нервная', 3500, 75000);

INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Лейла', 'кобыла', 21, 500, 5, 'Бельгия. Гент.', 'жизнерадостная', 2000, 55000);
INSERT INTO horses(name, gender, age, weight, breed_id, birthplace, habit, rentalprice, purchaseprice) 
VALUES('Орион', 'мерин', 14, 450, 5, 'Бельгия. Антверпен.', 'добросовестная', 2100, 60000);

-- targets
INSERT INTO tasks(title) VALUES('земледелие');
INSERT INTO tasks(title) VALUES('работа на каменоломне');
INSERT INTO tasks(title) VALUES('перевозки грузов');
INSERT INTO tasks(title) VALUES('верховые прогулки');
INSERT INTO tasks(title) VALUES('участие в соревнованиях');

-- horses_tasks
INSERT INTO horses_tasks(horse_id, task_id) VALUES(1, 4);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(1, 5);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(2, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(2, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(2, 3);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(3, 5);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(4, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(4, 3);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(5, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(6, 4);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(7, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(7, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(8, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(8, 4);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(9, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(10, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(10, 4);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(11, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(11, 3);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(11, 5);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(12, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(12, 5);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(13, 3);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(13, 4);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(14, 4);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(15, 2);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(15, 3);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(16, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(17, 1);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(17, 3);
INSERT INTO horses_tasks(horse_id, task_id) VALUES(17, 5);