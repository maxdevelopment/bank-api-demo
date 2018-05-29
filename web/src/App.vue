<template>
  <div id="app">
    <button v-on:click="openAccount">open account</button>
    <ul>
      <li v-for="account in accounts">
        {{ account }}<a v-on:click="closeAccount(account.id)">del</a>
      </li>
    </ul>
    <deposit></deposit>
    <withdraw></withdraw>
    <transfer></transfer>
  </div>
</template>

<script>
  import Vue from 'vue';
  import VueResource from 'vue-resource';
  import Deposit from './Deposit';
  import Withdraw from './Withdraw';
  import {EventBus} from "./EventBus.js";
  import Transfer from "./Transfer";

  Vue.use(VueResource);

  const openAccount = Vue.resource('/api/open-account');
  const closeAccount = Vue.resource('/api/close-account{/id}');
  const getAccounts = Vue.resource('/api/accounts');

  export default {
    name: 'app',
    components: {Transfer, Withdraw, Deposit},
    comments: {
      Deposit,
      Withdraw
    },
    data() {
      return {
        accounts: []
      }
    },
    created() {
      this.fetchAccounts()
    },
    mounted() {
      EventBus.$on('deposit-update', (value) => {
        console.log(value)
      });
      EventBus.$on('withdraw-update', (value) => {
        console.log(value)
      });
      EventBus.$on('transfer', (value) => {
        console.log(value)
      });
    },
    methods: {
      closeAccount: function (id) {
        closeAccount.delete({id: id}).then(response => {
          this.fetchAccounts()
        }, err => {})
      },
      fetchAccounts: function () {
        this.accounts = [];
        return getAccounts.get().then(response => {
          for (let val of Object.values(response.body)) {
            this.accounts.push(val)
          }
        }, err => {
          console.log(err)
        })
      },
      openAccount: function () {
        openAccount.get().then(response => {
          this.accounts.push(response.body)
        }, err => {
        })
      }
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }

  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
