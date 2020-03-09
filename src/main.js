import Vue from 'vue'
import App from './App.vue'
import router from './js/router/index'
// import './plugins/vuetify'
import axios from 'axios'
import VueAxios from 'vue-axios'
import store from './js/store/index'
import Menu from "@/components/Menu.vue"
Vue.component("menu-app", Menu);

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