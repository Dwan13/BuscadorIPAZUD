import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    imagen:"",
    numeroPiel:"",
    numeroRopa:"",
    numeroPelo:"",
    numeroOjos:""
  },
  mutations: {
    avatar(state, valor){
      state.imagen=valor
    },
    numeroPiel(state, number){
      state.numeroPiel=number
    },
    numeroRopa(state, nRopa){
      state.numeroRopa=nRopa
    },
    numeroPelo(state, nPelo){
      state.numeroPelo=nPelo
    },
    numeroOjos(state, nOjos){
      state.numeroOjos=nOjos
    }
  },
  actions: {
  },
  modules: {
  }
})
