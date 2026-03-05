<template>
  <v-container fluid class="pa-6">
    <v-row>
      <v-col cols="12" sm="6" md="4">
        <v-card
          class="mx-auto"
          prepend-icon="mdi-receipt"
          title="New Bill"
          subtitle="Create a new invoice"
          link
          @click="navigateToBilling"
        >
          <template v-slot:append>
            <v-icon icon="mdi-arrow-right" color="primary"></v-icon>
          </template>
          <v-card-text>
            Start creating a new bill for a patient or customer.
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <v-card
          class="mx-auto"
          prepend-icon="mdi-account-multiple"
          title="History"
          subtitle="View past bills"
          link
          disabled
        >
          <v-card-text>testing card</v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <div class="d-flex align-center justify-center mt-12">
      <div class="text-center opacity-60">
        <div class="d-flex justify-center mb-4">
          <TestLogo />
        </div>
        <div class="text-h6">
          Welcome, {{ currentUser ? currentUser.name : "Guest" }}
        </div>
        <div
          v-if="currentUser"
          class="text-subtitle-2 text-uppercase text-primary mb-2"
        >
          Role: {{ currentUser.role }}
        </div>
        <div>Select an option from the menu create a new bill.</div>
      </div>
    </div>
  </v-container>
</template>

<script>
import TestLogo from "@/components/TestLogo.vue";

export default {
  name: "DashboardPage",
  components: {
    TestLogo,
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    },
    // isAdminOrManager logic removed as it was only for sidebar links
  },
  methods: {
    navigateToBilling() {
      this.$router.push({ name: "billing" });
    },
    // navigateToLogs and navigateToUsers removed as they were for sidebar links, unless cards use them?
    // The code only had cards for 'New Bill' and 'History' (disabled).
    // So distinct methods for Logs/Users for cards aren't needed here unless I see them in the template...
    // Template had New Bill (billing) and History. No Logs/Users cards.
  },
};
</script>
