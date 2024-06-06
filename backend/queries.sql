


CREATE TABLE IF NOT EXISTS autos
(
    auto_id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    marca character varying(30),
    modelo integer,
    combustible character varying(10),
    transmision character varying(20),
    lujos character varying(2),
    sillas integer,
    sillabb character varying(2),
    seguros character varying(2),
    cedula character varying(30),
    img character varying(3000),
    precio integer
);

-- Insertar 30 registros de ejemplo en la tabla autos
INSERT INTO autos (marca, modelo, combustible, transmision, lujos, sillas, sillabb, seguros, cedula, img, precio)
VALUES 
('Toyota', 2022, 'Gasolina', 'Automático', 'Si', 5, 'Si', 'Si', '0', 'https://www.elcarrocolombiano.com/wp-content/uploads/2021/06/20210618-TOYOTA-TUNDRA-TRD-PRO-2022-PRIMERA-FOTO-OFICIAL-01.jpg', 20000),
('Honda', 2021, 'Gasolina', 'Manual', 'No', 4, 'No', 'Si', '0', 'https://upload.wikimedia.org/wikipedia/commons/7/76/2021_Honda_Accord_Sport_%28facelift%29%2C_front_11.30.21.jpg', 18000),
('Ford', 2020, 'Diesel', 'Automático', 'Si', 6, 'Si', 'No', '0', 'https://www.elcarrocolombiano.com/wp-content/uploads/2019/01/20190110-FORD-EXPLORER-2020-CARACTERISTICAS-VERSIONES-Y-EQUIPAMIENTO-01.jpg', 25000),
('Chevrolet', 2019, 'Gasolina', 'Manual', 'No', 5, 'No', 'Si', '0', 'https://www.shutterstock.com/image-illustration/red-chevrolet-spark-20192021-model-600nw-2136526829.jpg', 17000),
('Nissan', 2022, 'Híbrido', 'Automático', 'Si', 5, 'Si', 'Si', '0', 'https://www.elcarrocolombiano.com/wp-content/uploads/2021/02/20210204-NISSAN-PATHFINDER-2022-PORTADA.jpg', 22000),
('BMW', 2021, 'Gasolina', 'Automático', 'Si', 4, 'Si', 'No', '0', 'https://www.elcarrocolombiano.com/wp-content/uploads/2021/11/20211116-BMW-SERIE-4-GRAN-COUPE-420I-COLOMBIA-PRECIO-CARACTERISTICAS-VERSIONES-01.jpg', 30000),
('Mercedes', 2020, 'Diesel', 'Manual', 'No', 5, 'No', 'Si', '0', 'https://www.shutterstock.com/image-illustration/paris-france-may-02-2022-260nw-2212939915.jpg', 35000),
('Audi', 2019, 'Gasolina', 'Automático', 'Si', 4, 'Si', 'No', '0', 'https://www.shutterstock.com/image-photo/guangzhou-china-september-82022-blue-600nw-2199585715.jpg', 28000),
('Volkswagen', 2022, 'Gasolina', 'Manual', 'No', 5, 'No', 'Si', '0', 'https://acnews.blob.core.windows.net/imgnews/medium/NAZ_55fdf7bbb9b4489595a0b71c7b0382d2.jpg', 19000),
('Hyundai', 2021, 'Híbrido', 'Automático', 'Si', 6, 'Si', 'No', '0', 'https://acnews.blob.core.windows.net/imgnews/large/NAZ_be8f07e516974b8ca8be219033210b87.jpg', 21000);

CREATE TABLE IF NOT EXISTS usuarios(
  cedula varchar(30) PRIMARY KEY,
  contrasena varchar(12)
);