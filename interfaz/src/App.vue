<template>
 <!-- APP-Buscador Ipazud
Creada por DWAN FELIPE VELOZA PAEZ e IVAN RENÉ ARÉVALO VENEGAS
Este es el buscador del IPAZUD basado en vue.js y go.-->
<v-app>
    <!-- Llamamos al componente Menu -->
      <menu-app></menu-app>
      <v-content>
          <!-- router-view permite que naveguemos entre pantallas del buscador -->
          <router-view/>
      </v-content>
</v-app>
</template>

<style lang="scss">
/* Importamos estilos globales */
@import"sass/globales.scss";
</style>
