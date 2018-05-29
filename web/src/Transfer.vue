<template>
  <div>
    <input v-model="accountFrom">
    <input v-model="accountTo">
    <input v-model="sum">
    <button v-on:click="transfer()">transfer</button>
  </div>
</template>

<script>
  import {EventBus} from "./EventBus.js";

  export default {
    name: "Transfer",
    data() {
      return {
        sum: 0,
        accountFrom: "",
        accountTo: ""
      }
    },
    methods: {
      transfer: function()  {
        EventBus.$emit('transfer', {
          sum: this.sum,
          idFrom: this.accountFrom,
          idTo: this.accountTo
        })
      }
    }
  }
</script>

<style scoped>

</style>
