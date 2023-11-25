USE echo-crud;

CREATE TABLE IF NOT EXISTS pokemons(
  id          INT           UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT 'ID',
  no          INT           UNSIGNED NOT NULL                 COMMENT 'No.',
  name        VARCHAR(10)   NOT NULL                          COMMENT '名前',
  species     VARCHAR(10)   NOT NULL                          COMMENT '分類',
  h           INT           UNSIGNED NOT NULL                 COMMENT 'HP',
  a           INT           UNSIGNED NOT NULL                 COMMENT 'こうげき',
  b           INT           UNSIGNED NOT NULL                 COMMENT 'ぼうぎょ',
  c           INT           UNSIGNED NOT NULL                 COMMENT 'とくこう',
  d           INT           UNSIGNED NOT NULL                 COMMENT 'とくぼう',
  s           INT           UNSIGNED NOT NULL                 COMMENT 'すばやさ',
  created_on  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='ポケモン';