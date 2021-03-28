//
// Main starting point for Vue.js SPA
// ----------------------------------------------
// Ben C, Sept 2018
//

import Vue from 'vue'
import App from './App.vue'
import router from './router'
import auth from './services/auth'

// UI and Bootstrap stuff
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import './scss/theme.scss'

// Font Awesome has Vue.js support, import some icons we'll use
import { library as fontAwesomeLib } from '@fortawesome/fontawesome-svg-core'
import {
  faHome,
  faCogs,
  faTachometerAlt,
  faInfoCircle,
  faUmbrella,
  faBomb,
  faCommentDollar,
  faLaptopCode,
  faCube,
  faDharmachakra,
  faMicrochip,
  faWrench,
  faMemory,
  faFlask,
  faProjectDiagram,
  faUser,
  faSignInAlt
} from '@fortawesome/free-solid-svg-icons'
import { faGithub, faDocker } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
// Register icons and component
fontAwesomeLib.add([
  faHome,
  faCogs,
  faTachometerAlt,
  faInfoCircle,
  faGithub,
  faDocker,
  faUmbrella,
  faBomb,
  faCommentDollar,
  faLaptopCode,
  faCube,
  faDharmachakra,
  faMicrochip,
  faWrench,
  faMemory,
  faFlask,
  faProjectDiagram,
  faUser,
  faSignInAlt
])
// eslint-disable-next-line
Vue.component('fa', FontAwesomeIcon)

// We have to register this globally for some reason
import VueSkycons from 'vue-skycons'
Vue.component('Skycon', VueSkycons)

// Init Vue
Vue.use(BootstrapVue)
Vue.config.productionTip = false

// Let's go!
appStartup()

//
// App start up synchronized using await with the config API call
//
async function appStartup() {
  // Take Azure AD client-id from .env.development or .env.development.local
  // Or fall back to empty string which disables the auth feature
  let AUTH_CLIENT_ID = process.env.VUE_APP_AUTH_CLIENT_ID || ''

  // Load config at runtime from special `/config` endpoint on frontend-host
  const apiEndpoint = process.env.VUE_APP_API_ENDPOINT || '/api'
  try {
    let configResp = await fetch(`${apiEndpoint}/config`)
    if (configResp.ok) {
      const config = await configResp.json()
      AUTH_CLIENT_ID = config.AUTH_CLIENT_ID
      console.log('### Config loaded from server API:', config)
    }
  } catch (err) {
    console.warn(
      `### Failed to fetch remote '${apiEndpoint}' endpoint. Local value for AUTH_CLIENT_ID '${AUTH_CLIENT_ID}' will be used`
    )
  }

  // Setup auth helper but disable dummy user
  auth.configure(AUTH_CLIENT_ID, false)

  // Actually mount & start the Vue app, kinda important
  new Vue({
    router,
    render: (h) => h(App)
  }).$mount('#app')
}
