//
// App router
// ----------------------------------------------
// Ben C, April 2018, Updated for Vue3 2021
//

import Home from './components/Home.vue'
import About from './components/About.vue'
import Error from './components/Error.vue'
import Info from './components/Info.vue'
import Monitor from './components/Monitor.vue'
import Weather from './components/Weather.vue'
import User from './components/User.vue'
import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/home',
      name: 'apphome',
      component: Home
    },
    {
      path: '/info',
      name: 'info',
      component: Info
    },
    {
      path: '/monitor',
      name: 'monitor',
      component: Monitor
    },
    {
      path: '/weather',
      name: 'weather',
      component: Weather
    },
    {
      path: '/about',
      name: 'about',
      component: About
    },
    {
      path: '/user',
      name: 'user',
      component: User
    },
    {
      path: '/:catchAll(.*)',
      name: 'catchall',
      component: Error
    }
  ]
})

export default router
