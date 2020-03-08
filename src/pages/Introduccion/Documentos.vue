<template>
  <v-flex>
        <menu-app></menu-app>
        <div class="contenedor">
          <h3>Documentos</h3>
          <div class="contenido">
            <!-- Se definen etiquetas de aceurdo al critério de búsqueda y se solicita ingresar cada dato según corresponda -->
            <form class="buscador" action="index.html" method="post">
            <label for="Documento" >Tematica del documento:</label>
            <br />
            <input type="text" name="Documento" title="Buscar por temática del documento" placeholder="Ingrese palabras clave" id="temaDocumento" required>
            <br />
            <label for="Autor" >Ingrese autor:</label>
            <br />
            <input type="text" name="Autor" title="Buscar por autor del documento" placeholder="Ingrese un nombre" id="autorDocumento" required>
            <br />
            <label for="Tipo" >Tipo de documento:</label>
            <br />
            <!-- Se implementa un elemento de tipo select para escojer un tipo de documento -->
            <select class="desplegar" name="Tipo" id="tipoDocumento" title="Buscar por tipo de documento" >
              <option value="libro">Libro</option>
              <option value="revista">Revista</option>
              <option value="articulo">Artículo</option>
              <option value="tesis">Tesis</option>
              <option value="informe">Informe</option>
              <option value="entrevista">Entrevista</option>
              <option value="registro">Registro</option>
            </select>
            <br />
            <label for="Area" >Área de estudio:</label>
            <br />
            <!-- Se implementa un elemento de tipo select para escojer entre diferentes areas de estudio -->
            <select class="desplegar" name="Area" id="areaDocumento" title="Buscar por area temática del documento" >
              <option value="democracia">Democracia y Ciudadanía </option>
              <option value="memoria">Memoria y Conflicto</option>
              <option value="territorio">Territorio y Desarraigo</option>
            </select>
            <br />
            <!-- Etiqueta para el manejo de fechas por año -->
            <label for="FechaDocumentos">Fecha:</label>
            <br />
            <input class="text" id="fechaDocumento" title="Buscar por fecha del documento" placeholder="Ingrese un año">
            <br />
            <br />
          </form>
          <!-- Envío orden para búsqueda -->
          <input type="submit" name="" value="Enviar" class="btn" @click="creartablaDocumentos()">
            <br />
            <br />
          <!-- Esquema de la tabla contenedora de la búsqueda para documentos -->
          <table>
            <br />
            <br />
            <br />
            <br />
            <thead>
              <th width="20">id</th>
              <th width="40">tema</th>
              <th width="30">tipo</th>
              <th width="30">area</th>
              <th width="50">autor</th>
              <th width="30">fecha</th>
              <th width="150">link</th>
            </thead>
            <tbody id="cuerpoDocumentos">
            </tbody>
          </table>
          </div>
        </div>
    </v-flex>
</template>

<script>
export default {
/* methods:{
    creartablaDocumentos(){

    $("#cuerpoDocumentos tr").remove();

    var temaDocumento = document.getElementById("temaDocumento").value;
    var autorDocumento = document.getElementById("autorDocumento").value;
    var areaDocumento = document.getElementById("areaDocumento").value;
    var tipoDocumento = document.getElementById("tipoDocumento").value;
    var añoDocumento = document.getElementById("fechaDocumento").value;
    $.ajax({
        url     : "http://localhost:3000/documento/"+temaDocumento+","+tipoDocumento+","+areaDocumento+","+autorDocumento+","+añoDocumento,
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
                            <td>`+data[i].tipo+`</td>
                            <td>`+data[i].area+`</td>
                            <td>`+data[i].autor+`</td>
                            <td>`+data[i].fecha+`</td>
                            <td>`+data[i].link+`</td>
                        </tr>`;
                        $("#cuerpoDocumentos").append(tr)
                        }
                  })
    });
}
} */
}
</script>

<style>
#Documentos{
  padding: 90px 0;
}
 #Documentos h3{
   text-align:  center;
   font-size: 45px;
   font-weight: normal;
 }
</style>