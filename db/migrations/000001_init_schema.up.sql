
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(1000) NOT NULL,
  "price" float8 NOT NULL,
  "reorder_level" int NOT NULL,
  "description" text,
  "has_discount" smallint NOT NULL DEFAULT 0,
  "image_id" uuid NOT NULL,
  "discount_percentage" float4 NOT NULL DEFAULT 0,
  "brand_id" uuid,
  "size_id" int DEFAULT 1,
  "category_id" int,
  "subcategory_id" int,
  "color_id" int DEFAULT 1,
  "status_id" int DEFAULT 0,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "images" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(200) NOT NULL DEFAULT 'unnamed',
  "url" text NOT NULL,
  "category_id" smallint NOT NULL DEFAULT 2,
  "size" smallint NOT NULL DEFAULT 1
);

CREATE TABLE "image_categories" (
  "id" smallint PRIMARY KEY NOT NULL,
  "name" varchar(45) NOT NULL
);

CREATE TABLE "image_sizes" (
  "id" smallint PRIMARY KEY NOT NULL,
  "name" varchar(45) NOT NULL
);

CREATE TABLE "status" (
  "id" int PRIMARY KEY NOT NULL,
  "name" varchar(45) UNIQUE NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sizes" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(45) UNIQUE NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "brands" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(100) UNIQUE NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(45) UNIQUE NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "subcategories" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(45) UNIQUE NOT NULL,
  "categories_id" int NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "colors" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(20) UNIQUE NOT NULL,
  "hex" varchar(7) UNIQUE NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "stock" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "quantity" int NOT NULL DEFAULT 0,
  "product_id" uuid NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "locations" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(45) UNIQUE NOT NULL,
  "quantity" int NOT NULL,
  "stock_id" uuid NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now()),
  "update_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "inventory_log" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "client_id" uuid NOT NULL,
  "operation_type_id" int NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" int NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "operation_type" (
  "id" smallint PRIMARY KEY NOT NULL,
  "name" varchar(45) UNIQUE NOT NULL
);

ALTER TABLE "products" ADD FOREIGN KEY ("image_id") REFERENCES "images" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("brand_id") REFERENCES "brands" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("size_id") REFERENCES "sizes" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("subcategory_id") REFERENCES "subcategories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("color_id") REFERENCES "colors" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("category_id") REFERENCES "image_categories" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("size") REFERENCES "image_sizes" ("id");

ALTER TABLE "subcategories" ADD FOREIGN KEY ("categories_id") REFERENCES "categories" ("id");

ALTER TABLE "stock" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "locations" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");

ALTER TABLE "inventory_log" ADD FOREIGN KEY ("operation_type_id") REFERENCES "operation_type" ("id");

ALTER TABLE "inventory_log" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");



CREATE FUNCTION update_stock()
  RETURNS TRIGGER AS $update_quantity$
  DECLARE
    stockquantity integer := (SELECT quantity FROM stock WHERE product_id = New.product_id);
    ordertype integer := new.operation_type_id;
  BEGIN
    IF ordertype  = 1 THEN
      UPDATE stock
      SET quantity = stockquantity + New.quantity
      WHERE product_id = New.product_id;
    ELSEIF ordertype  = 2 THEN
      UPDATE stock
      SET quantity = stockquantity - New.quantity
      WHERE product_id = New.product_id;
    END IF;
	RETURN NEW;
  END; 
  $update_quantity$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_stock BEFORE INSERT OR UPDATE ON inventory_log
    FOR EACH ROW EXECUTE PROCEDURE update_stock();
	

CREATE FUNCTION insert_stock()
  RETURNS TRIGGER AS $insert_stock$
  BEGIN
    INSERT INTO stock (product_id) VALUES(New.id);
	RETURN NEW;
  END; 
  $insert_stock$ LANGUAGE 'plpgsql';

  CREATE TRIGGER insert_stock AFTER INSERT ON products
    FOR EACH ROW EXECUTE PROCEDURE insert_stock();


CREATE FUNCTION update_status_bystock()
RETURNS TRIGGER AS $update_status_bystock$
	BEGIN
		IF New.quantity <= 0 THEN
		UPDATE products
			SET status_id = 2
			WHERE id = New.product_id;
		ELSEIF New.quantity <= (SELECT reorder_level FROM products WHERE id = New.product_id) THEN
			UPDATE products
			SET status_id = 3
			WHERE id = New.product_id;
		ELSEIF New.quantity > (SELECT reorder_level FROM products WHERE id = New.product_id) THEN
			UPDATE products
			SET status_id = 1
		WHERE id = New.product_id;
	  END IF;
	  RETURN NEW;
  END; 
$update_status_bystock$ LANGUAGE 'plpgsql';

  CREATE TRIGGER update_status_bystock AFTER UPDATE ON stock
    FOR EACH ROW EXECUTE PROCEDURE update_status_bystock();

INSERT INTO brands(name) VALUES('NVIDIA');
INSERT INTO brands(name) VALUES('Xbox');
INSERT INTO brands(name) VALUES('PlayStation');
INSERT INTO brands(name) VALUES('MSI');
INSERT INTO brands(name) VALUES('ASUS');

INSERT INTO image_sizes(id, name) VALUES(1, 'Large');
INSERT INTO image_sizes(id, name) VALUES(2, 'Mid');
INSERT INTO image_sizes(id, name) VALUES(3, 'Small');

INSERT INTO sizes(name) VALUES('unidad');
INSERT INTO sizes(name) VALUES('metros');
INSERT INTO sizes(name) VALUES('centimetros');
INSERT INTO sizes(name) VALUES('kilometros');
INSERT INTO sizes(name) VALUES('12 unidades');
INSERT INTO sizes(name) VALUES('6 unidades');
INSERT INTO sizes(name) VALUES('3 unidades');

INSERT INTO colors(name, hex) VALUES('Gray', '#808080');
INSERT INTO colors(name, hex) VALUES('Maroon', '#800000');
INSERT INTO colors(name, hex) VALUES('Olive', '#808000');
INSERT INTO colors(name, hex) VALUES('Green', '#008000');
INSERT INTO colors(name, hex) VALUES('Purple', '#800080');
INSERT INTO colors(name, hex) VALUES('Teal', '#008080');
INSERT INTO colors(name, hex) VALUES('Navy', '#000080');
INSERT INTO colors(name, hex) VALUES('Black', '#000000');
INSERT INTO colors(name, hex) VALUES('White', '#FFFFFF');
INSERT INTO colors(name, hex) VALUES('Red', '#FF0000');
INSERT INTO colors(name, hex) VALUES('Lime', '#00FF00');
INSERT INTO colors(name, hex) VALUES('Blue', '#0000FF');
INSERT INTO colors(name, hex) VALUES('Yellow', '#FFFF00');
INSERT INTO colors(name, hex) VALUES('Cyan-Aqua', '#00FFFF');
INSERT INTO colors(name, hex) VALUES('Magenta-Fuchsia', '#FF00FF');

INSERT INTO categories(name) VALUES('Tecnología');
INSERT INTO categories(name) VALUES('Hogar');
INSERT INTO categories(name) VALUES('Papelería');

INSERT INTO subcategories(name, categories_id) VALUES('Videojuegos', (SELECT id FROM categories WHERE name = 'Tecnología'));

INSERT INTO status(id, name) VALUES(0, 'NOT STATUS');
INSERT INTO status(id, name) VALUES(1, 'AVAILABLE');
INSERT INTO status(id, name) VALUES(2, 'OUT STOCK');
INSERT INTO status(id, name) VALUES(3, 'LOW LEVEL STOCK');

INSERT INTO operation_type(id, name) VALUES(1, 'IN');
INSERT INTO operation_type(id, name) VALUES(2, 'OUT');

INSERT INTO image_categories(id, name) VALUES(1, 'Main');
INSERT INTO image_categories(id, name) VALUES(2, 'Gallery');
INSERT INTO image_categories(id, name) VALUES(3, 'Banner');
INSERT INTO image_categories(id, name) VALUES(4, 'Discount');

INSERT INTO images(name, url, category_id) VALUES('default', 'https://www.pngkey.com/png/full/124-1244676_box-package-delivery-shipping-product-comments-icon-for.png', 1);


INSERT INTO products(name, price, reorder_level, description, image_id, category_id, brand_id, subcategory_id,status_id) 
VALUES('Xbox Series X', 499.99, 20, 'Consola de siguiente generación, con 1tb de almacenamiento', 
(SELECT id FROM images WHERE name = 'default'),
1,
(SELECT id FROM brands WHERE name = 'Xbox'),
1,0);

INSERT INTO products(name, price, reorder_level, description, image_id, category_id, brand_id, subcategory_id,status_id) 
VALUES('PlayStation 5 Disck Unit', 499.99, 20, 'Consola de siguiente generación, con 1tb de almacenamiento', 
(SELECT id FROM images WHERE name = 'default'),
1,
(SELECT id FROM brands WHERE name = 'PlayStation'),
1,0);

INSERT INTO products(name, price, reorder_level, description, image_id, category_id, brand_id, subcategory_id,status_id) 
VALUES('RTX 3090', 1499.99, 10, 'Tarjeta de video de gama alta', 
(SELECT id FROM images WHERE name = 'default'),
1,
(SELECT id FROM brands WHERE name = 'MSI'),
1,0);

INSERT INTO products(name, price, reorder_level, description, image_id, category_id, brand_id, subcategory_id,status_id) 
VALUES('RTX 3080', 799.99, 15, 'Tarjeta de video de gama alta', 
(SELECT id FROM images WHERE name = 'default'),
1,
(SELECT id FROM brands WHERE name = 'MSI'),
1,0);

INSERT INTO inventory_log(client_id, operation_type_id, product_id, quantity) 
VALUES(uuid_generate_v4(), 1, (SELECT id FROM products WHERE name = 'Xbox Series X'), 5);

INSERT INTO inventory_log(client_id, operation_type_id, product_id, quantity) 
VALUES(uuid_generate_v4(), 1, (SELECT id FROM products WHERE name = 'PlayStation 5 Disck Unit'), 5);

INSERT INTO inventory_log(client_id, operation_type_id, product_id, quantity) 
VALUES(uuid_generate_v4(), 1, (SELECT id FROM products WHERE name = 'RTX 3090'), 5);

INSERT INTO inventory_log(client_id, operation_type_id, product_id, quantity) 
VALUES(uuid_generate_v4(), 1, (SELECT id FROM products WHERE name = 'RTX 3080'), 5);
