<!--
Core app component, just template no code/logic
----------------------------------------------
Ben C, April 2018
-->

<template>
  <div id="app">
    <b-navbar toggleable="md" type="dark" variant="primary">
      <b-navbar-toggle target="nav_collapse" />
      <b-navbar-brand to="/">
        <img src="./assets/vuejs.svg" height="50px">Vue.js &amp; Go
      </b-navbar-brand>
      <b-collapse id="nav_collapse" is-nav>
        <b-navbar-nav>
          <b-button to="/home" variant="dark">
            <fa icon="home" />&nbsp; Home
          </b-button>
          <b-button to="/info" variant="dark">
            <fa icon="cogs" />&nbsp; Info
          </b-button>
          <b-button to="/monitor" variant="dark">
            <fa icon="tachometer-alt" />&nbsp; Monitor
          </b-button>
          <b-button to="/weather" variant="dark">
            <fa icon="umbrella" />&nbsp; Weather
          </b-button>
          <b-button to="/about" variant="dark">
            <fa icon="info-circle" />&nbsp; About
          </b-button>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <br>

    <b-container>
      <transition name="slidedown">
        <router-view />
      </transition>
    </b-container>
  </div>
</template>

<style>
  html {
    font-size: 20px;
  }
  #app, html, body {
    height: 100%;
    background-color: #317256;
  }
  .b-navbar-nav {
    color: red;
  }
  .slidedown-enter-active {
    transition: transform 0.4s ease-out;
    transform-origin: top;
  }
  .slidedown-enter {
    transform: scaleY(0);
  }
</style>

<style scoped>
  .btn {
    margin-right: 0.3rem;
    width: 7rem;
  }
</style>
