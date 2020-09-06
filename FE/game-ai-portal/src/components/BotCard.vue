<template>
  <div class="bot-card">
    <h1>{{ botName }}</h1>
    <hr />
    <p> 개발자: ~~~</p> <br/>
    <p> 대략의 설명, 전적, 웹사이트</p>
    <div class="button">
      <router-link :to="{ name: 'PlayGround', params: { botName: botName }}">Challenge!</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BotCard',
  props: {
    botName: String
  }
}
</script>

<style scoped lang="less">
.bot-card {
  display: block;
  width: 200px;
  height: 300px;
  background-color: #2c3033;
  box-shadow: 2px 3px 2px #FFFCF3;
 
}
.button {
  width: 110px;
  height: 30px;
  background-color: #ffd770;
  border-radius: 15px;
  margin: auto;
  position: relative;
}

.button:hover {
  background-color: #c95311;
}

.button a {
  position: absolute;
  top: 50%;
  left: 20%;
  margin-top: -9px;
  color: #05090c;
  font-size: 14px;
}
</style>
