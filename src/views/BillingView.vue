<template>
  <BillingPage1 />
</template>

<script>
import { defineComponent } from "vue";

// Components
import BillingPage1 from "../components/BillingPage1.vue";

export default defineComponent({
  name: "BillingView",

  components: {
    BillingPage1,
  },
});
</script>
