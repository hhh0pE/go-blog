CREATE TABLE "MY_TABLE" (
  id serial,
  name varchar,
  file varchar,
  parent_id int4
);
INSERT INTO "MY_TABLE"(id, name, file, parent_id) VALUES (5, 'redirect', null, 0);
INSERT INTO "MY_TABLE"(id, name, file, parent_id) VALUES (7, 'main', 'layout.html', 0);
INSERT INTO "MY_TABLE"(id, name, file, parent_id) VALUES (1, 'index', 'index.html', 7);
INSERT INTO "MY_TABLE"(id, name, file, parent_id) VALUES (2, 'post', 'post.html', 7);
INSERT INTO "MY_TABLE"(id, name, file, parent_id) VALUES (3, 'category', 'category.html', 7);
INSERT INTO "MY_TABLE"(id, name, file, parent_id) VALUES (8, 'admin', 'admin_layout.html', 0);