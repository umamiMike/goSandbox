<template>
  <div id="app">
    <img src="https://jwt.io/img/pic_logo.svg">
    <h3>This is a demo of how the process of using <a href="http://jwt.io">Javascript Web Tokens</a> in a single sign on setup</h3>
    
    <div class="container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;

}
.container{
  width: 100%;
  max-width: 1200px;
  padding: 0 15px;
  margin: auto;
  overflow: hidden;
}
</style>
