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
              <input type="text" title="Buscar por tema en audio" placeholder="Ingrese un tema" id="TemaAudios" required>
              <br />
              <label for="ParticipanteAudios">Participante:</label>
              <br />
              <input type="text" title="Buscar por participante en audio" placeholder="Ingrese un nombre" id="ParticipanteAudios" required>
              <br />
              <!-- Etiqueta para el manejo de fechas por año -->

              <label for="FechaAudios">Fecha:</label>
              <br />
              <input class="text" id="FechaAudios" title="Buscar por fecha en audio" placeholder="Ingrese un año" required>

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
              <br />
              <br />
              <br />
              <br />
              <thead>
                <th width="20">id</th>
                <th width="40">tema</th>
                <th width="50">participantes</th>
                <th width="30">fecha</th>
                <th width="150">link</th>
              </thead>
              <tbody id="cuerpoAudios">
              </tbody>
            </table>
          </div>
        </div>

    </v-flex>
</template>

<script>
export default {
   /*  methods:{
       creartablaAudios(){
alert("Entre")
$("#cuerpoAudios tr").remove();

var temaAudio = document.getElementById("TemaAudios").value;
var participanteAudio = document.getElementById("ParticipanteAudios").value;
var añoAudio = document.getElementById("FechaAudios").value;
$.ajax({
    url     : "http://localhost:3000/audio/"+temaAudio+","+participanteAudio+","+añoAudio,
    type    : "GET",
    success : (function (data) {
                  console.log(temaAudio);
                  console.log(participanteAudio);
                  if(document.getElementById(data)==null){
                    alert("Datos no encontrados")
                  }
                  alert("Busqueda realizada con exito");
                    for(var i=0; i<data.length; i++){
                    var tr = `<tr>
                        <td>`+data[i].id+`</td>
                        <td>`+data[i].tema+`</td>
                        <td>`+data[i].participantes+`</td>
                        <td>`+data[i].fecha+`</td>
                        <td> <iframe width="150" height="150" scrolling="no" frameborder="no" allow="autoplay" src="`+data[i].link+`"></iframe></td>
                    </tr>`;
                    $("#cuerpoAudios").append(tr)
                    }
              })
});
} 
    } */
}
</script>

<style scoped>

#Audios .contenido{
  display: inline;
  justify-content: flex-start;
}

</style>