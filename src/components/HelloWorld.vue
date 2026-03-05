<template>
  <v-container>
    <v-card class="pa-6" elevation="3">
      <v-card-title class="text-h5 font-weight-bold">
        Billing Details
      </v-card-title>

      <v-card-text>
        <v-form ref="billingForm">
          <!-- Customer Information -->
          <v-divider class="my-4" />
          <h3 class="text-subtitle-1 font-weight-bold mb-3">
            Customer Information
          </h3>

          <v-row>
            <v-col cols="12" md="4">
              <v-text-field
                label="Customer Name"
                v-model="billing.customerName"
                required
              />
            </v-col>

            <v-col cols="12" md="4">
              <v-text-field
                label="Customer Email"
                v-model="billing.customerEmail"
              />
            </v-col>

            <v-col cols="12" md="4">
              <v-text-field
                label="Age"
                type="number"
                v-model.number="billing.age"
                min="0"
              />
            </v-col>
          </v-row>

          <!-- Invoice Details -->
          <v-divider class="my-4" />
          <h3 class="text-subtitle-1 font-weight-bold mb-3">Invoice Details</h3>

          <v-row>
            <v-col cols="12" md="4">
              <v-text-field
                label="Invoice Number"
                v-model="billing.invoiceNo"
              />
            </v-col>

            <v-col cols="12" md="4">
              <v-text-field
                label="Invoice Date"
                type="date"
                v-model="billing.invoiceDate"
              />
            </v-col>

            <v-col cols="12" md="4">
              <v-text-field
                label="Due Date"
                type="date"
                v-model="billing.dueDate"
              />
            </v-col>
          </v-row>

          <!-- Bank Details -->
          <v-divider class="my-4" />
          <h3 class="text-subtitle-1 font-weight-bold mb-3">
            Bank Account Details
          </h3>

          <v-row>
            <v-col cols="12" md="6">
              <v-text-field label="Bank Name" v-model="billing.bankName" />
            </v-col>

            <v-col cols="12" md="6">
              <v-text-field
                label="Account Number"
                v-model="billing.accountNumber"
              />
            </v-col>

            <v-col cols="12" md="6">
              <v-text-field label="IFSC Code" v-model="billing.ifsc" />
            </v-col>

            <v-col cols="12" md="6">
              <v-text-field
                label="Account Holder Name"
                v-model="billing.accountHolder"
              />
            </v-col>
          </v-row>

          <!-- Amount Details -->
          <v-divider class="my-4" />
          <h3 class="text-subtitle-1 font-weight-bold mb-3">Amount Details</h3>

          <v-row>
            <v-col cols="12" md="4">
              <v-text-field
                label="Subtotal"
                type="number"
                v-model.number="billing.subtotal"
              />
            </v-col>

            <v-col cols="12" md="4">
              <v-text-field
                label="Tax (%)"
                type="number"
                v-model.number="billing.tax"
              />
            </v-col>

            <v-col cols="12" md="4">
              <v-text-field
                label="Total Amount"
                :value="totalAmount"
                readonly
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-card-actions class="justify-end">
        <v-btn color="primary" @click="submitBilling"> Submit Bill </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: "HelloWorld",

  data() {
    return {
      billing: {
        customerName: "",
        customerEmail: "",
        age: null,
        invoiceNo: "",
        invoiceDate: "",
        dueDate: "",
        bankName: "",
        accountNumber: "",
        ifsc: "",
        accountHolder: "",
        subtotal: 0,
        tax: 0,
      },
    };
  },

  computed: {
    totalAmount() {
      const taxAmount = (this.billing.subtotal * this.billing.tax) / 100;
      return this.billing.subtotal + taxAmount;
    },
  },

  methods: {
    submitBilling() {
      const payload = {
        ...this.billing,
        totalAmount: this.totalAmount,
      };

      console.log("Billing Payload:", payload);
      alert("Billing data submitted. Check console.");
    },
  },
};
</script>
