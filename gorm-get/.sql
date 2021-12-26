create database weathers;

use weathers;

CREATE TABLE weathers
(
id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
name            VARCHAR(150) NOT NULL,                # Name city
data            VARCHAR(150) NOT NULL,                # weather of city in next 12h
PRIMARY KEY     (id)                                  # Make the id the primary key
);

ALTER TABLE weathers ADD INDEX `name_index` (`name`);

INSERT INTO weathers (name, data) VALUES ("HCM", "sunning"),("Beijing", "raining"),("Hanoi", "cloudy");