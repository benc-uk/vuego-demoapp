<template>
  <div id="app">

    <b-navbar toggleable="md" type="dark" variant="primary">
      <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
      <b-navbar-brand to="/"><img src="./assets/vuejs.svg" height="50px"></b-navbar-brand>
      <b-collapse is-nav id="nav_collapse">
        <b-navbar-nav>
          <b-nav-item to="/app/home"><i class="fas fa-home"></i>&nbsp; Home</b-nav-item>
          <b-nav-item to="/app/info"><i class="fas fa-cogs"></i>&nbsp; Sys Info</b-nav-item>
          <b-nav-item to="/app/about"><i class="far fa-info-circle"></i>&nbsp; About</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <br/>
    <b-container>
      <router-view/>
    </b-container>

  </div>
</template>

<style lang="scss">
  html {
    font-size: 20px;
  } 

  #app, html, body {
    height: 100%;
    background-color: #317256;
  }
</style>
