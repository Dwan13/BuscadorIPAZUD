<template>
  <v-flex>
        <menu-app></menu-app>
        <div class="contenedor">
          <h3>Publicaciones</h3>
          <div class="form-group">
            <!-- Se definen etiquetas de aceurdo al critério de búsqueda y se solicita ingresar cada dato según corresponda -->
            <form class="buscador" action="index.html" method="post">
              <label for="searchterm1"> Buscar: </label>
              <br />
              <input type="text" size="46" maxlength="256" id="temaPublicidad" name="temaPublicidad" title="Buscar por tema de publicación" placeholder="Igrese una palabra clave">
              <br />
              <label for="Participante">Participante:</label>
              <br />
              <input type="text" title="Buscar por participtante en la publicación" placeholder="Ingrese un nombre" id="participantePublicidad" required>
              <br />
              <label for="Tipo_Publicidad" >Tip de Publicación</label>
              <!-- Se implementa un elemento de tipo select para escojer el tipo de publicación -->
              <select class="" name="Tipo_Publicidad" title="Buscar por tipo de publicación" id="tipoPublicidad">
                <option value="afiche">Afiche</option>
                <option value="folleto">Folleto</option>
                <option value="separador">Separador</option>
                <option value="flyer">Flyer</option>
              </select>
              <br />
              <!-- Etiqueta para el manejo de fechas por año -->
              <label for="FechaPublica">Fecha:</label>
              <br />
              <input class="text" id="fechaPublicidad" title="Buscar por fehca de la publicación"  placeholder="Ingrese un año">
              <br />
          </form>

          <!-- Envío orden para búsqueda -->
          <input type="submit" name="" value="Enviar" class="btn" onclick="creartablaPublicidad()">
            <br />
            <br />
          <!-- Esquema de la tabla contenedora de la búsqueda para publicaciones -->
          <table>
            <br />
            <br />
            <br />
            <br />
            <thead>
              <th width="20">id</th>
              <th width="40">tema</th>
              <th width="30">tipo</th>
              <th width="50">participantes</th>
              <th width="30">fecha</th>
              <th width="150">link</th>
            </thead>
            <tbody id="cuerpoPublicidad">
            </tbody>
          </table>
          </div>
        </div>
    </v-flex>
</template>

<script>
export default {
/* methods:{
    creartablaPublicidad(){

    $("#cuerpoPublicidad tr").remove();

    var temaPublicidad = document.getElementById("temaPublicidad").value;
    var participantePublicidad = document.getElementById("participantePublicidad").value;
    var tipoPublicidad = document.getElementById("tipoPublicidad").value;
    var añoPublicidad = document.getElementById("fechaPublicidad").value;
    $.ajax({
        url     : "http://localhost:3000/publicidad/"+temaPublicidad+","+tipoPublicidad+","+participantePublicidad+","+añoPublicidad,
        type    : "GET",
        success : (function (data) {
          if(document.getElementById(data)==null){
                        alert("Datos no encontrados")
                      }
                      try {
                            for(var i=0; i<data.length; i++){
                            var tr = `<tr>
                                <td>`+data[i].id+`</td>
                                <td>`+data[i].tema+`</td>
                                <td>`+data[i].tipo+`</td>
                                <td>`+data[i].participantes+`</td>
                                <td>`+data[i].fecha+`</td>
                                <td>`+data[i].link+`</td>
                            </tr>`;
                            $("#cuerpoPublicidad").append(tr)
                            }
                            alert("Busqueda realizada con exito");
                          }
                          catch (e) {
                            if (e instanceof ReferenceError) {
                                // Handle error as necessary
                                alert("No hay resultados");
                            }


                          }

                  })
    });
}
} */
}
</script>

<style scoped>
#Publicaciones{
   padding: 90px 0;
 }
 #Publicaciones h3{
   text-align: center;
   font-size: 45px;
   font-weight: normal;
 }
 #Publicaciones img{
   transform: scale(1);
   transition: .3;
   width:"100";
   height:"100";
 }
 #Publicaciones img:hover{
   transform: scale(1.1);
   transition:.3;
 }
 #Publicaciones p{
   text-align: center;
 }
 #Publicaciones a{
  text-decoration:none;
 }
</style>