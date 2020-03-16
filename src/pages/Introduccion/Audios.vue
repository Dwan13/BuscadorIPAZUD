<template>
    <v-flex id="Audios">
        <menu-app></menu-app>
        <div class="contenedor">
          <h3>Audios</h3>
          <div class="contenido">
            <!-- Se definen etiquetas de aceurdo al critério de búsqueda y se solicita ingresar cada dato según corresponda -->
            <form class="buscador" action="index.html" method="post">
              <label for="TemaAudios">Tema:</label>
              <br />
              <input type="text" title="Buscar por tema en audio" placeholder="Ingrese un tema" id="TemaAudios" v-model="form.tema" required>
              <br />
              <label for="ParticipanteAudios">Participante:</label>
              <br />
              <input type="text" title="Buscar por participante en audio" placeholder="Ingrese un nombre" id="ParticipanteAudios" v-model="form.autor" required>
              <br />
              <!-- Etiqueta para el manejo de fechas por año -->

              <label for="FechaAudios">Fecha:</label>
              <br />
              <input class="text" id="FechaAudios" title="Buscar por fecha en audio" placeholder="Ingrese un año" v-model="form.fecha" required>

              <br />
              <br />

            </form>
            <br />
            <br />
            <br />
            <!-- Envío orden para búsqueda -->
            <input type="submit" name="" value="Enviar" class="btn" @click="creartablaAudios()">
            <br />
            <br />  
            <!-- Esquema de la tabla contenedora de la búsqueda para audios -->
            <table>
              <thead>
                <th width="20">id</th>
                <th width="40">tema</th>
                <th width="50">participantes</th>
                <th width="30">fecha</th>
                <th width="150">link</th>
              </thead>
              <tbody >
                  <tr v-for="(item,index) in misAudios" :key="index">
                    <th scope="row">{{item.id}}</th>
                    <th>{{item.tema}}</th>
                    <th>{{item.participantes}}</th>
                    <th>{{item.fecha}}</th>
                    <th><iframe width="150" height="150" scrolling="no" frameborder="no" allow="autoplay" :src="item.link"></iframe></th>
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
          misAudios:[],
        form: {
          tema: '',
          autor: '',
          fecha: ''
        },
        show: true
      }
    },
  methods:{
       creartablaAudios(){
         this.axios.get("https://cors-anywhere.herokuapp.com/http://10.20.200.180:3000/audio/"+this.form.tema+","+this.form.autor+","+this.form.fecha)
        .then(res=>{
            this.misAudios=res.data;
        })
        .catch(e=>{
            alert("Dato no encontrado"+" "+e.response);
        }) 
       }
}
}
</script>

<style>

#Audios .contenido{
  display: inline;
  justify-content: flex-start;
}


</style>