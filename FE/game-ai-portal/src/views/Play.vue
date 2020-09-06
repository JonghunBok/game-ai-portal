<template>
  <div class="play">
    <Nav />
    <div class="deck">
      <BotCard botName="PurpleWave" />
      <BotCard botName="WillBot" />
      <BotCard botName="Locutus" />
      <BotCard botName="SteamHammer" />
    </div>
  </div>
</template>

<script>
import BotCard from '@/components/BotCard.vue'
import Nav from '@/components/Nav.vue'

export default {
  name: 'PlayGround',
  components: {
    BotCard,
    Nav
  }
}
</script>

<style>
.play .deck {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
</style>
