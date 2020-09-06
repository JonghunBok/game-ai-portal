import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import seduVnc from 'vue-novnc-sedu'

Vue.config.productionTip = false
Vue.use(seduVnc);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
