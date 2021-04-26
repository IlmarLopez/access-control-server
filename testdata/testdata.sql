INSERT INTO roles (id, name, description)
VALUES ('04d078d4-907e-4f91-b908-b0721be9fca4', 'estudiante', 'Club cuervos'),
       ('04d078d4-907e-4f91-b908-b0721be9fca5', 'guardia', '...'),
       ('04d078d4-907e-4f91-b908-b0721be9fca6', 'visitante', '...'),
       ('04d078d4-907e-4f91-b908-b0721be9fca7', 'docente', '...'),
       ('04d078d4-907e-4f91-b908-b0721be9fca9', 'admin', '...'),
       ('04d078d4-907e-4f91-b908-b0721be9fca8', 'personal', '...');

INSERT INTO buildings (id, name, description, user_limit, created_at, updated_at, is_active) 
VALUES ('2b6146eb-e636-49cd-bb77-b7d61285d594', 'Biblioteca', 'Lugar de lectrura', 47, '2021-04-21 14:08:01', '2021-04-21 14:08:22', 1),
       ('34d078d4-907e-4f91-b908-b0721be9fc0g', 'Docencia 5', 'BIS', 4, '2021-04-18 01:15:38', '2021-04-21 13:55:32', 1),
       ('34d078d4-907e-4f91-b908-b0721be9fcf3', 'docencia 2', 'csdcds', 30, '2021-04-18 01:15:38', '2021-04-21 13:55:23', 1),
       ('34d078d4-907e-4f91-b908-b0721be9fck8', 'docencia 3', null, 15, '2021-04-18 01:15:38', null, 1),
       ('34d078d4-907e-4f91-b908-b0721be9fco9', 'docencia 4', null, 10, '2021-04-18 01:15:38', null, 1),
       ('34d078d4-907e-4f91-b908-b0721be9fcr1', 'docencia 1', null, 16, '2021-04-18 01:15:38', null, 1),
       ('4f4d94ec-30aa-43d8-bf88-6681abdca5dd', 'Biblioteca', 'Estación de cuervos para lectura.', 5, '2021-04-19 00:48:57', '2021-04-19 00:48:57', 1);

INSERT INTO careers (id, name, is_active) 
VALUES ('14d078d4-907e-4f91-b908-b0721be9fca5', 'career 1', 1),
       (id, name, is_active) VALUES ('24d078d4-907e-4f91-b908-b0721be9fca5', 'career 2', 1),
       (id, name, is_active) VALUES ('34d078d4-907e-4f91-b908-b0721be9fca5', 'career 3', 1),
       (id, name, is_active) VALUES ('44d078d4-907e-4f91-b908-b0721be9fca5', 'career 4', 1),
       (id, name, is_active) VALUES ('54d078d4-907e-4f91-b908-b0721be9fca5', 'career 5', 1);

INSERT INTO groups (id, career_id, name, is_active) 
VALUES ('55d078d4-907e-4f91-b908-b0721be9fca5', '14d078d4-907e-4f91-b908-b0721be9fca5', '10 a', 1),
       (id, career_id, name, is_active) VALUES ('56d078d4-907e-4f91-b908-b0721be9fca5', '14d078d4-907e-4f91-b908-b0721be9fca5', '10 b', 1),
       (id, career_id, name, is_active) VALUES ('57d078d4-907e-4f91-b908-b0721be9fca5', '14d078d4-907e-4f91-b908-b0721be9fca5', '10 c', 1),
       (id, career_id, name, is_active) VALUES ('58d078d4-907e-4f91-b908-b0721be9fca5', '24d078d4-907e-4f91-b908-b0721be9fca5', '9 a', 1),
       (id, career_id, name, is_active) VALUES ('59d078d4-907e-4f91-b908-b0721be9fca5', '24d078d4-907e-4f91-b908-b0721be9fca5', '9 b', 1),
       (id, career_id, name, is_active) VALUES ('60d078d4-907e-4f91-b908-b0721be9fca5', '24d078d4-907e-4f91-b908-b0721be9fca5', '9 c', 1),
       (id, career_id, name, is_active) VALUES ('61d078d4-907e-4f91-b908-b0721be9fca5', '34d078d4-907e-4f91-b908-b0721be9fca5', '6 a', 1),
       (id, career_id, name, is_active) VALUES ('62d078d4-907e-4f91-b908-b0721be9fca5', '34d078d4-907e-4f91-b908-b0721be9fca5', '6 b', 1),
       (id, career_id, name, is_active) VALUES ('63d078d4-907e-4f91-b908-b0721be9fca5', '34d078d4-907e-4f91-b908-b0721be9fca5', '6 c', 1),
       (id, career_id, name, is_active) VALUES ('64d078d4-907e-4f91-b908-b0721be9fca5', '44d078d4-907e-4f91-b908-b0721be9fca5', '7 a', 1),
       (id, career_id, name, is_active) VALUES ('65d078d4-907e-4f91-b908-b0721be9fca5', '44d078d4-907e-4f91-b908-b0721be9fca5', '7 b', 1),
       (id, career_id, name, is_active) VALUES ('66d078d4-907e-4f91-b908-b0721be9fca5', '44d078d4-907e-4f91-b908-b0721be9fca5', '7 c', 1),
       (id, career_id, name, is_active) VALUES ('67d078d4-907e-4f91-b908-b0721be9fca5', '54d078d4-907e-4f91-b908-b0721be9fca5', '11 a', 1),
       (id, career_id, name, is_active) VALUES ('68d078d4-907e-4f91-b908-b0721be9fca5', '54d078d4-907e-4f91-b908-b0721be9fca5', '11 b', 1),
       (id, career_id, name, is_active) VALUES ('69d078d4-907e-4f91-b908-b0721be9fca5', '54d078d4-907e-4f91-b908-b0721be9fca5', '11 c', 1);

INSERT INTO users (id, username, password, created_at, updated_at, is_active, role_id, first_name, last_name, email, registration_number, career_id, group_id) 
VALUES ('cd8846bf-d6bc-4cf4-b2ce-68a3537c2c92', 'admin', '$2a$04$sjt7VwmOj/6LlQzxSh.HJeCFQxB8Q/a358q2bm/v.dYBMS./2Ifnu', '2021-04-16 04:05:16', null, 1, '04d078d4-907e-4f91-b908-b0721be9fca9', 'admin', 'admin', 'me@ilmarlopez.com', '4545454545', '14d078d4-907e-4f91-b908-b0721be9fca5', '56d078d4-907e-4f91-b908-b0721be9fca5');
