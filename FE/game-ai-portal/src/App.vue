<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style lang="less">
html, body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  background: #05090C;
  color: #FFFCF3;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  width: 100%;
  height: 100%;
  margin: 0;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #fffcf3;

    &.router-link-exact-active {
      color: #ffd770;
    }
  }

}

h1 {
  font-size: 18px;
}

p {
  font-size: 12px;
}

hr {
  border-color: #df3127;
}

a {
  text-decoration: none;
}

</style>
