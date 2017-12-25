CREATE TABLE queue_items
(
  id           SERIAL      NOT NULL
    CONSTRAINT queue_items_pkey
    PRIMARY KEY,
  user_id      VARCHAR(30) NOT NULL,
  next_item_id INTEGER
    CONSTRAINT next_item_id
    REFERENCES queue_items
    ON DELETE SET NULL,
  prev_item_id INTEGER
    CONSTRAINT prev_item_id
    REFERENCES queue_items
    ON DELETE SET NULL
);

CREATE UNIQUE INDEX queue_items_id_uindex
  ON queue_items (id);

CREATE TABLE queues
(
  id            SERIAL NOT NULL
    CONSTRAINT queues_pkey
    PRIMARY KEY,
  name          VARCHAR(30),
  first_item_id INTEGER
    CONSTRAINT first_item_id
    REFERENCES queue_items
    ON UPDATE SET DEFAULT ON DELETE SET DEFAULT,
  last_item_id  INTEGER
    CONSTRAINT last_item_id
    REFERENCES queue_items
    ON UPDATE SET DEFAULT ON DELETE SET DEFAULT
);

CREATE UNIQUE INDEX queues_id_uindex
  ON queues (id);

