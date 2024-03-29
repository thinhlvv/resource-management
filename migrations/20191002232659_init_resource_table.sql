-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS resource(
	id INT UNSIGNED AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
	name VARCHAR(255) UNIQUE NOT NULL,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at DATETIME DEFAULT NULL,
	PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resource;
-- +goose StatementEnd

