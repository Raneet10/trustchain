<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->

      <promises-list type="promise" module="trustchain" :keeper="this.$store.state.cosmos.account.address" v-if="this.$store.state.cosmos.account.address != null" />

	<sp-type-form type="promise" :fields="['promiseDescription', 'promiseKeeper', 'reward', 'deadline']" module="trustchain" />
    </div>
  </div>
</template>

<script>
import {default as SpTypeForm} from "../components/SpTypeForm.vue";
import {default as PromisesList} from "../components/PromisesList.vue";
import {SpSignIn} from "@tendermint/vue";
import {SpBankBalances} from "@tendermint/vue";
import {SpTokenSend} from "@tendermint/vue";
import {SpH3} from "@tendermint/vue";

export default {
  components: { SpTypeForm, SpSignIn, SpBankBalances, SpTokenSend, SpH3, PromisesList },
};
</script>
