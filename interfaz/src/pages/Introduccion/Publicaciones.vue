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
              <input type="text" size="46" maxlength="256" id="temaPublicidad" name="temaPublicidad" title="Buscar por tema de publicación" placeholder="Igrese una palabra clave" v-model="form.tema">
              <br />
              <label for="Participante">Participante:</label>
              <br />
              <input type="text" title="Buscar por participtante en la publicación" placeholder="Ingrese un nombre" id="participantePublicidad" v-model="form.Participante" required>
              <br />
              <label for="Tipo_Publicidad" >Tip de Publicación</label>
              <!-- Se implementa un elemento de tipo select para escojer el tipo de publicación -->
              <select class="desplegar" name="Tipo_Publicidad" title="Buscar por tipo de publicación" id="tipoPublicidad" v-model="form.tipo">
                <option value="afiche">Afiche</option>
                <option value="folleto">Folleto</option>
                <option value="separador">Separador</option>
                <option value="flyer">Flyer</option>
              </select>
              <br />
              <!-- Etiqueta para el manejo de fechas por año -->
              <label for="FechaPublica">Fecha:</label>
              <br />
              <input class="text" id="fechaPublicidad" v-model="form.fecha" title="Buscar por fehca de la publicación"  placeholder="Ingrese un año">
              <br />
          </form>

          <!-- Envío orden para búsqueda -->
          <input type="submit" name="" value="Enviar" class="btn" @click="creartablaPublicidad()">
            <br />
            <br />
          <!-- Esquema de la tabla contenedora de la búsqueda para publicaciones -->
          <table>
            <thead>
              <th width="20">id</th>
              <th width="40">tema</th>
              <th width="30">tipo</th>
              <th width="50">participantes</th>
              <th width="30">fecha</th>
              <th width="150">link</th>
            </thead>
            <tbody >
                  <tr v-for="(item,index) in misPublicaciones" :key="index">
                    <th scope="row">{{item.id}}</th>
                    <th>{{item.tema}}</th>
                    <th>{{item.tipo}}</th>
                    <th>{{item.participantes}}</th>
                    <th>{{item.fecha}}</th>
                    <th><iframe :src="item.link" scrolling="auto"></iframe></th>
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
          misPublicaciones:[],
        form: {
          tema: '',
          tipo: '',
          Participante: '',
          fecha: ''
        },
        show: true
      }
    },
  methods:{
       creartablaPublicidad(){
         this.axios.get("http://localhost:3000/publicidad/"+this.form.tema+","+this.form.tipo+","+this.form.Participante+","+this.form.fecha)
        .then(res=>{
            this.misPublicaciones=res.data;
        })
        .catch(e=>{
            alert("Dato no encontrado");
            alert(e.response);
        }) 
       }
}
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