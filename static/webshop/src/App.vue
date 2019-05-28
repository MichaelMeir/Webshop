<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
@tailwind preflight;
@tailwind components;
@tailwind utilities;

@import url('https://fonts.googleapis.com/css?family=Roboto&display=swap');

html {
  font-family: 'Roboto', sans-serif;
  padding: 0;
  margin: 0;
}
</style>
