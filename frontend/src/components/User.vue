<!--
User page
-----------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div class="card border-light mb-3 rounded-3">
    <h4 class="card-header p-3 bg-info">
      <i class="fas fa-user fa-fw" /> User Account <button class="btn btn-secondary text-white float-end" @click="doLogout()">LOGOUT</button>
    </h4>
    <div v-if="user && graphDetails" class="card-body row">
      <ul class="list-group m-2 col">
        <li class="list-group-item d-flex justify-content-between align-items-center">
          Name <span class="badge bg-primary rounded-pill p-2 px-4 fs-6">{{ user.name }}</span>
        </li>
        <li class="list-group-item d-flex justify-content-between align-items-center">
          Username <span class="badge bg-primary rounded-pill p-2 px-4 fs-6"> {{ user.username }}</span>
        </li>
        <li class="list-group-item d-flex justify-content-between align-items-center">
          Email <span class="badge bg-primary rounded-pill p-2 px-4 fs-6"> {{ user.idTokenClaims.email }}</span>
        </li>
        <li class="list-group-item d-flex justify-content-between align-items-center">
          Phone <span class="badge bg-primary rounded-pill p-2 px-4 fs-6"> {{ graphDetails.mobilePhone }}</span>
        </li>
        <li class="list-group-item d-flex justify-content-between align-items-center">
          Job Totle <span class="badge bg-primary rounded-pill p-2 px-4 fs-6"> {{ graphDetails.jobTitle }}</span>
        </li>
        <li class="list-group-item d-flex justify-content-between align-items-center">
          Department <span class="badge bg-primary rounded-pill p-2 px-4 fs-6"> {{ graphDetails.department }}</span>
        </li>
      </ul>
      <div class="col photocol"><img class="graphphoto" :src="graphPhoto" /></div>
    </div>
  </div>
</template>

<script>
import auth from '../services/auth'
import graph from '../services/graph'
import router from '../router'

export default {
  data: function () {
    return {
      user: null,
      graphDetails: null,
      graphPhoto: null
    }
  },

  async mounted() {
    this.user = auth.user()
    this.graphDetails = await graph.getSelf()
    this.graphPhoto = await graph.getPhoto()
  },

  methods: {
    doLogout() {
      this.user = null
      this.graphDetails = null
      this.graphPhoto = null
      auth.clearLocal()
      this.$emit('logoutEvent')
      router.push('home')
    }
  }
}
</script>

<style scoped>
.photocol {
  display: grid;
}
.graphphoto {
  width: 260px;
  border-radius: 5%;
  margin: auto;
  align-self: center;
}
</style>
