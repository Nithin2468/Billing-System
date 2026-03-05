<template>
  <v-container fluid class="pa-6">
    <div class="d-flex justify-space-between align-center mb-6">
      <h2 class="text-h4 font-weight-bold text-grey-darken-3">
        <v-icon icon="mdi-pill" class="mr-2"></v-icon>
        Available Medicines
      </h2>
    </div>

    <v-card elevation="2" rounded="lg">
      <v-data-table
        :headers="headers"
        :items="flattenedMedicines"
        :search="search"
        hover
        class="stocks-table"
      >
        <template v-slot:top>
          <v-toolbar flat color="white" class="px-4">
            <v-text-field
              v-model="search"
              prepend-inner-icon="mdi-magnify"
              label="Search Medicines"
              single-line
              hide-details
              density="compact"
              variant="outlined"
              class="mr-4"
              style="max-width: 300px"
            ></v-text-field>
          </v-toolbar>
        </template>

        <template #[`item.qty`]="{ item }">
          <v-chip
            :color="getQtyColor(item.qty)"
            size="small"
            font-weight="bold"
            label
          >
            {{ item.qty }}
          </v-chip>
        </template>

        <template #[`item.mrp`]="{ item }">
          ₹{{ item.mrp.toFixed(2) }}
        </template>

        <template #[`item.actions`]="{ item }">
          <v-btn
            icon="mdi-delete"
            size="small"
            color="error"
            variant="text"
            @click="confirmRemove(item)"
          ></v-btn>
        </template>
      </v-data-table>
    </v-card>

    <!-- Remove Confirmation Dialog -->
    <v-dialog v-model="removeDialog" max-width="400">
      <v-card>
        <v-card-title class="bg-error text-white">
          <v-icon icon="mdi-alert" class="mr-2"></v-icon>
          Confirm Removal
        </v-card-title>
        <v-card-text class="pt-4">
          Are you sure you want to remove the stock for
          <strong>{{ itemToRemove?.name }}</strong> (Batch:
          {{ itemToRemove?.batchId }})? This action cannot be undone.
        </v-card-text>
        <v-card-actions class="justify-end px-4 pb-4">
          <v-btn color="grey" variant="text" @click="removeDialog = false"
            >Cancel</v-btn
          >
          <v-btn color="error" variant="elevated" @click="removeStock"
            >Remove</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
      {{ snackbarText }}
      <template v-slot:actions>
        <v-btn color="white" variant="text" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
export default {
  name: "AvailableMedicineView",
  data() {
    return {
      search: "",
      headers: [
        { title: "Medicine Name", key: "name", align: "start" },
        { title: "Batch ID", key: "batchId" },
        { title: "Expiry Date", key: "expiryDate" },
        { title: "Avail. Qty", key: "qty", align: "center" },
        { title: "MRP", key: "mrp", align: "end" },
        { title: "Actions", key: "actions", align: "center", sortable: false },
      ],
      removeDialog: false,
      itemToRemove: null,
      snackbar: false,
      snackbarText: "",
      snackbarColor: "success",
    };
  },
  computed: {
    allMedicines() {
      return this.$store.getters.allMedicines;
    },
    flattenedMedicines() {
      const flatList = [];
      this.allMedicines.forEach((med) => {
        if (med.batches && med.batches.length > 0) {
          med.batches.forEach((batch) => {
            flatList.push({
              name: med.name,
              batchId: batch.batchId,
              expiryDate: batch.expiryDate,
              qty: batch.qty,
              mrp: batch.mrp,
            });
          });
        } else {
          // Optional: Show medicines even if no batches (with 0 qty)
        }
      });
      return flatList;
    },
  },
  methods: {
    getQtyColor(qty) {
      if (qty <= 0) return "error";
      if (qty < 20) return "warning";
      return "success";
    },
    confirmRemove(item) {
      this.itemToRemove = item;
      this.removeDialog = true;
    },
    async removeStock() {
      if (!this.itemToRemove) return;

      try {
        await this.$store.dispatch("removeMedicineBatch", {
          name: this.itemToRemove.name,
          batchId: this.itemToRemove.batchId,
        });

        this.snackbarText = "Stock removed successfully";
        this.snackbarColor = "success";
        this.snackbar = true;
      } catch (error) {
        console.error("Error removing stock:", error);
        this.snackbarText = "Failed to remove stock";
        this.snackbarColor = "error";
        this.snackbar = true;
      } finally {
        this.removeDialog = false;
        this.itemToRemove = null;
      }
    },
  },
};
</script>

<style scoped>
.stocks-table :deep(th) {
  font-weight: bold !important;
  color: #424242;
}
</style>
