-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 17-03-2020 a las 01:02:55
-- Versión del servidor: 10.4.11-MariaDB
-- Versión de PHP: 7.2.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `pasantia`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `audios`
--

CREATE TABLE `audios` (
  `id` int(11) NOT NULL,
  `tema` varchar(100) NOT NULL,
  `participantes` varchar(255) NOT NULL,
  `fecha` int(4) NOT NULL,
  `link` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `audios`
--

INSERT INTO `audios` (`id`, `tema`, `participantes`, `fecha`, `link`) VALUES
(1, 'este audio es una muestra de progreso', 'ivan arevalo dwan veloza', 2018, 'https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/693822901&color=%23ff5500&auto_play=false&hide_related=false&show_comments=true&show_user=true&show_reposts=false&show_teaser=true&visual=true'),
(2, 'memoria y conflicto', 'personaje de prueba', 2019, 'https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/693822901&color=%23ff5500&auto_play=false&hide_related=false&show_comments=true&show_user=true&show_reposts=false&show_teaser=true&visual=true'),
(3, 'indigenas', 'ivan duque', 2019, 'https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/693822901&color=%23ff5500&auto_play=false&hide_related=false&show_comments=true&show_user=true&show_reposts=false&show_teaser=true&visual=true');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `documentos`
--

CREATE TABLE `documentos` (
  `id` int(11) NOT NULL,
  `tema` varchar(100) NOT NULL,
  `tipo` varchar(100) NOT NULL,
  `area` varchar(100) NOT NULL,
  `autor` varchar(100) NOT NULL,
  `fecha` int(4) NOT NULL,
  `link` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `documentos`
--

INSERT INTO `documentos` (`id`, `tema`, `tipo`, `area`, `autor`, `fecha`, `link`) VALUES
(1, 'tema de prueba', 'tesis', 'memoria', 'ivan arevalo', 2019, 'http://bienestar.udistrital.edu.co:8080/c/document_library/get_file?uuid=1648f261-4726-42eb-a08f-97f'),
(2, 'territorios indigenas', 'libro', 'democracia', 'dwan veloza', 2018, 'http://www.scielo.org.co/pdf/rhel/v19n28/v19n28a09.pdf'),
(3, 'paz', 'registro', 'territorio', 'dwan veloza', 2012, 'https://santillanaplus.com.co/pdf/cartilla-catedra-de-paz.pdf');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `publicidad`
--

CREATE TABLE `publicidad` (
  `id` varchar(11) NOT NULL,
  `tema` varchar(255) NOT NULL,
  `tipo` varchar(255) NOT NULL,
  `participantes` varchar(255) NOT NULL,
  `fecha` int(4) NOT NULL,
  `link` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `publicidad`
--

INSERT INTO `publicidad` (`id`, `tema`, `tipo`, `participantes`, `fecha`, `link`) VALUES
('1', 'el primer tema es un tema de prueba', 'folleto', 'ivan arevalo', 2019, 'https://www.google.com/url?sa=i&url=http%3A%2F%2Fwww.editoriallapaz.org%2Fsimiente-JPEG-folleto-imprimir.htm&psig=AOvVaw23AMQcUmHHU5GCHm20Gxnb&ust=1583989928542000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCOiw9t_UkegCFQAAAAAdAAAAABAJ'),
('2', 'anuncio paz conflicto', 'afiche', 'dwan veloza', 2019, 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.pinterest.com%2Fpin%2F492299803010917599%2F&psig=AOvVaw23AMQcUmHHU5GCHm20Gxnb&ust=1583989928542000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCOiw9t_UkegCFQAAAAAdAAAAABAa');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `videos`
--

CREATE TABLE `videos` (
  `id` int(11) NOT NULL,
  `tema` varchar(100) NOT NULL,
  `autor` varchar(100) NOT NULL,
  `fecha` int(4) NOT NULL,
  `link` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `videos`
--

INSERT INTO `videos` (`id`, `tema`, `autor`, `fecha`, `link`) VALUES
(1, 'video de prueba ipazud', 'ivan arevalo', 2019, 'https://www.youtube.com/embed/RIed1Olb_6s'),
(2, 'un tema diferente', 'dwan veloza', 2018, 'https://www.youtube.com/embed/RIed1Olb_6s'),
(3, 'latinoamerica', 'calle 13', 2019, 'https://www.youtube.com/embed/DkFJE8ZdeG8');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `audios`
--
ALTER TABLE `audios`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `documentos`
--
ALTER TABLE `documentos`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `publicidad`
--
ALTER TABLE `publicidad`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `videos`
--
ALTER TABLE `videos`
  ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
