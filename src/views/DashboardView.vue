<template>
  <v-layout class="rounded rounded-md">
    <v-app-bar color="primary" density="compact">
      <template v-slot:prepend>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      </template>

      <v-app-bar-title>Billing App</v-app-bar-title>

      <div
        v-if="currentUser"
        class="text-subtitle-2 mr-4 text-white d-none d-sm-flex align-center"
      >
        <v-icon icon="mdi-account" size="small" class="mr-2"></v-icon>
        <span>{{ currentUser.name }}</span>
        <v-chip
          size="x-small"
          class="ml-2 text-uppercase"
          color="secondary"
          label
        >
          {{ currentUser.role }}
        </v-chip>
      </div>

      <template v-slot:append>
        <v-btn
          icon="mdi-logout"
          variant="text"
          color="white"
          @click="logout"
          title="Logout"
        ></v-btn>
      </template>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" :permanent="drawer">
      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-view-dashboard"
          title="Dashboard"
          :to="{ name: 'dashboard-home' }"
          color="primary"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-receipt-text"
          title="Billing"
          :to="{ name: 'billing' }"
          color="primary"
          v-if="!isAdmin"
        >
        </v-list-item>
        <v-list-item
          v-if="isManager"
          prepend-icon="mdi-pill"
          title="Medicine Stocks"
          :to="{ name: 'stocks' }"
          color="primary"
        ></v-list-item>

        <v-list-item
          prepend-icon="mdi-clipboard-list-outline"
          title="Available Stocks"
          :to="{ name: 'available-medicines' }"
          color="primary"
        ></v-list-item>

        <v-list-item
          v-if="isAdminOrManager"
          prepend-icon="mdi-chart-bar"
          title="Stock & Sales Report"
          :to="{ name: 'stock-report' }"
          color="primary"
        ></v-list-item>

        <!-- Admin/Manager Only Links -->
        <template v-if="isAdminOrManager">
          <v-divider class="my-2"></v-divider>
          <v-list-item
            prepend-icon="mdi-history"
            title="Billing History"
            :to="{ name: 'history' }"
            color="primary"
          ></v-list-item>
          <v-list-item
            prepend-icon="mdi-textBox-multiple"
            title="Activity Logs"
            :to="{ name: 'logs' }"
            color="primary"
          ></v-list-item>
          <v-list-item
            prepend-icon="mdi-account-group"
            title="User Management"
            :to="{ name: 'users' }"
            color="primary"
          ></v-list-item>
        </template>
      </v-list>
    </v-navigation-drawer>

    <v-main class="bg-grey-lighten-3 fill-height">
      <!-- Nested Views Render Here -->
      <router-view></router-view>
    </v-main>
  </v-layout>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "DashboardView",
  data() {
    return {
      drawer: true,
    };
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    },
    isAdmin() {
      return this.$store.getters.userRole === "admin";
    },
    isManager() {
      return this.$store.getters.userRole === "manager";
    },
    isBiller() {
      return this.$store.getters.userRole === "biller";
    },
    isAdminOrManager() {
      return this.isAdmin || this.isManager;
    },
  },
  methods: {
    async logout() {
      await this.$store.dispatch("logout");
      this.$router.push({ name: "login" });
    },
  },
});
</script>
