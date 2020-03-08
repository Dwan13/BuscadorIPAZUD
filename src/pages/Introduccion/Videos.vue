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
              <input type="text" title="Buscar por tema en Videos" placeholder="Buscar" id="TemaVideos" required>
              <br />
              <label for="ParticipanteVideos">Participante:</label>
              <br />
              <input type="text" title="Buscar por participante en Videos" placeholder="Ingrese un nombre" id="ParticipanteVideos" required>
              <br />
              <label for="FechaVideos">Fecha:</label>
              <br />
              <!-- Etiqueta para el manejo de fechas por año -->

              <input class="text" id="FechaVideos" title="Buscar por fecha en Videos" placeholder="Ingrese un año">

              <br />
            </form>
            <!-- Envío orden para búsqueda -->
              <input type="submit" name="" value="Enviar" class="btn" id="botonvideos" @click="creartablaVideos()">
            <!-- Esquema de la tabla contenedora de la búsqueda para videos -->
            <br />
            <br />
           <table>
             <br />
             <br />
             <br />
             <br />
                <thead>
                  <th width="20">id</th>
                  <th width="40">tema</th>
                  <th width="50">autor</th>
                  <th width="30">fecha</th>
                  <th width="150">link</th>
                </thead>
                <tbody id="cuerpoVideos">
                </tbody>
              </table>
          </div>
        </div>
    </v-flex>
</template>

<script>
export default {
/* methods:{
    creartablaVideos(){

    $("#cuerpoVideos tr").remove();

    var temaVideo = document.getElementById("TemaVideos").value;
    var participanteVideo = document.getElementById("ParticipanteVideos").value;
    var añoVideo = document.getElementById("FechaVideos").value;
    $.ajax({
        url     : "http://localhost:3000/video/"+temaVideo+","+participanteVideo+","+añoVideo,
        type    : "GET",
        success : (function (data) {
          if(document.getElementById(data)==null){
                        alert("Datos no encontrados")
                      }
                      alert("Busqueda realizada con exito");
                        for(var i=0; i<data.length; i++){
                        var tr = `<tr>
                            <td>`+data[i].id+`</td>
                            <td>`+data[i].tema+`</td>
                            <td>`+data[i].autor+`</td>
                            <td>`+data[i].fecha+`</td>
                            <td><iframe width="200" height="185" src="`+data[i].link+`" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></td>
                        </tr>`;
                        $("#cuerpoVideos").append(tr)
                        }
                  })
    });
}
} */
}
</script>

<style>

</style>