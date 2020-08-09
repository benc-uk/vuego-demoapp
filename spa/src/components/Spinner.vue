<!--
Loading spinner, just HTML & CSS3
----------------------------------------------
Ben C, April 2018
-->

<template>
  <div class="spinner">
    <div class="bounce1" />
    <div class="bounce2" />
    <div class="bounce3" />
  </div>
</template>

<script>
export default {
  name: 'Spinner'
}
</script>

<style>
.spinner {
  margin: 10px auto 0;
  width: 90px;
  text-align: center;
}

.spinner > div {
  width: 28px;
  height: 28px;
  background-color: #1f573e;

  border-radius: 0%;
  margin: 1px;
  display: inline-block;
  -webkit-animation: sk-bouncedelay 1.4s infinite ease-in-out both;
  animation: sk-bouncedelay 1.4s infinite ease-in-out both;
}

.spinner .bounce1 {
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}

.spinner .bounce2 {
  -webkit-animation-delay: -0.24s;
  animation-delay: -0.24s;
}

.spinner .bounce3 {
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}

@-webkit-keyframes sk-bouncedelay {
  0%, 80%, 100% { -webkit-transform: scale(0) }
  40% { -webkit-transform: scale(1.0) }
}

@keyframes sk-bouncedelay {
  0%, 80%, 100% {
    -webkit-transform: scale(0);
    transform: scale(0);
  } 40% {
    -webkit-transform: scale(1.0);
    transform: scale(1.0);
  }
}
</style>