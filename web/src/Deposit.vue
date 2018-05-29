<template>
  <div>
    <input v-model="account">
    <input v-model="deposit">
    <button v-on:click="updateDeposit()">deposit update</button>
  </div>
</template>

<script>
  import {EventBus} from "./EventBus.js";

  export default {
    name: "Deposit",
    data() {
      return {
        deposit: 0,
        account: ""
      }
    },
    methods: {
      updateDeposit: function()  {
        EventBus.$emit('deposit-update', {
          sum: this.deposit,
          id: this.account
        })
      }
    }
  }
</script>

<style scoped>

</style>
