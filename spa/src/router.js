import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import About from './components/About.vue'
import Info from './components/Info.vue'

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
      path: '/app/about',
      name: 'about',
      component: About
    }    
  ]
})
