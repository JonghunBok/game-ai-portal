<template>
  <div class="play-ground">

    <div class="arena-box">
      <Arena :msg="$route.params.botName" />
    </div>
    <div class="panel">
      <h1> {{ $route.params.botName }} </h1>
      <hr/>
      <p> 갖가지 설명 </p>
      <p> 이 봇의 전적 </p>
      <p> esc: 전체화면 나가기 </p>
      <button @click="fullScreen"> 전체화면 </button>
    </div>
    

  </div>
</template>

<script>
import Arena from '@/components/Arena.vue'

export default {
  name: 'Play',
  props: {
  },
  components: {
    Arena
  },
  methods: {
    fullScreen() {
      console.log("full screen mode on!")
      const canvas = document.querySelector("canvas")
      canvas.requestFullscreen();
    }
  }
}
</script>

<style>

.play-ground {
  display: flex;
  height: 100%;
}

.arena-box {
  display: block;
  height: 100%;
  flex-grow: 1;

}

.panel {
  display: block;
  width: 200px;
  height: 100%;
  box-shadow: 10px 0px 13px 9px grey;
}

</style>
