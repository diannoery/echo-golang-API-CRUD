CREATE DATABASE pegawai;
use pegawai

CREATE TABLE `pegawai`.`data_pegawai`
(
  `id_pegawai` INT NOT NULL,
  `nama` VARCHAR
(45) NULL,
  `alamat` VARCHAR
(45) NULL,
  `hp` INT NULL,
  PRIMARY KEY
(`id_pegawai`));

CREATE TABLE `pegawai`.`user` (
  `iduser` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  PRIMARY KEY (`iduser`));
