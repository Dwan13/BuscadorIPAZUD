<template>
  <!-- CLASE PUBLICACIONES
Creada por DWAN FELIPE VELOZA PAEZ
Tenga en cuenta la documentación de Vuetify: Grillas, botones y data-tables
Se implementa un buscador dinámico que permite seleccionar dos tipos de búsqueda: Por filtros y rápida
La búsqueda por filtros requiere de formularios para garantizar que se incluyan los criterios de búsqueda,
mientras que la búsqueda rápida trae todo el contenido de la tabla en la Base de Datos y realiza filtros
  mediante search.-->
  <v-container class="animated fadeInUp dura-1">
    <h3>PUBLICACIONES</h3>
    <!-- Botones de inicio -->
    <v-layout v-if="VerBotones" wrap mt-5 xs12 justify-center>
      <div class="Seleccion">
        <v-btn
          title="Búsqueda por Filtros"
          class="mt-5"
          color="success"
          @click="elegirFiltros"
        >Búsqueda Por Filtros</v-btn>
        <v-btn
          color="success"
          class="mt-5"
          @click="elegirBusquedaRapida"
          title="Búsqueda Rápida"
        >Búsqueda Rápida</v-btn>
      </div>
    </v-layout>
    <!-- Invocamos el modal para control de errores -->
    <modal-app v-if="showModal" @close="showModal=!showModal">
      <h3 slot="header">Publicación no encontrada</h3>
    </modal-app>
    <!-- Formulario para la búsqueda por filtros -->
    <v-layout v-if="VerFiltros" wrap class="animated fadeIn dura-1">
      <v-flex xs12>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field v-model="form.tema" :rules="temaRules" label="Tema" required></v-text-field>

          <v-text-field
            v-model="form.Participante"
            :rules="autorRules"
            label="Participante"
            required
          ></v-text-field>
          <v-flex xs6 md3>
            <v-select
              v-model="form.tipo"
              :items="form.select"
              :rules="[v => !!v || 'Item es requerido']"
              label="Tipo de Publicidad"
              required
            ></v-select>
          </v-flex>

          <v-text-field v-model="form.fecha" :rules="fechaRules" label="Fecha" required></v-text-field>

          <btn :disabled="!valid" class="mr-4" @click="creartablaPublicidad">Enviar</btn>

          <btn class="mr-4" @click="reset">Reestablecer Búsqueda</btn>
        </v-form>
      </v-flex>
      <v-flex mt-4>
        <v-card-title v-if="showSearch">
          <v-text-field
            v-model="search"
            append-icon="mdi-magnify"
            label="Search"
            single-line
            hide-details
          ></v-text-field>
        </v-card-title>
        <v-data-table :headers="headers" :items="misPublicaciones" :search="search">
          <template v-slot:item.link="{ item }">
            <iframe :src="item.link" scrolling="auto" width="550" height="450"></iframe>
          </template>
        </v-data-table>
      </v-flex>
    </v-layout>
    <!-- Búsqueda Rápida -->

    <v-flex v-if="Fast" mt-4 class="animated fadeIn dura-1">
      <v-card-title>
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table :headers="headers" :items="misPublicaciones" :search="search">
        <template v-slot:item.link="{ item }">
          <iframe :src="item.link" scrolling="auto" width="550" height="450"></iframe>
        </template>
      </v-data-table>
    </v-flex>
  </v-container>
</template>

<script>
export default {
  /* En esta sección declaramos variables que permiten cambiar de tipo de búsqueda
  tambien se incluyen lo v-model necesarios para generar la búsqueda más adelante. */
  data() {
    return {
      VerBotones: true,
      Fast: false,
      VerFiltros: false,
      showSearch: false,
      headers: [
        {
          text: "Id",
          align: "start",
          sortable: false,
          value: "id"
        },
        { text: "Tema", value: "tema" },
        { text: "Tipo", value: "tipo" },
        { text: "Participante", value: "participantes" },
        { text: "Fecha", value: "fecha" },
        { text: "Link", value: "link" }
      ],
      search: "",
      valid: true,
      showModal: false,
      temaRules: [v => !!v || "Tema es requerido"],
      autorRules: [v => !!v || "Autor es requerido"],
      fechaRules: [v => !!v || "Fecha es requerida"],
      misPublicaciones: [],
      form: {
        tema: "",
        tipo: "",
        select: ["Afiche", "Folleto", "Separador", "Flyer"],
        Participante: "",
        fecha: ""
      },
      show: true
    };
  },
  methods: {
    /* Los métodos creartablaAudios y  elegirBusquedaRapida requieren de Axios para generar búsquedas por método GET
    Una vez se despliegue el servidor en la red de datos con una ip pública se debe copiar y reemplazar por el localhost*/
    creartablaPublicidad() {
      this.axios
        .get(
          "http://localhost:3000/publicidad/" +
            this.form.tema +
            "," +
            this.form.tipo +
            "," +
            this.form.Participante +
            "," +
            this.form.fecha
        )
        .then(res => {
          this.misPublicaciones = res.data;
          this.showSearch = true;
        })
        .catch(e => {
          this.showModal = true;
          console(e.response);
        });
    },
    /* Limpiamos campos del formulario */
    reset() {
      this.$refs.form.reset();
    },
    /* Método para solicitar una búsqueda por filtros */
    elegirFiltros() {
      this.VerBotones = false;
      this.VerFiltros = true;
    },
    elegirBusquedaRapida() {
      this.VerBotones = false;
      this.Fast = true;
      this.axios
        .get("http://localhost:3000/publicaciones")
        .then(res => {
          this.misPublicaciones = res.data;
        })
        .catch(e => {
          this.showModal = true;
          console(e.response);
        });
    }
  }
};
</script>