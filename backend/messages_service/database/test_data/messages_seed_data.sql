INSERT INTO dialogs (user_1_id, user_2_id, last_message, last_message_at) VALUES
(1, 2, 'Hello!', '2023-01-01 10:00:00'),
(2, 3, 'Hi there!', '2023-01-02 11:00:00'),
(3, 4, 'Good morning!', '2023-01-03 12:00:00'),
(5, 6, 'Greetings!', '2023-01-04 13:00:00'),
(6, 7, 'Hey!', '2023-01-05 14:00:00'),
(7, 8, 'Morning!', '2023-01-06 15:00:00'),
(9, 10, 'Salutations!', '2023-01-07 16:00:00');

INSERT INTO messages (dialog_id, sender_id, content, created_at) VALUES
(1, 1, 'Hello!', '2023-01-01 10:00:00'),
(1, 2, 'Hi!', '2023-01-01 10:05:00'),
(2, 2, 'Hi there!', '2023-01-02 11:00:00'),
(2, 3, 'Hello!', '2023-01-02 11:05:00'),
(3, 3, 'Good morning!', '2023-01-03 12:00:00'),
(3, 4, 'Hello!', '2023-01-03 12:05:00'),
(4, 5, 'Greetings!', '2023-01-04 13:00:00'),
(4, 6, 'Hello!', '2023-01-04 13:05:00'),
(5, 6, 'Hey!', '2023-01-05 14:00:00'),
(5, 7, 'Hi!', '2023-01-05 14:05:00'),
(6, 7, 'Morning!', '2023-01-06 15:00:00'),
(6, 8, 'Good day!', '2023-01-06 15:05:00'),
(7, 9, 'Salutations!', '2023-01-07 16:00:00'),
(7, 10, 'Hello!', '2023-01-07 16:05:00');