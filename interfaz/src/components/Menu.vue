<template>
  <v-content>
    <v-app-bar-nav-icon class="Responsive" @click="drawer=!drawer"></v-app-bar-nav-icon>
    <v-toolbar app color="#BB681A" class="BarraNavegacion">
      <v-toolbar-title>
        <router-link to="./">
          <a class="ItemPrincipal">Buscador</a>
        </router-link>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <router-link to="./audios">
        <a class="Secundarios">Audios</a>
      </router-link>
      <router-link to="./videos">
        <a class="Secundarios">Videos</a>
      </router-link>
      <router-link to="./documentos">
        <a class="Secundarios">Documentos</a>
      </router-link>
      <router-link to="./publicaciones">
        <a class="Secundarios">Publicidad</a>
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
  </v-content>
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