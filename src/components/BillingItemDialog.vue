<template>
  <v-dialog
    :model-value="modelValue"
    @update:model-value="$emit('update:modelValue', $event)"
    max-width="600px"
  >
    <v-card rounded="lg" elevation="4">
      <v-card-title class="bg-deep-purple-darken-1 text-white">
        <v-icon icon="mdi-cart-plus" class="mr-2"></v-icon>
        Add Item
      </v-card-title>
      <v-card-text class="pt-6">
        <v-form ref="form" v-model="valid" @submit.prevent="addItem">
          <v-autocomplete
            v-model="selectedMedicineName"
            :items="medicineNames"
            label="Search Medicine"
            variant="outlined"
            density="comfortable"
            color="deep-purple"
            prepend-inner-icon="mdi-pill"
            :rules="[rules.required]"
            @update:model-value="onMedicineSelect"
          ></v-autocomplete>

          <v-row v-if="selectedMedicineName">
            <v-col cols="12" sm="6">
              <v-select
                v-model="selectedBatchId"
                :items="availableBatches"
                item-title="batchId"
                item-value="batchId"
                label="Select Batch"
                variant="outlined"
                density="comfortable"
                color="deep-purple"
                :rules="[rules.required]"
                @update:model-value="onBatchSelect"
                no-data-text="No stock available"
              >
                <template v-slot:item="{ props, item }">
                  <v-list-item
                    v-bind="props"
                    :subtitle="`Exp: ${item.raw.expiryDate} | MRP: ₹${item.raw.mrp}`"
                  ></v-list-item>
                </template>
              </v-select>
            </v-col>
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="item.expiryDate"
                label="Expiry"
                readonly
                variant="outlined"
                density="comfortable"
                color="grey"
              ></v-text-field>
            </v-col>
          </v-row>

          <v-row>
            <v-col cols="6">
              <v-text-field
                v-model.number="item.mrp"
                label="MRP (₹)"
                type="number"
                readonly
                variant="outlined"
                density="comfortable"
                color="grey"
                prefix="₹"
              ></v-text-field>
            </v-col>
            <v-col cols="6">
              <v-text-field
                v-model.number="item.gst"
                label="GST (%)"
                type="number"
                readonly
                variant="outlined"
                density="comfortable"
                color="grey"
                suffix="%"
              ></v-text-field>
            </v-col>
          </v-row>

          <v-text-field
            v-model.number="item.quantity"
            label="Quantity"
            type="number"
            variant="outlined"
            density="comfortable"
            color="deep-purple"
            :rules="[rules.required, rules.positive, rules.maxStock]"
            ref="qtyField"
          >
            <template v-slot:append-inner>
              <span v-if="maxQty" class="text-caption text-grey"
                >Max: {{ maxQty }}</span
              >
            </template>
          </v-text-field>

          <div class="d-flex justify-end mt-4">
            <v-btn
              color="grey"
              variant="text"
              class="mr-2"
              @click="$emit('update:modelValue', false)"
            >
              Cancel
            </v-btn>
            <v-btn
              type="submit"
              color="deep-purple-darken-1"
              variant="elevated"
            >
              Add Item
            </v-btn>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "BillingItemDialog",
  props: {
    modelValue: Boolean,
    medicines: {
      type: Array,
      default: () => [],
    },
    currentItems: {
      type: Array,
      default: () => [],
    },
  },
  emits: ["update:modelValue", "add-item"],
  data() {
    return {
      valid: false,
      selectedMedicineName: null,
      selectedBatchId: null,
      item: {
        medicineName: "",
        batchId: "",
        expiryDate: "",
        quantity: 1,
        mrp: 0,
        gst: 0,
      },
      maxQty: null,
      rules: {
        required: (v) => !!v || "Field is required",
        positive: (v) => v > 0 || "Must be > 0",
        maxStock: (v) =>
          !this.maxQty ||
          v <= this.maxQty ||
          `Exceeds available stock (${this.maxQty})`,
      },
    };
  },
  computed: {
    medicineNames() {
      return this.medicines.map((m) => m.name);
    },
    availableBatches() {
      if (!this.selectedMedicineName) return [];
      const med = this.medicines.find(
        (m) => m.name === this.selectedMedicineName
      );
      return med ? med.batches : [];
    },
  },
  watch: {
    modelValue(val) {
      if (!val) {
        this.resetForm();
      }
    },
  },
  methods: {
    onMedicineSelect(name) {
      this.selectedBatchId = null;
      this.item = {
        ...this.item,
        medicineName: name,
        batchId: "",
        expiryDate: "",
        mrp: 0,
        gst: 0,
        quantity: 1,
      };
      this.maxQty = null;
    },
    onBatchSelect(batchId) {
      const batch = this.availableBatches.find((b) => b.batchId === batchId);
      if (batch) {
        this.item.batchId = batchId;
        this.item.expiryDate = batch.expiryDate;
        this.item.mrp = batch.mrp;
        this.item.gst = batch.gst;

        // Calculate already drafted quantity for this batch
        const draftedQty = this.currentItems
          .filter((i) => i.batchId === batchId)
          .reduce((acc, i) => acc + (parseInt(i.quantity) || 0), 0);

        this.maxQty = batch.qty - draftedQty;
        if (this.maxQty < 0) this.maxQty = 0;

        // If current qty > max, reset it
        if (this.item.quantity > this.maxQty) {
          this.item.quantity = this.maxQty > 0 ? 1 : 0;
        }
      }
    },
    async addItem() {
      const { valid } = await this.$refs.form.validate();
      if (valid) {
        this.$emit("add-item", { ...this.item });
        this.$emit("update:modelValue", false);
      }
    },
    resetForm() {
      this.selectedMedicineName = null;
      this.selectedBatchId = null;
      this.item = {
        medicineName: "",
        batchId: "",
        expiryDate: "",
        quantity: 1,
        mrp: 0,
        gst: 0,
      };
      this.$refs.form?.resetValidation();
      this.maxQty = null;
    },
  },
};
</script>
