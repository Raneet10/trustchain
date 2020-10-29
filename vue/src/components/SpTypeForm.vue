<template>
  <div>
    <div class="container">
      <sp-h3>New {{ type }}</sp-h3>
      <div v-for="field in fields" :key="field">
        <div  v-if="field == 'deadline'">
            <label>Deadline</label>
            <date-picker mode="dateTime" v-model="fieldsList['deadline']" :disabled="flight" />
        </div>
        <sp-input
          v-else
          v-model="fieldsList[field]"
          type="text"
          :placeholder="title(field)"
          :disabled="flight"
        />
      </div>
      <div class="button">
        <sp-button
          :loading="flight"
          :disabled="!valid || !hasAddress || flight"
          @click="submit"
        >
          Create {{ type }}
        </sp-button>
      </div>
      <sp-type-list :type="type" :module="module"/>
    </div>
  </div>
</template>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;531;600;700;800&display=swap");

.container {
  font-family: "Inter";
}
.button {
  display: inline-block;
}
</style>

<script>
import {SpInput} from "@tendermint/vue";
import {IconLoading2} from "@tendermint/vue";
import {SpH3} from "@tendermint/vue";
import {SpButton} from "@tendermint/vue";
import {DatePicker} from "v-calendar";
import {default as SpTypeList} from "../components/SpTypeList.vue";

export default {
  props: {
    type: {
      type: String,
    },
    fields: {
      default: () => [],
    },
    preflight: {
      default: () => {
        return (obj) => obj;
      },
    },
    module: {
      type: String,
    },
  },
  components: {
    SpInput,
    IconLoading2,
    SpH3,
    SpButton,
    SpTypeList,
    DatePicker,
  },
  data: function() {
    return {
      fieldsList: {},
      flight: false,
    };
  },
  created() {
    (this.fields || []).forEach((field) => {
      this.$set(this.fieldsList, field, "");
    });
  },
  computed: {
    hasAddress() {
      return !!this.$store.state.cosmos.account.address;
    },
    valid() {
      return Object.values(this.fieldsList).every((el) => {
        return !isNaN(Date.parse(el)) || el.trim().length > 0;
      });
    },
  },
  methods: {
    title(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    },
    async submit() {
      if (this.valid && !this.flight && this.hasAddress) {
        this.flight = true;

        this.fieldsList.reward = this.fieldsList.reward + 'trust';
        this.fieldsList.deadline = Date.parse(this.fieldsList.deadline).toString();

        const payload = {
          type: this.type,
          body: this.preflight(this.fieldsList),
          module: this.module,
        };
        try {
            await this.$store.dispatch("cosmos/entitySubmit", payload);
            await this.$store.dispatch("cosmos/entityFetch", {
              type: this.type,
              body: this.fieldsList,
              module: this.module,
            });
            this.flight = false;
        } catch (error) {
            console.error(error);
        }
        Object.keys(this.fieldsList).forEach((f) => {
          this.fieldsList[f] = "";
        });
      }
    },
  },
};
</script>

