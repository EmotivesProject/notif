\c notif_db;

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    username VARCHAR(128) NOT NULL,
	title VARCHAR(128) NOT NULL,
	message VARCHAR(256) NOT NULL,
	link VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    seen boolean DEFAULT false
);