<!--
User page
-----------------
Ben C, April 2021
-->

<template>
  <div class="card">
    <div class="card-header bg-success h3"><fa icon="user" />&nbsp; User Account</div>
    <div v-if="user" class="card-body">
      <b-button variant="warning" class="float-right" size="lg" @click="doLogout()">LOGOUT</b-button>

      <table>
        <tr>
          <td style="width: 40%">Name</td>
          <td>{{ user.name }}</td>
        </tr>
        <tr>
          <td style="width: 40%">Username</td>
          <td>{{ user.username }}</td>
        </tr>
        <tr v-if="user && user.idTokenClaims">
          <td style="width: 40%">Email</td>
          <td>{{ user.idTokenClaims.email }}</td>
        </tr>
        <tr v-if="user && user.idTokenClaims">
          <td style="width: 40%">Phone</td>
          <td>{{ graphDetails.mobilePhone }}</td>
        </tr>
        <tr v-if="user && user.idTokenClaims">
          <td style="width: 40%">Job Title</td>
          <td>{{ graphDetails.jobTitle }}</td>
        </tr>
        <tr v-if="user && user.idTokenClaims">
          <td style="width: 40%">Department</td>
          <td>{{ graphDetails.department }}</td>
        </tr>
        <tr v-if="user && user.idTokenClaims">
          <td style="width: 40%">Photo</td>
          <td><img class="graphphoto" :src="graphPhoto" alt="user" /></td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import auth from '../services/auth'
import graph from '../services/graph'
import router from '../router'

export default {
  data: function() {
    return {
      user: {},
      graphDetails: {},
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
.graphphoto {
  border-radius: 50%;
}
tr {
  height: 2.5rem;
}
</style>
