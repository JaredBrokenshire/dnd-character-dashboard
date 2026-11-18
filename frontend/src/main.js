import Vue from 'vue'
import App from '@/App.vue'
import store from '@/store'
import router from '@/router'

// Import Styles
import "@/assets/css/styles.css"

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
