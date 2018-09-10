//
// App router
// ----------------------------------------------
// Ben C, April 2018
//

import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import About from './components/About.vue'
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
      path: '/app/home',
      name: 'apphome',
      component: Home
    },
    {
      path: '/app/info',
      name: 'info',
      component: Info
    },
    {
      path: '/app/monitor',
      name: 'monitor',
      component: Monitor
    },    
    {
      path: '/app/weather',
      name: 'weather',
      component: Weather
    },   
    {
      path: '/app/about',
      name: 'about',
      component: About
    }    
  ]
})
