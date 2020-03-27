/* Onjetos de cada vista pasados por parámetro mediante el método de Vue para comunicar pantallas llamado Router */
import Vue from 'vue'
import VueRouter from 'vue-router'
import Portada from '../../pages/portada.vue'
import Audios from '../../pages/Introduccion/Audios'
import Videos from '../../pages/Introduccion/Videos'
import Documentos from '../../pages/Introduccion/Documentos'
import Publicaciones from '../../pages/Introduccion/Publicaciones'

Vue.use(VueRouter)

const routes = [{
        path: '/audios',
        /* Dirección asosiada a un puerto */
        name: 'audios',
        /* Nombre de la clase */
        component: Audios /* Paso de parámetros desde import */
    },
    {
        path: '/videos',
        name: 'videos',
        component: Videos
    },
    {
        path: '/documentos',
        name: 'documentos',
        component: Documentos
    },
    {
        path: '/publicaciones',
        name: 'publicaciones',
        component: Publicaciones
    },
    {
        path: '/',
        name: 'portada',
        component: Portada
    },
]


const router = new VueRouter({
    routes
})

export default router