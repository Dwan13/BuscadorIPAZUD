<template>
  <div>
    <!-- Uso de vuetify nav y toolbar para desarrollar el menu responsive
    1. drawer es una variable lógica que eprmite desplegar o no el menu en un entorno responsive
    2. La barra de navegación primaria es la que contiene la clase BarraNavegacion.
     -->
    <v-app-bar-nav-icon class="Responsive" @click="drawer=!drawer"></v-app-bar-nav-icon>
    <v-toolbar app color="#BB681A" class="BarraNavegacion">
      <v-toolbar-title>
        <router-link to="./">
          <a class="ItemPrincipal" title="Ir a la página principal">
            Buscador</a>
        </router-link>  
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <router-link to="./audios">
        <a class="Secundarios" title="Ir a la categoría Audios">Audios</a>
      </router-link>
      <router-link to="./videos">
        <a class="Secundarios" title="Ir a la categoría Videos">Videos</a>
      </router-link>
      <router-link to="./documentos">
        <a class="Secundarios" title="Ir a la categoría Documentos">Documentos</a>
      </router-link>
      <router-link to="./publicaciones">
        <a class="Secundarios" title="Ir a la categoría Publicidad">Publicidad</a>
      </router-link>
    </v-toolbar>
    <v-navigation-drawer app v-model="drawer" temporary color="#BB681A">
      <v-layout mt-4 ml-2 column align-center>
        <v-flex >
          <v-avatar>
              <router-link to="./">
              <img src="@/assets/LogoIPAZUD.jpg" alt="Logo Ipazud">
        </router-link>
          </v-avatar>
        </v-flex>
          <v-flex class="white--text mt-2 headline" >
              IPAZUD
          </v-flex>
          <v-divider></v-divider>
          <v-flex>
          <v-list-item class="Items">
            <v-list-item-content>
              <router-link to="./audios">
                <a>Audios</a>
              </router-link>
            </v-list-item-content>
            <v-list-item-content>
              <router-link to="./videos">
                <a>Videos</a>
              </router-link>
            </v-list-item-content>
            <v-list-item-content>
              <router-link to="./documentos">
                <a>Documentos</a>
              </router-link>
            </v-list-item-content>
            <v-list-item-content>
              <router-link to="./publicaciones">
                <a>Publicidad</a>
              </router-link>
            </v-list-item-content>
          </v-list-item>
        </v-flex>
      </v-layout>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      drawer: false
    };
  }
};
</script>

<style lang="scss" scoped>
/* Estilos de MENU
Display: none permite la interacción entre el responsive con drawer
 */
@media (min-width: 600px) {
    .Responsive {
      display: none;
}
.Secundarios{
    padding: 15px;
    color: white;
}
.Secundarios:hover{
  background: #2E8E83;
  color: #fff;
}
  }
@media (max-width: 600px) {
    .BarraNavegacion {
      display: none;
}
  }

.Items{
    list-style-type: square;
    display: inline;
   padding: 0 1em 0 0;
}
img{
    max-width: 100%;
    width: 100%;
}
.v-application a {
        color: white;
}

</style>