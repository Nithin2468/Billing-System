<template>
  <v-container fluid class="pa-6">
    <v-card elevation="2" rounded="lg">
      <v-card-title class="d-flex align-center bg-primary text-white py-4 px-6">
        <v-icon icon="mdi-history" class="mr-3"></v-icon>
        <span class="text-h5 font-weight-bold">Activity Logs</span>
        <v-spacer></v-spacer>
        <v-chip color="white" variant="outlined" class="font-weight-bold">
          Total Logs: {{ logs.length }}
        </v-chip>
      </v-card-title>

      <v-card-text class="pa-0">
        <v-data-table
          :headers="headers"
          :items="logs"
          :items-per-page="10"
          class="elevation-0"
          hover
        >
          <template v-slot:[`item.event`]="{ item }">
            <v-chip
              :color="getEventColor(item.event)"
              size="small"
              class="text-uppercase font-weight-bold"
              label
            >
              {{ item.event }}
            </v-chip>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: "LogPage",
  data() {
    return {
      headers: [
        {
          title: "Timestamp",
          key: "timestamp",
          align: "start",
          sortable: true,
        },
        { title: "User ID", key: "userid", align: "start" },
        { title: "Name", key: "name", align: "start" },
        { title: "Event", key: "event", align: "center" },
      ],
    };
  },
  computed: {
    logs() {
      return this.$store.getters.allLogs;
    },
  },
  methods: {
    getEventColor(event) {
      switch (event) {
        case "Login":
          return "success";
        case "Logout":
          return "warning";
        default:
          return "grey";
      }
    },
  },
};
</script>
