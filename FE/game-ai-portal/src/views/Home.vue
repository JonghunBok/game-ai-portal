<template>
  <div class="home">
    <Nav />
    <div class="welcoming-msg">
      <p>Can you beat it?</p>
    </div>
  </div>
</template>

<script>
// <HelloWorld msg="Welcome to Your Vue.js App"/>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Nav from '@/components/Nav.vue'

export default {
  name: 'Home',
  components: {
    Nav
  }
}
</script>
