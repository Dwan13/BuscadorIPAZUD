<template>
     <v-flex class="Presentacion">
       <!-- Se consume el componente menu -->
        <menu-app></menu-app>
        <div class="contenedor">
      <div class="contenedor-titulo">
        <div class="titulo">
          <h1>IPAZUD</h1>
          <h2>Instituto Universidad Distrital</h2>
          <!-- Enlace rápido con la página principal del IPAZUD -->
          <a href="https://ipazud.udistrital.edu.co/" target="_blank" title="Ir a la página oficial">Ver página oficial</a>
        </div>
      </div>
    </div>
           <!-- Se consume el componente footer -->
    <footer-app></footer-app>
     </v-flex>
</template>

<script>
export default {

}
</script>

<style lang="scss" scoped>
/* Se declaran variables para estilos de la portada */
@import "@/sass/Principal.scss";
</style>