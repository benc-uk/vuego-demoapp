import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import About from './views/About.vue'
import Goat from './views/Goat.vue'

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
      path: '/app/about',
      name: 'about',
      component: About
    },
    {
      path: '/app/goats',
      name: 'about',
      component: Goat
    }    
  ]
})
