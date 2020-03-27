<template>
<v-layout>
  <!-- Pantalla de inicio -->
     <v-flex class="Presentacion">
        <div class="contenedor">
      <div class="contenedor-titulo">
        <div class="titulo" >
          <h1 class="animated fadeInLeft dura-2">IPAZUD</h1>
          <h2 class="animated fadeInLeft dura-2">Instituto Universidad Distrital</h2>
          <!-- Enlace rápido con la página principal del IPAZUD -->
          <a class="animated fadeInLeft dura-2" href="https://ipazud.udistrital.edu.co/" target="_blank" title="Ir a la página oficial">Ver página oficial</a>
        </div>
      </div>
    </div>
    <!-- Llamamos el componente Footer -->
    <footer-app class="animated fadeInUp dura-1"></footer-app>
     </v-flex>
     </v-layout>
</template>

<style lang="scss" scoped>
/* Importamos clase de estilo sass Principal.scss */
@import "@/sass/Principal.scss";
</style>