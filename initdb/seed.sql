-- sql yang akan dijalankan saat container appDb started,
-- contoh:

CREATE DATABASE IF NOT EXISTS contoh;

CREATE TABLE IF NOT EXISTS contoh.orang(
    nama varchar(20),
    alamat varchar(20)
);