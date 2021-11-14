<!--
Core app component
----------------------------------------------
// Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
      <div class="container-fluid">
        <a class="navbar-brand text-light" href="#">Vue &amp; Go Demo</a>
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div id="navbarSupportedContent" class="collapse navbar-collapse">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <router-link class="nav-link active" to="/home"> <i class="fas fa-home fa-fw"></i> Home</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link active" to="/info"> <i class="fas fa-cog fa-fw"></i> Info</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link active" to="/monitor"> <i class="fas fa-tachometer-alt fa-fw"></i> Monitor</router-link>
            </li>
            <li v-if="weatherEnabled" class="nav-item">
              <router-link class="nav-link active" to="/weather"> <i class="fas fa-umbrella fa-fw"></i> Weather</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link active" to="/about"> <i class="fas fa-info-circle fa-fw"></i> About</router-link>
            </li>
          </ul>
          <div v-if="authEnabled" class="d-flex">
            <button v-if="!user" class="btn btn-info btn-lg" variant="light" @click="doLogin()">
              <i class="fas fa-sign-in-alt"></i>&nbsp; Login
            </button>
            <router-link v-if="user" class="btn btn-info" to="/user" variant="light"> <i class="fas fa-user"></i>&nbsp; User </router-link>
          </div>
        </div>
      </div>
    </nav>

    <div class="container mt-5">
      <router-view v-slot="{ Component }">
        <transition name="route" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </div>
  </div>
</template>

<script>
import auth from './services/auth'
import router from './router'
import { config } from './main'

export default {
  data: function () {
    return {
      user: null,
      // We switch the user sign-in feature on or off depending if auth has been configured
      authEnabled: auth.isConfigured(),
      weatherEnabled: config.weatherEnabled
    }
  },

  async created() {
    // Restore any cached or saved local user
    if (auth.isConfigured()) {
      this.user = auth.user()
    }
  },

  methods: {
    async doLogin() {
      try {
        await auth.login()
        this.user = auth.user()
        router.push('user')
      } catch (err) {
        this.error = err.toString()
      }
    },

    clearUser() {
      this.user = null
    }
  }
}
</script>

<style>
html {
  min-height: 100%;
}

.route-enter-from {
  opacity: 0;
  transform: translateX(100px);
}
.route-enter-active {
  transition: all 0.3s ease-out;
}
.route-leave-to {
  opacity: 0;
  transform: translateX(-100px);
}
.route-leave-active {
  transition: all 0.3s ease-in;
}
</style>
