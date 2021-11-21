//
// Main starting point for Vue.js SPA
// ----------------------------------------------
// Ben C, April 2018, Updated for Vue3 2021
//

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import auth from './services/auth'

// Bootstrap and icons
import 'bootswatch/dist/vapor/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.bundle'
import '@fortawesome/fontawesome-free/js/all'
// Weather Icons
import 'open-weather-icons/dist/css/open-weather-icons.css'
import 'open-weather-icons/dist/fonts/OpenWeatherIcons.svg'

const app = createApp(App)
app.use(router)

// Let's go!
startup(app)

// Default config
let config = {
  AUTH_CLIENT_ID: null,
  WEATHER_ENABLED: false
}

//
// App start up synchronized using await with the config API call
//
async function startup(app) {
  // Take Azure AD client-id from .env.development or .env.development.local if it's set
  // Fall back to empty string which disables the auth feature
  let AUTH_CLIENT_ID = process.env.VUE_APP_AUTH_CLIENT_ID || ''

  // Load config at runtime from special `/config` endpoint on Go server backend
  const apiEndpoint = process.env.VUE_APP_API_ENDPOINT || '/api'
  try {
    const configResp = await fetch(`${apiEndpoint}/config`)
    if (configResp.ok) {
      config = await configResp.json()
      AUTH_CLIENT_ID = config.authClientId
      console.log('### Config loaded from server API:', config)
    }
  } catch (err) {
    console.warn(`### Failed to fetch remote '${apiEndpoint}' endpoint. Local value for AUTH_CLIENT_ID '${AUTH_CLIENT_ID}' will be used`)
  }

  // Setup auth helper but disable dummy user
  // if AUTH_CLIENT_ID isn't set at this point, then the user sign-in will be dynamically disabled
  auth.configure(AUTH_CLIENT_ID, false)

  app.mount('#app')
}

export { config }
