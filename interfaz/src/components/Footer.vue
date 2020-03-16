<template>
   <footer>
        <div class="contenedor">
          <div class="redes">
            <a href="https://www.facebook.com/ipazud/" target="_blank"><img src='../assets/facebook.png' alt="imagen de facebook"></a>
            <a href="https://twitter.com/InstitutoPazUD" target="_blank"><img src='../assets/twitter.png' width="1.6%" alt="imagen de twitter"></a>
            <a href="https://www.instagram.com/institutodepaz_ud/ " target="_blank"><img src='../assets/instagram.png' width="1.6%" alt="imagen de instagram"></a>
          </div>
          <div class="parrafo">
            <p>© 2019 Instituto para la Pedagogia, la Paz y el Conflicto Urbano IPAZUD  | Universidad Distrital Francisco José de Caldas. NIT. 899.999.230.7</p>
            <p>Prueba 3</p>
          </div>
        </div>
      </footer>
</template>

<script>
export default {

}
</script>

<style>
footer {
  background: #2F2F2F;
  padding: 20px 0;
  text-align: center;
  margin-top: 13%;
  position: relative;
  bottom: 0;
}

footer .contenedor {
  display: flex;
  justify-content: center;
  flex-direction: column;
}

footer img {
  margin: 0 40px;
  transform: rotate(0);
  transition: 0.4s;
}

footer img:hover {
  transform: rotate(360deg);
  transition: 0.4s;
}

footer p {
  color: #fff;
  margin-top: 1%;
}
</style>