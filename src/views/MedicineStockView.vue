<template>
  <v-container fluid class="bg-grey-lighten-4 fill-height align-start">
    <v-main class="w-100">
      <v-container class="pa-6">
        <h1 class="text-h4 mb-4 text-indigo-darken-3 font-weight-bold">
          Medicine Stock Management
        </h1>
        <v-card elevation="2" rounded="lg" class="border-t-4 border-indigo">
          <v-card-title class="text-indigo-darken-2">
            <v-icon icon="mdi-pill" class="mr-2"></v-icon>
            Add New Stock
          </v-card-title>
          <v-card-text class="pt-4">
            <v-form ref="form" v-model="valid" @submit.prevent="addStock">
              <v-row>
                <v-col cols="12" md="6">
                  <!-- Name Combobox -->
                  <v-combobox
                    v-model="stock.name"
                    :items="existingMedicineNames"
                    label="Medicine Name"
                    variant="outlined"
                    density="comfortable"
                    color="indigo"
                    :rules="[rules.required]"
                    prepend-inner-icon="mdi-format-title"
                    @update:model-value="onNameSelect"
                  ></v-combobox>
                </v-col>
                <v-col cols="12" md="6">
                  <!-- Batch Combobox -->
                  <v-combobox
                    v-model="stock.batchId"
                    :items="existingBatches"
                    label="Batch ID"
                    variant="outlined"
                    density="comfortable"
                    color="indigo"
                    :rules="[rules.required]"
                    prepend-inner-icon="mdi-barcode"
                    @update:model-value="onBatchSelect"
                  ></v-combobox>
                </v-col>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="stock.expiryDate"
                    label="Expiry Date"
                    type="date"
                    variant="outlined"
                    density="comfortable"
                    color="indigo"
                    :rules="[rules.required]"
                    prepend-inner-icon="mdi-calendar-alert"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model.number="stock.qty"
                    label="Quantity"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    color="indigo"
                    :rules="[rules.required, rules.positive]"
                    prepend-inner-icon="mdi-package-variant-closed"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model.number="stock.mrp"
                    label="MRP (₹)"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    color="indigo"
                    :rules="[rules.required, rules.positive]"
                    prefix="₹"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model.number="stock.gst"
                    label="GST (%)"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    color="indigo"
                    :rules="[rules.required, rules.nonNegative]"
                    suffix="%"
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-divider class="my-4"></v-divider>
              <div class="d-flex justify-end">
                <v-btn
                  type="submit"
                  color="indigo-darken-3"
                  size="large"
                  prepend-icon="mdi-plus-circle"
                  :loading="loading"
                >
                  Add to Stock
                </v-btn>
              </div>
            </v-form>
          </v-card-text>
        </v-card>

        <v-snackbar
          v-model="snackbar.show"
          :color="snackbar.color"
          location="top right"
        >
          {{ snackbar.text }}
          <template v-slot:actions>
            <v-btn variant="text" @click="snackbar.show = false">Close</v-btn>
          </template>
        </v-snackbar>
      </v-container>
    </v-main>
  </v-container>
</template>

<script>
import eventServices from "@/services/eventServices";

export default {
  name: "MedicineStockView",
  data() {
    return {
      valid: false,
      loading: false,
      stock: {
        name: "",
        batchId: "",
        expiryDate: "",
        qty: null,
        mrp: null,
        gst: null,
      },
      rules: {
        required: (v) => !!v || "Field is required",
        positive: (v) => v > 0 || "Must be greater than 0",
        nonNegative: (v) => v >= 0 || "Must be non-negative",
      },
      snackbar: {
        show: false,
        text: "",
        color: "success",
      },
      allMedicines: [],
      existingBatches: [], // For combobox items
    };
  },
  computed: {
    existingMedicineNames() {
      // Unique names
      const names = this.allMedicines.map((m) => m.name);
      return [...new Set(names)];
    },
  },
  mounted() {
    this.fetchMedicines();
  },
  methods: {
    async fetchMedicines() {
      try {
        const response = await eventServices.getMedicines();
        this.allMedicines = response.data || [];
      } catch (error) {
        console.error("Failed to fetch medicines:", error);
        this.showSnackbar("Failed to fetch medicines", "error");
      }
    },
    onNameSelect(name) {
      if (!name) {
        this.existingBatches = [];
        return;
      }
      // Find all batches for this medicine name
      const medicines = this.allMedicines.filter(
        (m) => m.name.toLowerCase() === (name.title || name).toLowerCase()
      );

      if (medicines.length > 0) {
        this.existingBatches = medicines.map((m) => m.batchId);
      } else {
        this.existingBatches = [];
      }
    },
    onBatchSelect(batchId) {
      const name =
        typeof this.stock.name === "object"
          ? this.stock.name.title
          : this.stock.name;
      if (!name) return;

      const medicine = this.allMedicines.find(
        (m) =>
          m.name.toLowerCase() === name.toLowerCase() && m.batchId === batchId
      );

      if (medicine) {
        // Auto-fill existing batch details
        this.stock.expiryDate = medicine.expiryDate;
        this.stock.mrp = medicine.mrp;
        this.stock.gst = medicine.gst;
        // Qty should be empty (user adds new qty)
      } else {
        // New batch, maybe clear details if needed?
        this.stock.expiryDate = "";
        this.stock.mrp = null;
        this.stock.gst = null;
        this.stock.qty = null;
      }
    },
    async addStock() {
      // Combobox returns object if item is selected object, or string if typed
      const name =
        typeof this.stock.name === "object"
          ? this.stock.name.title || this.stock.name
          : this.stock.name;
      const batchId =
        typeof this.stock.batchId === "object"
          ? this.stock.batchId.title || this.stock.batchId
          : this.stock.batchId;

      const payload = { ...this.stock, name, batchId };

      const { valid } = await this.$refs.form.validate();
      if (!valid) return;

      this.loading = true;
      try {
        const response = await eventServices.addMedicine(payload);
        if (response.data.status) {
          this.showSnackbar("Stock added successfully", "success");
          this.$refs.form.reset();
          this.stock = {
            name: "",
            batchId: "",
            // ... reset other fields naturally by form reset
            qty: null,
            mrp: null,
            gst: null,
          };
          this.existingBatches = [];
          await this.fetchMedicines(); // Refresh list
        } else {
          throw new Error(response.data.message);
        }
      } catch (error) {
        this.showSnackbar(error.message || "Failed to add stock", "error");
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
    showSnackbar(text, color) {
      this.snackbar.text = text;
      this.snackbar.color = color;
      this.snackbar.show = true;
    },
  },
};
</script>
