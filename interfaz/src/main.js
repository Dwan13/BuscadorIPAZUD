/* CLASSE main.js
Esta clase permite el uso e interacción entre componentes
*/
import Vue from 'vue'
import App from './App.vue'
import router from './js/router/index'
// import './plugins/vuetify'
import axios from 'axios'
import VueAxios from 'vue-axios'
import store from './js/store/index'
/* Declaramos el componente Menu */
import Menu from "@/components/Menu.vue"
Vue.component("menu-app", Menu);
/* Declaramos el componente Modal */
import modal from "@/components/modal.vue"
Vue.component("modal-app", modal);
/* Declaramos el componente Footer */
import Footer from "@/components/Footer.vue"
import vuetify from './plugins/vuetify';
Vue.component("footer-app", Footer);

Vue.use(VueAxios, axios)

Vue.config.productionTip = false


new Vue({
    router,
    store,
    vuetify,
    render: h => h(App)
}).$mount('#app')