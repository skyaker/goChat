INSERT INTO users (username, password, email) VALUES
('user1', 'password1', 'user1@example.com'),
('user2', 'password2', 'user2@example.com'),
('user3', 'password3', 'user3@example.com'),
('user4', 'password4', 'user4@example.com'),
('user5', 'password5', 'user5@example.com'),
('user6', 'password6', 'user6@example.com'),
('user7', 'password7', 'user7@example.com'),
('user8', 'password8', 'user8@example.com'),
('user9', 'password9', 'user9@example.com'),
('user10', 'password10', 'user10@example.com');


INSERT INTO friends (user_id, friend_id, status) VALUES
(1, 2, 'accepted'),
(1, 3, 'pending'), 
(2, 4, 'blocked'), 
(5, 6, 'accepted'),
(6, 7, 'pending'), 
(7, 8, 'blocked'), 
(9, 10, 'accepted');