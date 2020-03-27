<template>
   <footer>
   <!-- Contenido del FOOTER -->
        <v-layout class="contenedor" wrapp>
          <!-- Enlazo a travez de etiquetas "a" las redes sociales que maneja el IPAZUD -->
          <v-flex class="redes">
            <a title="Facebook" href="https://www.facebook.com/ipazud/" target="_blank"><img class="GirarIcon" src='../assets/facebook.png' alt="imagen de facebook"></a>
            <a title="Twitter" href="https://twitter.com/InstitutoPazUD" target="_blank"><img class="GirarIcon" src='../assets/twitter.png'  alt="imagen de twitter"></a>
            <a title="Instagram" href="https://www.instagram.com/institutodepaz_ud/ " target="_blank"><img class="GirarIcon" src='../assets/instagram.png'  alt="imagen de instagram"></a>
          </v-flex>
          <!-- Se define la clase parrafo para identificar aspectos de la Universidad y el IPAZUD -->
          <v-flex class="parrafo">
            <p>© 2020 Instituto para la Pedagogia, la Paz y el Conflicto Urbano IPAZUD  | Universidad Distrital Francisco José de Caldas. NIT. 899.999.230.7</p>
            <p>Código Postal: 11021-110231588</p>
            <a   title="Ubicación" id="DatosContacto" target="_blank" href="https://www.google.com.co/maps/place/Cra.+7+%2340b-53,+Bogot%C3%A1/@4.6280856,-74.0674698,17z/data=!3m1!4b1!4m5!3m4!1s0x8e3f9a287591013f:0x5cce5fbab6b77b9b!8m2!3d4.6280856!4d-74.0652811?hl=es"><i class="fa fa-map-marker" aria-hidden="true"></i><img img src='../assets/mapa.png' alt="imagen de GoogleMaps">Carrera 7 No. 40B - 53 <br>Bogotá D.C. - República de Colombia</a>
            <br>
            <a  id="DatosContacto" href="https://www.udistrital.edu.co/directorio" title="Directorio institucional"><i class="fa fa-phone" aria-hidden="true"></i>Directorio institucional</a>
          </v-flex>
        </v-layout>
      </footer>
</template>

<style scoped>
/* Estilos del FOOTER */
footer {
  position: absolute;
  bottom: 0;
  justify-content: center;
  background: #2F2F2F;
  padding: 1% 25%;
  text-align: center;
  bottom: 0; 
}


footer .contenedor {
  display: flex;
  justify-content: center;
  flex-direction: column;
}

footer .GirarIcon {
  margin: 0 40px;
  transform: rotate(0);
  transition: 0.4s;
}
footer img{
  margin: 0 10px;
}
footer .GirarIcon:hover {
  transform: rotate(360deg);
  transition: 0.4s;
}

footer p {
  color: #fff;
  margin-top: 1%;
}
#DatosContacto{
  color: rgb(238, 207, 29);
}
@media (max-width: 600px) {
    footer {
      padding: 1% 5%;
}
footer .GirarIcon{
  margin: 0 15px;
}}

</style>