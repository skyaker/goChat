INSERT INTO users (username, password, email) VALUES
('alice', 'password1', 'alice@example.com'),
('bob', 'password2', 'bob@example.com'),
('charlie', 'password3', 'charlie@example.com'),
('dave', 'password4', 'dave@example.com'),
('eve', 'password5', 'eve@example.com'),
('frank', 'password6', 'frank@example.com'),
('grace', 'password7', 'grace@example.com'),
('heidi', 'password8', 'heidi@example.com'),
('ivan', 'password9', 'ivan@example.com'),
('judy', 'password10', 'judy@example.com');


INSERT INTO relations (user_1_id, user_2_id, status, status_creator) VALUES
(1, 2, 'pending', 1),  -- Alice (id=1) отправляет запрос Bobу (id=2)
(2, 3, 'accepted', 3), -- Charlie (id=3) принял запрос от Bob (id=2)
(3, 4, 'blocked', 3),  -- Charlie (id=3) блокирует Dave (id=4)
(4, 5, 'pending', 5),  -- Eve (id=5) отправляет запрос Dave (id=4)
(5, 6, 'accepted', 5), -- Eve (id=5) приняла запрос от Frank (id=6)
(6, 7, 'pending', 7),  -- Grace (id=7) отправляет запрос Frank (id=6)
(7, 8, 'blocked', 7),  -- Grace (id=7) блокирует Heidi (id=8)
(8, 9, 'accepted', 9), -- Ivan (id=9) принял запрос от Heidi (id=8)
(9, 10, 'pending', 9), -- Ivan (id=9) отправляет запрос Judy (id=10)
(10, 1, 'accepted', 1); -- Alice (id=1) приняла запрос от Judy (id=10)