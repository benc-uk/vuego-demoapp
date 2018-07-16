//
// Main starting point for Vue.js SPA
// ----------------------------------------------
// Ben C, April 2018
//

import Vue from 'vue'
import App from './App.vue'
import router from './router'

// UI and Bootstrap stuff
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import './scss/theme.scss'

// Font Awesome has Vue.js support, import some icons we'll use
import { library as fontAwesomeLib } from '@fortawesome/fontawesome-svg-core'
import { faHome, faCogs, faTachometerAlt, faInfoCircle } from '@fortawesome/free-solid-svg-icons'
import { faGithub, faDocker } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
// Register icons and component
fontAwesomeLib.add([faHome, faCogs, faTachometerAlt, faInfoCircle, faGithub, faDocker])
Vue.component('font-awesome-icon', FontAwesomeIcon)

// Init Vue 
Vue.use(BootstrapVue);
Vue.config.productionTip = false

// Root Vue instance
// Mount on the <div id="root"> and render the template of the App component
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
