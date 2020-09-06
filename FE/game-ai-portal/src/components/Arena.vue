<template>
  <div class="arena">

    <div id="view"></div>
  </div>
</template>

<script>
//<VueNoVNC websockify="http://park-terminal:6080/vnc.html?host=park-terminal&port=6080" />
import RFB from '@novnc/novnc/core/rfb';
//import VueNoVNC from '@phasedlogix/vue-novnc/dist/vue-novnc';

export default {
  name: 'Arena',
  props: {
    msg: String,
  },
  data() {
    return {
      websockify: "ws://scbw.jonghun.onl:6080"
    }
  },

  methods: {
    connect() {
      //this.rfb.sendCredentials({ password: this.passwd });
      console.log(this.rfb)
      console.log(this.$el.lastChild)
      
  //    document.querySelector('canvas').requestFullScreen();
    }
  },

  mounted() {
    this.$nextTick(() => {
      console.log("Ticekd!")
      this.rfb = new RFB(this.$el.lastChild, this.websockify);

      fetch("http://scbw.jonghun.onl:9090")
        .then(function(response) {
          console.log(response);
          return response.json();
        })

      this.connect();
      

    })
  }
}
</script>

<style scoped lang="less">
* {
  height: 100%;
}
</style>
