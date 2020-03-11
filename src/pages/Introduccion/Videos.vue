<template>
  <v-flex>
        <menu-app></menu-app>
        <div class="contenedor">
          <h3>Videos</h3>
          <div class="contenido">
            <!-- Se definen etiquetas de aceurdo al critério de búsqueda y se solicita ingresar cada dato según corresponda -->
            <form class="buscador" action="index.html" method="post">
              <label for="TemaVideos">Tema:</label>
              <br />
              <input type="text" title="Buscar por tema en Videos" placeholder="Buscar" id="TemaVideos" v-model="form.tema" required>
              <br />
              <label for="ParticipanteVideos">Participante:</label>
              <br />
              <input type="text" title="Buscar por participante en Videos" placeholder="Ingrese un nombre" id="ParticipanteVideos" v-model="form.autor" required>
              <br />
              <label for="FechaVideos">Fecha:</label>
              <br />
              <!-- Etiqueta para el manejo de fechas por año -->

              <input class="text" id="FechaVideos" title="Buscar por fecha en Videos" placeholder="Ingrese un año" v-model="form.fecha">

              <br />
            </form>
            <!-- Envío orden para búsqueda -->
              <input type="submit" name="" value="Enviar" class="btn" id="botonvideos" @click="creartablaVideos()">
            <!-- Esquema de la tabla contenedora de la búsqueda para videos -->
            <br />
            <br />
           <table>
                <thead>
                  <th width="20">id</th>
                  <th width="40">tema</th>
                  <th width="50">autor</th>
                  <th width="30">fecha</th>
                  <th width="150">link</th>
                </thead>
                <tbody >
                  <tr v-for="(item,index) in misVideos" :key="index">
                    <th scope="row">{{item.id}}</th>
                    <th>{{item.tema}}</th>
                    <th>{{item.autor}}</th>
                    <th>{{item.fecha}}</th>
                    <th><iframe width="200" height="185" :src="item.link" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></th>
                  </tr>
              </tbody>
              </table>
          </div>
        </div>
    </v-flex>
</template>

<script>
export default {
  data() {
      return {
          misVideos:[],
        form: {
          tema: '',
          autor: '',
          fecha: ''
        },
        show: true
      }
    },
  methods:{
       creartablaVideos(){
         this.axios.get("http://localhost:3000/video/"+this.form.tema+","+this.form.autor+","+this.form.fecha)
        .then(res=>{
            this.misVideos=res.data;
        })
        .catch(e=>{
            alert("Dato no encontrado");
            alert(e.response);
        }) 
       }
}
}
</script>

<style>

</style>