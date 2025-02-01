CREATE table users (username text, password text);
CREATE table global_messages (sender text, message_text text, created_at timestamp default current_timestamp);
CREATE table direct_messages (sender text, receiver text, message_text text, created_at timestamp default current_timestamp);

INSERT INTO users (username, password) VALUES ('skibidi', 'skibidi');
INSERT INTO users (username, password) VALUES ('svetlana', '9999');

INSERT INTO global_messages (sender, message_text) VALUES ('skibidi', 'Hello, world!');
INSERT INTO global_messages (sender, message_text) VALUES ('svetlana', 'Hello, skibidi!');

INSERT INTO direct_messages (sender, receiver, message_text) VALUES ('skibidi', 'svetlana', 'Hello, svetlana!');