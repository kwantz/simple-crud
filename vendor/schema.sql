CREATE TABLE IF NOT EXISTS Product (
	id INTEGER AUTO_INCREMENT,
	shop_id INTEGER NOT NULL,
	category_id INTEGER NOT NULL,
	name VARCHAR(30) NOT NULL,
	price INTEGER NOT NULL,
	stock INTEGER NOT NULL,
	description VARCHAR(255) NOT NULL,

    PRIMARY KEY (id),
    INDEX (shop_id),
    INDEX (category_id)
)
