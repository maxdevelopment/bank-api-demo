<template>
  <div>
    <input v-model="account">
    <input v-model="deposit">
    <button v-on:click="updateWithdraw()">withdraw update</button>
  </div>
</template>

<script>
  import {EventBus} from "./EventBus.js";

  export default {
    name: "Withdraw",
    data() {
      return {
        deposit: 0,
        account: ""
      }
    },
    methods: {
      updateWithdraw: function()  {
        EventBus.$emit('withdraw-update', {
          sum: this.deposit,
          id: this.account
        })
      }
    }
  }
</script>

<style scoped>

</style>
