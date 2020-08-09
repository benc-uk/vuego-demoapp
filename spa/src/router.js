//
// App router
// ----------------------------------------------
// Ben C, April 2018
//

import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import About from './components/About.vue'
import Error from './components/Error.vue'
import Info from './components/Info.vue'
import Monitor from './components/Monitor.vue'
import Weather from './components/Weather.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
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
      path: '*',
      name: 'catchall',
      component: Error
    }
  ]
})
