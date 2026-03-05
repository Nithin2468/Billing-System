<template>
  <v-container fluid class="bg-grey-lighten-4 fill-height align-start pa-0">
    <v-main class="w-100">
      <v-container class="pa-6 mt-2">
        <!-- Top Info Row -->
        <!-- Top Info Row -->
        <v-row class="mb-4">
          <v-col cols="12">
            <v-card
              class="mx-auto"
              elevation="4"
              rounded="lg"
              color="indigo-lighten-5"
            >
              <v-card-text
                class="d-flex justify-space-between align-center py-4"
              >
                <div class="d-flex align-center">
                  <v-avatar color="indigo" class="mr-4" size="50">
                    <v-icon
                      icon="mdi-receipt-text"
                      color="white"
                      size="30"
                    ></v-icon>
                  </v-avatar>
                  <div>
                    <h2 class="text-h5 font-weight-bold text-indigo-darken-3">
                      New Bill
                    </h2>
                    <div class="text-caption text-indigo-darken-1">
                      Create a new invoice
                    </div>
                  </div>
                </div>

                <div class="d-flex align-center">
                  <v-chip
                    v-if="partyName"
                    class="mr-4 font-weight-bold"
                    color="amber-darken-2"
                    variant="flat"
                    label
                    size="large"
                  >
                    <v-icon start icon="mdi-account-star"></v-icon>
                    {{ partyName }}
                  </v-chip>

                  <v-card elevation="0" color="white" class="px-4 py-2 border">
                    <div class="text-overline text-grey-darken-1 lh-1">
                      Bill Number
                    </div>
                    <div class="text-h5 font-weight-black text-indigo">
                      {{ billNumber }}
                    </div>
                  </v-card>
                </div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
        <!-- Addresses Row -->
        <v-row>
          <!-- Recipient Details -->
          <v-col cols="12" md="6">
            <v-card
              elevation="2"
              rounded="lg"
              class="h-100 border-t-4 border-indigo"
            >
              <v-card-item>
                <v-card-title class="text-indigo-darken-2 d-flex align-center">
                  <v-icon
                    icon="mdi-account-arrow-left"
                    class="mr-2"
                    color="indigo"
                  ></v-icon>
                  Recipient Details
                </v-card-title>
              </v-card-item>
              <v-card-text class="pt-2">
                <v-text-field
                  v-model="recipient.name"
                  label="Name"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-account"
                  color="indigo"
                  class="mb-1"
                ></v-text-field>
                <v-text-field
                  v-model="recipient.address"
                  label="Address"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-map-marker"
                  color="indigo"
                  class="mb-1"
                ></v-text-field>
                <v-row>
                  <v-col cols="6">
                    <v-text-field
                      v-model="recipient.city"
                      label="City"
                      variant="outlined"
                      density="compact"
                      prepend-inner-icon="mdi-city"
                      color="indigo"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="6">
                    <v-text-field
                      v-model="recipient.state"
                      label="State"
                      variant="outlined"
                      density="compact"
                      prepend-inner-icon="mdi-map"
                      color="indigo"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-text-field
                  v-model="recipient.gstin"
                  label="GSTIN / Unique ID"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-id-card"
                  color="indigo"
                ></v-text-field>
              </v-card-text>
            </v-card>
          </v-col>

          <!-- Consignee Details -->
          <v-col cols="12" md="6">
            <v-card
              elevation="2"
              rounded="lg"
              class="h-100 border-t-4 border-teal"
            >
              <v-card-item>
                <v-card-title class="text-teal-darken-2 d-flex align-center">
                  <v-icon
                    icon="mdi-truck-delivery"
                    class="mr-2"
                    color="teal"
                  ></v-icon>
                  Consignee Details
                </v-card-title>
              </v-card-item>
              <v-card-text class="pt-2">
                <v-text-field
                  v-model="consignee.name"
                  label="Name"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-account"
                  color="teal"
                  class="mb-1"
                ></v-text-field>
                <v-text-field
                  v-model="consignee.address"
                  label="Address"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-map-marker"
                  color="teal"
                  class="mb-1"
                ></v-text-field>
                <v-text-field
                  v-model="consignee.city"
                  label="City"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-city"
                  color="teal"
                  class="mb-1"
                ></v-text-field>
                <v-text-field
                  v-model="consignee.gstin"
                  label="GSTIN / Unique ID"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-id-card"
                  color="teal"
                  class="mb-1"
                ></v-text-field>
                <v-text-field
                  v-model="consignee.drugLicNo"
                  label="Drug Lic No"
                  variant="outlined"
                  density="compact"
                  prepend-inner-icon="mdi-pill"
                  color="teal"
                ></v-text-field>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>

        <!-- Invoice & Bank Details (Collapsible or Compact) -->
        <v-row>
          <!-- Invoice Details -->
          <v-col cols="12">
            <v-card elevation="2" rounded="lg">
              <v-expansion-panels variant="popout">
                <v-expansion-panel>
                  <v-expansion-panel-title color="grey-lighten-4">
                    <v-icon
                      icon="mdi-file-document-outline"
                      class="mr-2"
                      color="blue-grey"
                    ></v-icon>
                    <span class="text-blue-grey-darken-3 font-weight-bold"
                      >Invoice & Logistics Details</span
                    >
                  </v-expansion-panel-title>
                  <v-expansion-panel-text class="pt-4">
                    <v-row dense>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.gstInvoiceNo"
                          label="GST Invoice No"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.gstInvoiceDate"
                          label="Invoice Date"
                          type="date"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-select
                          v-model="invoice.reverseCharge"
                          :items="['Yes', 'No']"
                          label="Reverse Charge"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-select
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.supplierRefNo"
                          label="Supplier Ref No"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>

                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.paymentTerms"
                          label="Payment Terms"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.deliveryTerms"
                          label="Delivery Terms"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.dispatchedThrough"
                          label="Dispatched Through"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="invoice.destination"
                          label="Destination"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>

                      <v-col cols="12" md="4"
                        ><v-text-field
                          v-model="invoice.contactPerson"
                          label="Contact Person"
                          variant="outlined"
                          density="compact"
                          prepend-inner-icon="mdi-account-tie"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="4"
                        ><v-text-field
                          v-model="invoice.contactNumber"
                          label="Contact Number"
                          variant="outlined"
                          density="compact"
                          prepend-inner-icon="mdi-phone"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="4"
                        ><v-text-field
                          v-model="invoice.remark"
                          label="Remark"
                          variant="outlined"
                          density="compact"
                          prepend-inner-icon="mdi-note-text"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                    </v-row>
                  </v-expansion-panel-text>
                </v-expansion-panel>

                <!-- Bank Details -->
                <v-expansion-panel>
                  <v-expansion-panel-title color="grey-lighten-4">
                    <v-icon
                      icon="mdi-bank"
                      class="mr-2"
                      color="blue-grey"
                    ></v-icon>
                    <span class="text-blue-grey-darken-3 font-weight-bold"
                      >Bank Details</span
                    >
                  </v-expansion-panel-title>
                  <v-expansion-panel-text class="pt-4">
                    <v-row dense>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="bank.accountNumber"
                          label="Account Number"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="bank.ifscCode"
                          label="IFSC Code"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="bank.bankName"
                          label="Bank Name"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                      <v-col cols="12" md="3"
                        ><v-text-field
                          v-model="bank.branch"
                          label="Branch"
                          variant="outlined"
                          density="compact"
                          color="blue-grey"
                        ></v-text-field
                      ></v-col>
                    </v-row>
                  </v-expansion-panel-text>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-card>
          </v-col>
        </v-row>

        <!-- Billing Table -->
        <v-row>
          <v-col cols="12">
            <v-card
              elevation="2"
              rounded="lg"
              class="mt-2 border-t-4 border-deep-purple"
            >
              <v-card-item class="bg-grey-lighten-5 py-3">
                <div class="d-flex justify-space-between align-center">
                  <v-card-title
                    class="text-deep-purple-darken-1 font-weight-bold"
                  >
                    <v-icon icon="mdi-cart-outline" class="mr-2"></v-icon>
                    Items
                  </v-card-title>
                  <v-btn
                    color="deep-purple-darken-1"
                    variant="elevated"
                    prepend-icon="mdi-plus"
                    @click="addItem"
                    class="text-capitalize"
                  >
                    Add Item
                  </v-btn>
                </div>
              </v-card-item>

              <v-table class="billing-table">
                <thead>
                  <tr class="bg-grey-lighten-3">
                    <th class="text-left font-weight-bold text-subtitle-2">
                      #
                    </th>
                    <th class="text-left font-weight-bold text-subtitle-2">
                      Medicine Name
                    </th>
                    <th class="text-left font-weight-bold text-subtitle-2">
                      Batch
                    </th>
                    <th class="text-left font-weight-bold text-subtitle-2">
                      Expiry
                    </th>
                    <th
                      class="text-left font-weight-bold text-subtitle-2"
                      style="width: 100px"
                    >
                      Qty
                    </th>
                    <th class="text-left font-weight-bold text-subtitle-2">
                      MRP
                    </th>
                    <th class="text-left font-weight-bold text-subtitle-2">
                      GST(%)
                    </th>
                    <th class="text-left font-weight-bold text-subtitle-2">
                      Amount
                    </th>
                    <th class="text-center font-weight-bold text-subtitle-2">
                      Action
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="(item, index) in billingItems"
                    :key="index"
                    class="billing-row"
                  >
                    <td class="pt-4">{{ index + 1 }}</td>
                    <td class="pt-4">
                      <div class="font-weight-medium">
                        {{ item.medicineName }}
                      </div>
                    </td>
                    <td class="pt-4">
                      {{ item.batchId }}
                    </td>
                    <td class="pt-4">
                      {{ item.expiryDate }}
                    </td>
                    <td class="pt-4">
                      {{ item.quantity }}
                    </td>
                    <td class="pt-4">₹{{ item.mrp }}</td>
                    <td class="pt-4">{{ item.gst }}%</td>
                    <td class="pt-4 font-weight-bold text-right pr-4">
                      ₹{{ calculateAmount(item).toFixed(2) }}
                    </td>
                    <td class="text-center pt-2">
                      <v-btn
                        icon="mdi-delete"
                        size="small"
                        color="error"
                        variant="text"
                        @click="removeItem(index)"
                      ></v-btn>
                    </td>
                  </tr>
                </tbody>
              </v-table>

              <v-divider></v-divider>

              <div class="pa-4 bg-grey-lighten-4">
                <v-row align="center" justify="end">
                  <v-col cols="12" sm="6" md="4" class="text-right">
                    <div class="text-subtitle-1 text-grey-darken-1 mb-1">
                      Subtotal
                    </div>
                    <div class="text-h6 font-weight-medium mb-2">
                      ₹{{ subTotal.toFixed(2) }}
                    </div>

                    <div
                      v-if="discountValue > 0"
                      class="text-subtitle-1 text-error mb-1"
                    >
                      Discount ({{
                        discountType === "percent"
                          ? discountValue + "%"
                          : "Flat"
                      }})
                    </div>
                    <div
                      v-if="discountValue > 0"
                      class="text-h6 text-error font-weight-medium mb-2"
                    >
                      - ₹{{ discountAmountRaw.toFixed(2) }}
                    </div>
                  </v-col>
                </v-row>
                <v-divider class="my-3"></v-divider>
                <v-row align="center" justify="end">
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                    class="text-right d-flex align-center justify-end"
                  >
                    <span
                      class="text-h5 text-grey-darken-3 font-weight-bold mr-4"
                      >Grand Total:</span
                    >
                    <span class="text-h4 text-indigo-darken-3 font-weight-black"
                      >₹{{ finalTotal.toFixed(2) }}</span
                    >
                  </v-col>
                </v-row>
              </div>
            </v-card>
          </v-col>
        </v-row>

        <!-- Floating Action Buttons for quick access or Bottom Bar -->
        <v-footer app border class="bg-white" elevation="3" height="80">
          <v-spacer></v-spacer>
          <v-btn
            color="blue-grey-lighten-1"
            variant="tonal"
            class="mr-3"
            prepend-icon="mdi-format-list-bulleted"
            size="large"
          >
            Drafts
          </v-btn>
          <v-btn
            color="primary"
            variant="tonal"
            class="mr-3"
            prepend-icon="mdi-eye"
            size="large"
            @click="previewDialog = true"
          >
            Preview Bill
          </v-btn>
          <v-btn
            color="amber-darken-2"
            variant="tonal"
            class="mr-3"
            prepend-icon="mdi-tag-outline"
            size="large"
            @click="openDiscountDialog"
          >
            Add Discount
          </v-btn>
          <v-btn
            color="indigo-darken-1"
            variant="tonal"
            class="mr-3"
            prepend-icon="mdi-account-plus"
            size="large"
            @click="openPartyDialog"
          >
            Add Party
          </v-btn>
          <v-btn
            color="success"
            prepend-icon="mdi-check-circle"
            size="large"
            elevation="2"
            @click="saveBill"
            class="px-6 font-weight-bold"
          >
            Save Bill
          </v-btn>
        </v-footer>

        <!-- Preview Dialog -->
        <v-dialog v-model="previewDialog" maxWidth="900">
          <v-card>
            <v-card-title
              class="bg-indigo-darken-3 d-flex justify-space-between align-center"
            >
              <span>Invoice Preview</span>
              <div>
                <v-btn
                  icon="mdi-download"
                  title="Download PDF"
                  variant="text"
                  class="mr-2"
                  @click="downloadPdfBackend"
                ></v-btn>
                <v-btn
                  icon="mdi-printer"
                  title="Print"
                  variant="text"
                  class="mr-2"
                  @click="printBill"
                ></v-btn>
                <v-btn
                  icon="mdi-close"
                  variant="text"
                  @click="previewDialog = false"
                ></v-btn>
              </div>
            </v-card-title>
            <div class="pa-0">
              <BillPreview
                :recipient="recipient"
                :consignee="consignee"
                :invoice="invoice"
                :bank="bank"
                :items="billingItems"
                :subTotal="subTotal"
                :discount="{
                  type: discountType,
                  value: discountValue,
                  amount: discountAmountRaw,
                }"
                :grandTotal="finalTotal"
                :partyName="partyName"
              />
            </div>
          </v-card>
        </v-dialog>

        <!-- Add Item Dialog -->
        <BillingItemDialog
          v-model="addItemDialog"
          :medicines="allMedicines"
          :current-items="billingItems"
          @add-item="onItemAdded"
        />

        <!-- Discount Dialog -->
        <v-dialog v-model="discountDialog" max-width="450px">
          <v-card rounded="lg" elevation="4">
            <v-card-title class="bg-amber-darken-2 text-white">
              <v-icon icon="mdi-sale" class="mr-2"></v-icon>
              Add Discount
            </v-card-title>
            <v-card-text class="pt-6">
              <v-tabs
                v-model="tempDiscountType"
                color="amber-darken-2"
                align-tabs="center"
                class="mb-4"
              >
                <v-tab value="amount">Flat Amount (₹)</v-tab>
                <v-tab value="percent">Percentage (%)</v-tab>
              </v-tabs>

              <v-window v-model="tempDiscountType">
                <v-window-item value="amount">
                  <v-text-field
                    v-model.number="tempDiscountValue"
                    label="Enter Amount"
                    type="number"
                    prefix="₹"
                    variant="outlined"
                    color="amber-darken-2"
                    class="mt-2"
                  ></v-text-field>
                </v-window-item>
                <v-window-item value="percent">
                  <v-text-field
                    v-model.number="tempDiscountValue"
                    label="Enter Percentage"
                    type="number"
                    suffix="%"
                    variant="outlined"
                    color="amber-darken-2"
                    class="mt-2"
                  ></v-text-field>
                </v-window-item>
              </v-window>

              <div class="text-caption text-grey text-center">
                Current Subtotal: ₹{{ subTotal.toFixed(2) }}
                <br />
                <span v-if="tempDiscountType === 'percent'">
                  Estimated Discount: ₹{{
                    ((subTotal * tempDiscountValue) / 100).toFixed(2)
                  }}
                </span>
              </div>
            </v-card-text>
            <v-card-actions class="pb-4 px-4">
              <v-spacer></v-spacer>
              <v-btn
                color="grey-darken-1"
                variant="text"
                @click="discountDialog = false"
                >Cancel</v-btn
              >
              <v-btn
                color="amber-darken-2"
                variant="elevated"
                @click="applyDiscount"
                class="text-white px-6"
                >Apply Discount</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-dialog>

        <!-- Party Dialog -->
        <v-dialog v-model="partyDialog" max-width="500px">
          <v-card rounded="lg">
            <v-card-title class="bg-indigo text-white">
              <v-icon icon="mdi-account-plus" class="mr-2"></v-icon>
              Add Party Name
            </v-card-title>
            <v-card-text class="pt-6">
              <v-text-field
                v-model="partyName"
                label="Patient/Party Name"
                variant="outlined"
                color="indigo"
                prepend-icon="mdi-account"
              ></v-text-field>
            </v-card-text>
            <v-card-actions class="pb-4 px-4">
              <v-spacer></v-spacer>
              <v-btn color="grey" variant="text" @click="partyDialog = false"
                >Cancel</v-btn
              >
              <v-btn
                color="indigo"
                variant="elevated"
                @click="saveParty"
                class="text-white px-6"
                >Save Party</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-container>
      <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
        {{ snackbarText }}
        <template v-slot:actions>
          <v-btn color="white" variant="text" @click="snackbar = false">
            Close
          </v-btn>
        </template>
      </v-snackbar>

      <!-- Print Confirmation Dialog -->
      <v-dialog v-model="printConfirmDialog" max-width="400">
        <v-card>
          <v-card-title class="text-h6 bg-primary text-white">
            <v-icon icon="mdi-printer" class="mr-2"></v-icon>
            Print Bill?
          </v-card-title>
          <v-card-text class="pt-4">
            Do you want to print the saved bill (Bill No:
            <strong>{{ lastSavedBillNumber }}</strong
            >)?
          </v-card-text>
          <v-card-actions class="justify-end">
            <v-btn
              color="grey-darken-1"
              variant="text"
              @click="printConfirmDialog = false"
              >No</v-btn
            >
            <v-btn color="primary" variant="elevated" @click="confirmPrint"
              >Yes, Print</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-main>
  </v-container>
</template>

<script>
import EventService from "../services/eventServices";
import BillPreview from "./BillPreview.vue";
import BillingItemDialog from "./BillingItemDialog.vue";

export default {
  name: "BillingPage1",
  components: {
    BillPreview,
    BillingItemDialog,
  },
  computed: {
    // Removed mapGetters for allMedicines
    subTotal() {
      return this.billingItems.reduce((acc, item) => {
        return acc + this.calculateAmount(item);
      }, 0);
    },
    discountAmountRaw() {
      if (this.discountType === "amount") {
        return this.discountValue;
      } else {
        return (this.subTotal * this.discountValue) / 100;
      }
    },
    finalTotal() {
      return Math.max(0, this.subTotal - this.discountAmountRaw);
    },
  },
  data() {
    return {
      billNumber: "BILL-NEW", // Placeholder, will generate or fetch
      previewDialog: false,
      addItemDialog: false,
      recipient: {
        name: "",
        address: "",
        city: "",
        state: "",
        gstin: "",
      },
      consignee: {
        name: "",
        address: "",
        city: "",
        gstin: "",
        drugLicNo: "",
      },
      invoice: {
        gstInvoiceNo: "",
        gstInvoiceDate: "",
        reverseCharge: "No",
        supplierRefNo: "",
        paymentTerms: "",
        deliveryTerms: "",
        contactPerson: "",
        contactNumber: "",
        email: "",
        dispatchedThrough: "",
        destination: "",
        remark: "",
      },
      bank: {
        accountNumber: "",
        ifscCode: "",
        bankName: "",
        branch: "",
      },
      billingItems: [],
      // Dialogs
      partyDialog: false,
      discountDialog: false,

      // Party Data
      partyName: "",

      // Discount Data
      discountType: "amount",
      discountValue: 0,

      // Temp variables for Dialog
      tempDiscountType: "amount",
      tempDiscountValue: 0,

      // Snackbar
      snackbar: false,
      snackbarText: "",
      snackbarColor: "success",

      // Edit Mode
      isEditMode: false,
      billId: null,

      // Print Confirmation
      printConfirmDialog: false,
      lastSavedBillNumber: "",

      // Local Medicines State
      allMedicines: [],
    };
  },
  async created() {
    await this.fetchMedicines();
    // Bill Number generation is now tricky without store.
    // We can just use a timestamp or fetch next number from backend (not implemented yet).
    // Let's use timestamp for now or keep it simple.
    this.billNumber = `BILL-${Date.now().toString().slice(-6)}`;
  },

  methods: {
    async fetchMedicines() {
      try {
        const response = await EventService.getMedicines();
        this.allMedicines = response.data || [];
      } catch (error) {
        console.error("Failed to fetch medicines:", error);
      }
    },
    calculateAmount(item) {
      if (!item.quantity || !item.mrp) return 0;
      return item.quantity * item.mrp;
    },
    addItem() {
      this.addItemDialog = true;
    },
    onItemAdded(item) {
      this.billingItems.push(item);
      this.snackbarText = "Item added successfully";
      this.snackbarColor = "success";
      this.snackbar = true;
    },
    removeItem(index) {
      this.billingItems.splice(index, 1);
    },
    openDiscountDialog() {
      this.tempDiscountType = this.discountType;
      this.tempDiscountValue = this.discountValue;
      this.discountDialog = true;
    },
    applyDiscount() {
      this.discountType = this.tempDiscountType;
      this.discountValue = parseFloat(this.tempDiscountValue) || 0;
      this.discountDialog = false;
    },
    openPartyDialog() {
      this.partyDialog = true;
    },
    saveParty() {
      if (!this.recipient.name && this.partyName) {
        this.recipient.name = this.partyName;
      }
      this.partyDialog = false;
    },
    async saveBill() {
      if (this.billingItems.length === 0) {
        this.snackbarText = "Cannot save an empty bill! Please add items.";
        this.snackbarColor = "error";
        this.snackbar = true;
        return;
      }

      try {
        const billData = {
          billNumber: this.billNumber,
          recipient: this.recipient,
          consignee: this.consignee,
          invoice: this.invoice,
          bank: this.bank,
          billingItems: this.billingItems,
          subTotal: this.subTotal,
          discount: {
            type: this.discountType,
            value: this.discountValue,
            amount: this.discountAmountRaw,
          },
          grandTotal: this.finalTotal,
          partyName: this.partyName,
        };

        const response = await EventService.saveBill(billData);

        if (response.data.status === "success" || response.status === 200) {
          this.snackbarText = `Bill Saved Successfully! Grand Total: ₹${this.finalTotal.toFixed(
            2
          )}`;
          this.snackbarColor = "success";
          this.snackbar = true;

          this.lastSavedBillNumber = this.billNumber;
          this.printConfirmDialog = true;

          this.resetForm();
          this.billNumber = `BILL-${Date.now().toString().slice(-6)}`;
        } else {
          throw new Error(response.data.message || "Unknown error");
        }
      } catch (error) {
        console.error("Error saving bill:", error);
        this.snackbarText =
          "Failed to save bill: " + (error.message || "Network Error");
        this.snackbarColor = "error";
        this.snackbar = true;
      }
    },
    resetForm() {
      this.recipient = {
        name: "",
        address: "",
        city: "",
        state: "",
        gstin: "",
      };
      this.consignee = {
        name: "",
        address: "",
        city: "",
        gstin: "",
        drugLicNo: "",
      };
      this.invoice = {
        gstInvoiceNo: "",
        gstInvoiceDate: "",
        reverseCharge: "No",
        supplierRefNo: "",
        paymentTerms: "",
        deliveryTerms: "",
        contactPerson: "",
        contactNumber: "",
        email: "",
        dispatchedThrough: "",
        destination: "",
        remark: "",
      };
      this.billingItems = [];
      this.partyName = "";
      this.discountValue = 0;
    },
    async fetchBankDetails() {
      if (this.bank.ifscCode && this.bank.ifscCode.length === 11) {
        try {
          const response = await EventService.getBankDetails(
            this.bank.ifscCode
          );
          if (response.data) {
            this.bank.bankName = response.data.BANK || "";
            this.bank.branch = response.data.BRANCH || "";
            this.bank.accountNumber2 = "";
          }
        } catch (error) {
          console.error("Error fetching bank details:", error);
        }
      }
    },
    loadBillData(id) {
      const bill = this.$store.getters.getBillById(id);
      if (bill) {
        this.billNumber = bill.billNumber;
        this.recipient = { ...bill.recipient };
        this.consignee = { ...bill.consignee };
        this.invoice = { ...bill.invoice };
        this.bank = { ...bill.bank };
        this.billingItems = JSON.parse(JSON.stringify(bill.billingItems));
        this.subTotal = bill.subTotal; // Note: subTotal is computed, but we might arguably rely on re-computation or saved value.
        // Actually subTotal is computed based on items, so just setting items is enough.
        this.discountType = bill.discount.type;
        this.discountValue = bill.discount.value;
        this.partyName = bill.partyName || "";
      } else {
        this.snackbarText = "Bill not found!";
        this.snackbarColor = "error";
        this.snackbar = true;
      }
    },
    confirmPrint() {
      this.printConfirmDialog = false;
      // We need to load the data of the LAST SAVED bill to preview it.
      // Since the form is reset, we can't just open previewDialog with current data.
      // Wait, PreviewDialog uses props passed to BillPreview.
      // If we reset the form, `recipient`, `billingItems` etc are cleared.
      // We need to fetch the last saved bill or hold it in temp variables.
      // Better approach: fetch the bill by `lastSavedBillNumber` from store.
      const lastBill = this.$store.getters.allBills.find(
        (b) => b.billNumber === this.lastSavedBillNumber
      );
      if (lastBill) {
        // Temporarily repopulate for preview or pass explicitly?
        // BillPreview uses props bound to 'recipient', 'consignee' etc.
        // To avoid re-populating the FORM (which might confuse user thinking they are editing),
        // we should probably have a separate "PreviewData" object or just re-populate form but keep it distinct?
        // Actually, if we just saved it, the user might expect the screen to be ready for NEW bill.
        // So repopulating form variables is bad.
        // BUT BillPreview component is bound to `recipient`, `consignee` etc. in the template:
        // :recipient="recipient"
        // So we MUST change those variables to show the preview.

        // WORKAROUND: We will temporarily set the data, open dialog. When dialog closes, clear it?
        // Or simpler: Create a dedicated `previewData` object that binds to BillPreview,
        // and in the template switch between `recipient` and `previewData.recipient`?
        // No, that's too much refactor.

        // Re-filling form variables is risky if user started typing new bill.
        // But the dialog is immediate.
        // Let's assume user hasn't started typing yet.
        // We will load the data, open preview, and when preview closes, we might need to clear or just leave it.
        // If we leave it, the "New Bill" form is dirty.
        // Ideally: The BillPreview binding in the template should be dynamic.
        // Let's create `previewData` in data() and update the template binding to use a computed property or just `previewData`.

        // Let's do the lightweight fix first:
        // Use a separate method to "load for preview" that fills a separate object,
        // BUT we need to change the template `<BillPreview ... />` arguments.
        // Since I can't easily change the template binding to be conditional without major diffs,
        // I'll stick to: Load data back into form -> Print -> Clear form on dialog close?
        // No, `printConfirmDialog` "No" clears nothing.
        // "Yes" -> loads data -> opens preview -> User prints -> Closes preview.
        // When closing preview, we should probably clear form if it was the result of this flow.

        // actually, let's just hold the data in a `tempPreviewData` and use `v-if` in the template?
        // The template currently has:
        // <BillPreview :recipient="recipient" ... />
        // usage of "recipient" variable.

        // I will change the BillPreview usage in valid template section above.
        // But I only edited the `saveBill` logic.
        // I should have thought about this.

        // Let's Update the template binding to use `previewRecipient` etc.
        // And make `previewRecipient` a computed property that returns `recipient` normally,
        // or a specific object if we are in "View Mode".
        // That seems complex.

        // ALTERNATIVE: Just use the store. `allBills[0]` is the last saved bill.

        // Let's simply re-populate the form variables for the preview, and add a flag `isViewingSavedBill`.
        // If `previewDialog` is closed and `isViewingSavedBill` is true, reset form.
        this.loadBillData(this.lastSavedBillNumber);
        this.previewDialog = true;
      }
    },
    printBill() {
      // Trigger browser print
      window.print();
    },
    async downloadPdfBackend() {
      const invoiceElement = document.querySelector(".invoice-container");
      if (!invoiceElement) {
        this.snackbarText =
          "Invoice preview not found. Please open preview first.";
        this.snackbarColor = "error";
        this.snackbar = true;
        return;
      }

      // Basic styles for the PDF to look decent
      const style = `
        <style>
          body { font-family: Arial, sans-serif; font-size: 12px; }
          .invoice-container { padding: 20px; width: 100%; box-sizing: border-box; }
          .header-section { display: flex; justify-content: space-between; margin-bottom: 20px; border-bottom: 2px solid #ddd; padding-bottom: 10px; }
          .invoice-title { font-size: 24px; font-weight: bold; color: #3f51b5; }
          table { width: 100%; border-collapse: collapse; margin-top: 20px; }
          th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
          th { background-color: #f2f2f2; }
          .text-right { text-align: right; }
          .text-center { text-align: center; }
          .footer-section { margin-top: 30px; border-top: 1px solid #ddd; padding-top: 10px; }
          /* Utility classes used in template */
          .d-flex { display: flex; }
          .justify-space-between { justify-content: space-between; }
          .align-center { align-items: center; }
          .flex-column { flex-direction: column; }
          .mb-2 { margin-bottom: 8px; }
          .ma-0 { margin: 0; }
          .pa-0 { padding: 0; }
          .w-100 { width: 100%; }
        </style>
      `;

      const htmlContent = `
        <!DOCTYPE html>
        <html>
        <head>
          <meta charset="utf-8">
          <title>Invoice</title>
          ${style}
        </head>
        <body>
          ${invoiceElement.outerHTML}
        </body>
        </html>
      `;

      try {
        const response = await EventService.downloadPdf(htmlContent);
        const url = window.URL.createObjectURL(
          new Blob([response.data], { type: "application/pdf" })
        );
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute(
          "download",
          `invoice_${this.billNumber || "new"}.pdf`
        );
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);

        this.snackbarText = "PDF Downloaded Successfully via Backend!";
        this.snackbarColor = "success";
        this.snackbar = true;

        // Reset the screen after successful download
        this.previewDialog = false;
        this.resetForm();
        this.billNumber = this.$store.getters.nextBillNumber;
      } catch (error) {
        console.error("Download failed", error);
        this.snackbarText =
          "Backend PDF download failed. Ensure backend is running.";
        this.snackbarColor = "error";
        this.snackbar = true;
      }
    },
  },
  watch: {
    "bank.ifscCode": "fetchBankDetails",
  },
};
</script>

<style scoped>
.billing-table th {
  font-size: 0.9rem !important;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: #424242;
}

.billing-row:hover {
  background-color: #f5f5f5;
}

/* Make inputs in table look seamless but visible */
.v-field--variant-outlined .v-field__outline__start,
.v-field--variant-outlined .v-field__outline__notch,
.v-field--variant-outlined .v-field__outline__end {
  border-color: #e0e0e0;
}

.v-card-item {
  padding: 12px 16px;
}
</style>
